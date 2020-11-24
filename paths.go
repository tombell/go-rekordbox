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

func getAgentOptionsPath(root string) string {
	return filepath.Join(root, appSupportPath, agentOptionsPath)
}

func getAsarPath(root string) string {
	return filepath.Join(root, asarPath)
}

func getDatabasePath(root string) string {
	return filepath.Join(root, masterDBPath)
}

func getImagePath(dataPath, imagePath string) string {
	return filepath.Join(dataPath, "share", imagePath)
}

func getLibraryPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("os user home dir failed: %w", err)
	}

	return filepath.Join(home, appDataPath), nil
}
