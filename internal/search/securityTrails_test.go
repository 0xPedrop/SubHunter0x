package search

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock da API do SecurityTrails para testes
func setupMockServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		// Simule uma resposta de sucesso com subdomínios
		fmt.Fprintln(w, `{"subdomains": ["www.example.com", "api.example.com"]}`)
	})

	return httptest.NewServer(handler)
}

func TestFindSubdomains(t *testing.T) {
	// Configurar o servidor mock
	server := setupMockServer()
	defer server.Close()

	api := SecurityTrailsAPI{APIKey: "YOUR_API_KEY", BaseURL: server.URL} // Use a URL do mock

	// Testar com um domínio conhecido
	domain := "example.com"
	subdomains, err := api.FindSubdomains(domain)
	if err != nil {
		t.Fatalf("Erro ao buscar subdomínios: %v", err)
	}

	if len(subdomains) == 0 {
		t.Error("Nenhum subdomínio encontrado")
	}

	expected := []string{"www.example.com", "api.example.com"}
	for _, e := range expected {
		found := false
		for _, s := range subdomains {
			if s == e {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Subdomínio esperado %s não encontrado", e)
		}
	}
}
