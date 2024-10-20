package utils

import (
	"SubHunter0x/SecurityTrails/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct para a API
type SecurityTrailsAPI struct {
	APIKey  string
	BaseURL string
}

// Struct para a resposta da API
type SubdomainResponse struct {
	Subdomains []string `json:"subdomains"`
}

// Busca subdomínios na API SecurityTrails
func (api *SecurityTrailsAPI) FindSubdomains(domain string) ([]string, error) {
	// Cria a requisição utilizando a função utilitária
	req, err := utils.CreateAPIRequest(api.APIKey, api.BaseURL, fmt.Sprintf("/v1/domain/%s/subdomains", domain))
	if err != nil {
		return nil, err
	}

	// Executa a requisição (não precisa passar o client)
	resp, err := utils.ExecuteAPIRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica se o status code é 200 (OK)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao buscar subdomínios: %s", resp.Status)
	}

	// Decodifica a resposta JSON
	var subdomainResponse SubdomainResponse
	if err := json.NewDecoder(resp.Body).Decode(&subdomainResponse); err != nil {
		return nil, err
	}

	// Retorna os subdomínios
	return subdomainResponse.Subdomains, nil
}
