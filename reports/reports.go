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

func SaveReport(domain string, subdomains []string, filePath string) error {
	var reports []Report

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("não foi possível abrir o arquivo: %w", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&reports); err != nil {
			var singleReport Report
			if err := json.NewDecoder(file).Decode(&singleReport); err == nil {
				reports = append(reports, singleReport)
			} else {
				return fmt.Errorf("erro ao decodificar o arquivo JSON: %w", err)
			}
		}
	}

	report := Report{
		Domain:     domain,
		Subdomains: subdomains,
	}
	reports = append(reports, report)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("não foi possível criar o arquivo: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(reports); err != nil {
		return fmt.Errorf("erro ao salvar o relatório: %w", err)
	}

	return nil
}
