package rekordbox

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/tombell/go-rekordbox/internal/crypto"
)

func getDatabasePassword(appPath string) (string, error) {
	cfg, err := parseAgentConfig()
	if err != nil {
		return "", fmt.Errorf("parse agent config failed: %w", err)
	}

	encodedPasswordData := cfg.Options[1][1]

	decodedPasswordData, err := base64.StdEncoding.DecodeString(encodedPasswordData)
	if err != nil {
		return "", fmt.Errorf("base64 decode string failed: %w", err)
	}

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

	passwordBytes := []byte(password)

	decryptedBytes, err := crypto.Decrypt(decodedPasswordData, passwordBytes)
	if err != nil {
		return "", fmt.Errorf("crypto decrypt failed: %w", err)
	}

	return string(decryptedBytes), nil
}
