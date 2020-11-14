package rekordbox

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	agentOptionsPath = "/rekordboxAgent/storage/options.json"
	appDataPath      = "/Library/Pioneer/rekordbox"
	appSupportPath   = "/Library/Application Support/Pioneer"
	asarPath         = "/Contents/MacOS/rekordboxAgent.app/Contents/Resources/app.asar"
	masterDBPath     = "master.db"
)

// GetAsarPath returns the path of the Rekordbox Agent ASAR file.
func GetAsarPath(root string) string {
	return filepath.Join(root, asarPath)
}

// GetLibraryPath returns the path to the Rekordbox directory in the user
// library directory.
func GetLibraryPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("os user home dir failed: %w", err)
	}

	return filepath.Join(home, appDataPath), nil
}

// GetDatabasePath returns the path of the Rekordbox master database file.
func GetDatabasePath(root string) string {
	return filepath.Join(root, masterDBPath)
}

// GetImagePath returns the full path of the given image path.
func GetImagePath(dataPath, imagePath string) string {
	return filepath.Join(dataPath, "share", imagePath)
}

func getAgentOptionsPath(root string) string {
	return filepath.Join(root, appSupportPath, agentOptionsPath)
}
