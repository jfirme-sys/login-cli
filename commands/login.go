// Package commands implementa funcionalidades específicas da CLI.
package commands

import (
	"envsecret-cli/auth" // Importa o pacote de autenticação
	"fmt"
)

// Login realiza a operação de login usando o email e senha fornecidos.
// Chama a função Authenticate do pacote auth para autenticar o usuário.
// Imprime o token JWT obtido se o login for bem-sucedido, ou imprime um
// erro caso contrário.
func Login(email, password string) {
	// Chama a função de autenticação do pacote auth para obter o token JWT
	token, err := auth.Authenticate(email, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Imprime o token JWT se a autenticação for bem-sucedida
	fmt.Println("Token:", token)
}
