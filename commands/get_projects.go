package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Projeto struct {
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetProjects(token string) ([]Projeto, error) {
	// Cria uma requisição GET para a URL da API de projetos
	req, err := http.NewRequest("GET", "http://localhost:3000/api/project", nil)
	if err != nil {
		return nil, err
	}

	// Adiciona o token JWT como um cabeçalho de autorização na requisição
	req.Header.Set("Authorization", "Bearer "+token)

	// Cria um cliente HTTP
	client := &http.Client{}

	// Envia a requisição e obtém a resposta
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica o código de status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao obter projetos: status code %d", resp.StatusCode)
	}

	// Decodifica o JSON da resposta para uma slice de Projetos
	var projetos []Projeto
	if err := json.NewDecoder(resp.Body).Decode(&projetos); err != nil {
		return nil, err
	}

	fmt.Println(projetos)

	return projetos, nil
}
