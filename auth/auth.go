package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Authenticate(email, password string) (string, error) {
	reqBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:3000/api/users/login", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("authentication failed: status code %d", resp.StatusCode)
	}

	var res map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	token, ok := res["token"]
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}
