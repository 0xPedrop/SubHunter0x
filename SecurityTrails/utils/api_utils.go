package utils

import (
	"fmt"
	"net/http"
	"time"
)

// Cria uma nova requisição GET para a API SecurityTrails
func CreateAPIRequest(apiKey, baseURL, endpoint string) (*http.Request, error) {
	// Concatena a baseURL e o endpoint para formar a URL completa
	url := fmt.Sprintf("%s%s", baseURL, endpoint)

	// Cria a requisição HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Adiciona o cabeçalho com a API Key
	req.Header.Set("APIKEY", apiKey)
	return req, nil
}

// Executa a requisição e retorna a resposta
func ExecuteAPIRequest(req *http.Request) (*http.Response, error) {
	// Cria o cliente HTTP dentro da função
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
