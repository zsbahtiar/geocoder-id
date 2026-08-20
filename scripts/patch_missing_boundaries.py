#!/usr/bin/env python3
"""
Patch missing administrative boundaries into geocoder.duckdb.

Adds districts/villages that exist in the official Kepmendagri 2025 reference
(via lokabisa-oss/region-id) but are absent from the DB, using polygons from
cahyadsn/wilayah_boundaries (aligned to Kepmendagri No 300.2.2-2430/2025).

Only inserts entries that have a real polygon in the source. Entries whose
boundary is not yet digitized by BIG are skipped (reported), and points there
degrade gracefully to the parent polygon via ST_Contains.

Usage: python3 scripts/patch_missing_boundaries.py --input data/geocoder.duckdb --output data/geocoder.duckdb
"""
import argparse, csv, io, json, os, re, sys, urllib.request
import duckdb

REGION_ID = "https://github.com/lokabisa-oss/region-id/releases/download/v1.0.1"
WB_RAW    = "https://raw.githubusercontent.com/cahyadsn/wilayah_boundaries/main/db"
CACHE     = os.environ.get("PATCH_CACHE", "/tmp/wb_cache")

def dot(u):
    if len(u)==2:  return u
    if len(u)==4:  return f"{u[:2]}.{u[2:]}"
    if len(u)==6:  return f"{u[:2]}.{u[2:4]}.{u[4:]}"
    if len(u)==10: return f"{u[:2]}.{u[2:4]}.{u[4:6]}.{u[6:]}"
    raise ValueError(u)

def fetch(url, path):
    os.makedirs(os.path.dirname(path), exist_ok=True)
    if not os.path.exists(path):
        try:
            urllib.request.urlretrieve(url, path)
        except Exception:
            open(path,"w").close()  # cache miss (404) as empty
    return path

def load_csv(name):
    p = fetch(f"{REGION_ID}/{name}", f"{CACHE}/{name}")
    with open(p, newline="") as f:
        return list(csv.DictReader(f))

ROW_RE_TMPL = r"\('{code}','((?:[^']|'')*)',\s*(-?[\d.]+|NULL),\s*(-?[\d.]+|NULL),\s*'(\[[^']*\])'\)"
def extract_geom(code_dot, sqlfile):
    if not os.path.getsize(sqlfile): return None
    txt = open(sqlfile).read()
    m = re.search(ROW_RE_TMPL.format(code=re.escape(code_dot)), txt)
    if not m: return None
    name = m.group(1).replace("''","'")
    lat  = None if m.group(2)=="NULL" else float(m.group(2))
    lng  = None if m.group(3)=="NULL" else float(m.group(3))
    coords = json.loads(m.group(4))
    def depth(a):
        d=0
        while isinstance(a,list): a=a[0]; d+=1
        return d
    def ring(r):
        s=[[p[1],p[0]] for p in r]          # [lat,lng] -> [lng,lat]
        if s and s[0]!=s[-1]: s.append(s[0]) # close ring
        return s
    dp = depth(coords)
    if dp==3:   gj={"type":"Polygon","coordinates":[ring(r) for r in coords]}
    elif dp==4: gj={"type":"MultiPolygon","coordinates":[[ring(r) for r in poly] for poly in coords]}
    else: return None
    return name, lat, lng, json.dumps(gj)

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--input",  default="data/geocoder.duckdb")
    ap.add_argument("--output", default="data/geocoder.duckdb")
    a = ap.parse_args()
    if a.output != a.input:
        import shutil; shutil.copy(a.input, a.output)

    con = duckdb.connect(a.output); con.execute("INSTALL spatial; LOAD spatial;")

    # reference (Kepmendagri 2025) + hierarchy lookups
    prov,reg,dist,vill = {},{},{},{}
    for r in load_csv("regions_id.csv"):
        prov[r["province_code"]] = r["province_name"]
        reg[r["regency_code"]]   = (r["province_code"], r["regency_name"])
        dist[r["district_code"]] = (r["province_code"], r["regency_code"], r["district_name"])
        vill[r["village_code"]]  = (r["province_code"], r["regency_code"], r["district_code"], r["village_name"])
    ref_dist = set(dist); ref_vill = set(vill)

    db_dist = {c[0].replace(".","") for c in con.execute("SELECT code FROM locations WHERE level='district'").fetchall()}
    db_vill = {c[0].replace(".","") for c in con.execute("SELECT code FROM locations WHERE level='village'").fetchall()}
    miss_dist = sorted(ref_dist - db_dist)
    miss_vill = sorted(ref_vill - db_vill)
    print(f"missing districts={len(miss_dist)} villages={len(miss_vill)}")

    def kec_file(u): return fetch(f"{WB_RAW}/kec/wilayah_boundaries_kec_{u[:2]}.sql", f"{CACHE}/kec_{u[:2]}.sql")
    def kel_file(u): return fetch(f"{WB_RAW}/kel/{u[:2]}/wilayah_boundaries_kel_{u[:2]}.{u[2:4]}.sql", f"{CACHE}/kel_{u[:2]}.{u[2:4]}.sql")

    add_d = add_v = skip = 0
    def insert(code, name, level, lat, lng, gj):
        con.execute("INSERT OR REPLACE INTO locations(code,name,level,latitude,longitude,geom) "
                    "VALUES (?,?,?,?,?, ST_MakeValid(ST_GeomFromGeoJSON(?)))",[code,name,level,lat,lng,gj])

    for u in miss_dist:
        g = extract_geom(dot(u), kec_file(u))
        if not g: skip+=1; continue
        _,lat,lng,gj = g; pc,rc,dn = dist[u]; cd=dot(u)
        insert(cd, dn, "district", lat, lng, gj)
        con.execute("INSERT OR REPLACE INTO hierarchy VALUES (?,?,?,?,?,?,?,?,?)",
                    [cd,pc,prov[pc],rc,reg[rc][1],cd,dn,None,None]); add_d+=1

    for u in miss_vill:
        g = extract_geom(dot(u), kel_file(u))
        if not g: skip+=1; continue
        _,lat,lng,gj = g; pc,rc,dc,vn = vill[u]; cd=dot(u)
        insert(cd, vn, "village", lat, lng, gj)
        con.execute("INSERT OR REPLACE INTO hierarchy VALUES (?,?,?,?,?,?,?,?,?)",
                    [cd,pc,prov[pc],rc,reg[rc][1],dc,dist[dc][2],cd,vn]); add_v+=1

    con.execute("DROP INDEX IF EXISTS idx_locations_geom;")
    con.execute("CREATE INDEX idx_locations_geom ON locations USING RTREE (geom);")
    con.commit()
    print(f"added districts={add_d} villages={add_v} | skipped(no polygon in source)={skip}")
    for lvl,n in con.execute("SELECT level,count(*) FROM locations GROUP BY 1 ORDER BY 1").fetchall():
        print(f"  {lvl:9} {n}")
    con.close()

if __name__ == "__main__":
    main()
