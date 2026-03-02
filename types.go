package geocoder

type Result struct {
	Province *Location `json:"province,omitempty"`
	Regency  *Location `json:"regency,omitempty"`
	District *Location `json:"district,omitempty"`
	Village  *Location `json:"village,omitempty"`
}

type Location struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Coord struct {
	Lat float64
	Lon float64
}

type Level string

const (
	LevelProvince Level = "province"
	LevelRegency  Level = "regency"
	LevelDistrict Level = "district"
	LevelVillage  Level = "village"
)
