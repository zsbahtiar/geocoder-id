package geocoder

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	dbDownloadURL = "https://github.com/zsbahtiar/geocoder-id/releases/download/v%s/geocoder.duckdb"
	latestVersion = "0.1.2"
)

var Version = "dev"

func downloadDatabase(destPath string) error {
	version := Version
	if version == "dev" || version == "" {
		version = latestVersion
	}

	url := fmt.Sprintf(dbDownloadURL, version)

	fmt.Fprintf(os.Stderr, "Downloading database from %s...\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: HTTP %d", resp.StatusCode)
	}

	dir := filepath.Dir(destPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpPath := destPath + ".tmp"
	out, err := os.Create(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	written, err := io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("failed to write file: %w", err)
	}

	if err := os.Rename(tmpPath, destPath); err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("failed to finalize file: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Downloaded %d MB to %s\n", written/1024/1024, destPath)
	return nil
}

func getDefaultDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".geocoder-id", defaultDBName), nil
}
