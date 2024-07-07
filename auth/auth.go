package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Authenticate realiza a autenticação usando um endpoint de login da API.
// Recebe o email e senha como parâmetros, envia uma requisição POST para o
// endpoint de login, e retorna o token JWT obtido na resposta se a autenticação
// for bem-sucedida.
func Authenticate(email, password string) (string, error) {
	// Prepara o corpo da requisição JSON com o email e senha
	reqBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	// Envia a requisição POST para o endpoint de login da API
	resp, err := http.Post("http://localhost:3000/api/users/login", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Verifica se a resposta tem o status HTTP 201 Created (sucesso na criação)
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("authentication failed: status code %d", resp.StatusCode)
	}

	// Decodifica a resposta JSON para obter o token JWT
	var res map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	// Extrai e retorna o token JWT da resposta
	token, ok := res["token"]
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}
