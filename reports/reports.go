package reports

import (
	"encoding/json"
	"fmt"
	"os"
)

type Report struct {
	Domain     string   `json:"domain"`
	Subdomains []string `json:"subdomains"`
}

// SaveReport salva o relatório de subdomínios no arquivo reports.json
func SaveReport(domain string, subdomains []string, filePath string) error {
	var reports []Report

	// Verifica se o arquivo já existe
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("não foi possível abrir o arquivo: %w", err)
		}
		defer file.Close()

		// Decodifica o arquivo existente para recuperar os relatórios antigos
		if err := json.NewDecoder(file).Decode(&reports); err != nil {
			return fmt.Errorf("erro ao decodificar o arquivo JSON: %w", err)
		}
	}

	// Cria um novo relatório para o domínio atual
	report := Report{
		Domain:     domain,
		Subdomains: subdomains,
	}
	reports = append(reports, report)

	// Cria ou sobrescreve o arquivo JSON com os relatórios
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("não foi possível criar o arquivo: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	// Salva os relatórios no arquivo
	if err := encoder.Encode(reports); err != nil {
		return fmt.Errorf("erro ao salvar o relatório: %w", err)
	}

	return nil
}
