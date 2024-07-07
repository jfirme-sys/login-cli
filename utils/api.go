package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Secret struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Login(username, password string) (string, error) {
	loginData := map[string]string{"username": username, "password": password}
	body, _ := json.Marshal(loginData)
	req, err := http.NewRequest("POST", "http://localhost:3000/login", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed: %s", resp.Status)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["token"], nil
}

func FetchProjects(token string) ([]Project, error) {
	req, err := http.NewRequest("GET", "https://api.example.com/projects", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch projects: %s", resp.Status)
	}

	var projects []Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func SelectProject(projectId string) error {
	// You can store this selection in a config file or environment variable
	// For simplicity, we print it here
	fmt.Println("Project selected:", projectId)
	return nil
}

func FetchSecrets(token string) (map[string]string, error) {
	req, err := http.NewRequest("GET", "https://api.example.com/secrets", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch secrets: %s", resp.Status)
	}

	var secrets []Secret
	if err := json.NewDecoder(resp.Body).Decode(&secrets); err != nil {
		return nil, err
	}

	secretMap := make(map[string]string)
	for _, secret := range secrets {
		secretMap[secret.Key] = secret.Value
	}

	return secretMap, nil
}
