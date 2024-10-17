package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SecurityTrailsAPI struct {
	APIKey  string
	BaseURL string // Adicionando uma URL base
}

type SubdomainResponse struct {
	Subdomains []string `json:"subdomains"`
}

func (api *SecurityTrailsAPI) FindSubdomains(domain string) ([]string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Usando a URL base definida no struct
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/subdomains/%s", api.BaseURL, domain), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("APIKEY", api.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao buscar subdomínios: %s", resp.Status)
	}

	var subdomainResponse SubdomainResponse
	if err := json.NewDecoder(resp.Body).Decode(&subdomainResponse); err != nil {
		return nil, err
	}
	return subdomainResponse.Subdomains, nil
}
