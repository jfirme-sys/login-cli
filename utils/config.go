package utils

import (
	"os"
	"path/filepath"
)

func GetConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".envsecret-cli")
}

func EnsureConfigDir() error {
	dir := GetConfigDir()
	return os.MkdirAll(dir, 0755)
}

func SaveSelectedProject(projectId string) error {
	dir := GetConfigDir()
	filePath := filepath.Join(dir, "selected_project")
	return os.WriteFile(filePath, []byte(projectId), 0644)
}

func LoadSelectedProject() (string, error) {
	dir := GetConfigDir()
	filePath := filepath.Join(dir, "selected_project")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
