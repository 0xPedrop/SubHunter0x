package reports

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSaveReport(t *testing.T) {
	domain := "example.com"
	subdomains := []string{"www.example.com", "api.example.com", "mail.example.com"}
	outputPath := "test_report.json"

	err := SaveReport(domain, subdomains, outputPath)
	if err != nil {
		t.Fatalf("Erro ao salvar o relatório: %v", err)
	}

	file, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Erro ao abrir o relatório JSON: %v", err)
	}
	defer file.Close()

	var report Report
	if err := json.NewDecoder(file).Decode(&report); err != nil {
		t.Fatalf("Erro ao decodificar o JSON: %v", err)
	}

	if report.Domain != domain {
		t.Errorf("Esperado: %s, mas obteve: %s", domain, report.Domain)
	}

	if len(report.Subdomains) != len(subdomains) {
		t.Errorf("Esperado: %d subdomínios, mas obteve: %d", len(subdomains), len(report.Subdomains))
	}

	os.Remove(outputPath)
}
