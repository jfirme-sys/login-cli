package utils

import (
	"io/ioutil"
)

const (
	tokenFile           = "token.txt"
	selectedProjectFile = "selected_project.txt"
)

func SaveToken(token string) error {
	return ioutil.WriteFile(tokenFile, []byte(token), 0644)
}

func LoadToken() (string, error) {
	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SaveSelectedProject(projectId string) error {
	return ioutil.WriteFile(selectedProjectFile, []byte(projectId), 0644)
}

func LoadSelectedProject() (string, error) {
	data, err := ioutil.ReadFile(selectedProjectFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
