package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SecurityTrailsAPI struct {
	APIKey  string
	BaseURL string
}

type SubdomainResponse struct {
	Subdomains []string `json:"subdomains"`
}

// FindSubdomains faz a busca de subdomínios via API SecurityTrails
func (api *SecurityTrailsAPI) FindSubdomains(domain string) ([]string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/domain/%s/subdomains", api.BaseURL, domain), nil)
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

	// Concatena o domínio principal com cada subdomínio retornado, se necessário
	var fullSubdomains []string
	for _, subdomain := range subdomainResponse.Subdomains {
		// Verifica se o subdomínio já contém o domínio principal
		if !strings.Contains(subdomain, domain) {
			// Concatena o subdomínio com o domínio
			fullSubdomains = append(fullSubdomains, fmt.Sprintf("%s.%s", subdomain, domain))
		} else {
			// Subdomínio já está completo, adicione diretamente
			fullSubdomains = append(fullSubdomains, subdomain)
		}
	}

	return fullSubdomains, nil
}
