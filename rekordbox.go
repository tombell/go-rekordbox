package rekordbox

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Track ...
type Track struct {
	ID        string
	Number    int
	Artist    string
	Name      string
	ImagePath string
}

// GetEncryptedPassword ...
func GetEncryptedPassword(appPath string) (string, error) {
	asarPath := getAsarPath(appPath)

	f, err := os.Open(asarPath)
	if err != nil {
		return "", fmt.Errorf("os open failed: %w", err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	re := regexp.MustCompile(`pass: ".*"`)
	result := re.FindAllString(string(data), 10)[0]
	password := strings.Split(result, ": ")[1]
	password = strings.Replace(password, `"`, "", -1)

	return password, nil
}

// OpenDatabase ...
func OpenDatabase(appPath, encryptionKey string) (*sql.DB, error) {
	appDataPath, err := getLibraryPath()
	if err != nil {
		return nil, fmt.Errorf("get library path failed: %w", err)
	}

	databasePath := getDatabasePath(appDataPath)

	dsn := fmt.Sprintf("file:"+databasePath+"?_key=%s", encryptionKey)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql open failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	return db, nil
}

// GetRecentTrack ...
func GetRecentTrack(db *sql.DB) (*Track, error) {
	row := db.QueryRow(`
		SELECT
			h.ID,
			h.TrackNo,
			Name,
			Title,
			ImagePath
		FROM djmdSongHistory AS h
		JOIN djmdContent AS c on h.ContentID = c.ID
		LEFT JOIN djmdArtist as a on c.ArtistID = a.ID
		GROUP BY h.created_at
		ORDER BY h.created_at DESC
		LIMIT 1
	`)

	var id string
	var trackNo int
	var artist string
	var name string
	var imagePath string

	if err := row.Scan(&id, &trackNo, &artist, &name, &imagePath); err != nil {
		return nil, fmt.Errorf("row scan failed: %w", err)
	}

	return &Track{
		ID:        id,
		Number:    trackNo,
		Artist:    artist,
		Name:      name,
		ImagePath: imagePath,
	}, nil
}