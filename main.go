package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <API_KEY>")
		return
	}

	apiKey := os.Args[1]
	domain := "example.com" // Substitua pelo domínio que você deseja testar

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.securitytrails.com/v1/subdomains/%s", domain), nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	req.Header.Set("APIKEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao chamar a API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("A chave API é válida e permite acesso ao endpoint de subdomínios.")
	} else {
		fmt.Printf("Erro: %s (Código de status: %d)\n", resp.Status, resp.StatusCode)
	}
}
