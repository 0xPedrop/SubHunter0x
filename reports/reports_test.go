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

	var reports []Report // Mude aqui para um slice de Report
	if err := json.NewDecoder(file).Decode(&reports); err != nil {
		t.Fatalf("Erro ao decodificar o JSON: %v", err)
	}

	// Verifique se o último relatório salvo corresponde ao domínio e subdomínios esperados
	lastReport := reports[len(reports)-1] // Pegue o último relatório
	if lastReport.Domain != domain {
		t.Errorf("Esperado: %s, mas obteve: %s", domain, lastReport.Domain)
	}

	if len(lastReport.Subdomains) != len(subdomains) {
		t.Errorf("Esperado: %d subdomínios, mas obteve: %d", len(subdomains), len(lastReport.Subdomains))
	}

	os.Remove(outputPath)
}
