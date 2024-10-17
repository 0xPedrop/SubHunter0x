package main

import (
	"SubHunter0x/internal/search"
	"SubHunter0x/reports"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go-subdomain-finder <domínio>")
		return
	}

	domain := os.Args[1]

	// Carregar Lista com os Subdominios do Arquivo
	wordlist, err := search.LoadWordList("wordlist/common_subdomains.txt")
	if err != nil {
		fmt.Printf("Erro ao carregar a wordlist: %v\n", err)
		return
	}

	// Encontrar Subdominios Existentes e Funcionais

	foundSubdomains, err := search.FindSubdomains(domain, wordlist)
	if err != nil {
		fmt.Printf("Erro ao encontrar subdomínios: %v\n", err)
		return
	}

	for _, subdomain := range foundSubdomains {
		fmt.Println(subdomain)
	}

	// Relatórios/Reports

	err = reports.SaveReport(domain, foundSubdomains, "reports/report.json")
	if err != nil {
		fmt.Printf("Erro ao gerar o relatório: %v\n", err)
		return
	}

	fmt.Println("Relatório JSON gerado com sucesso em reports/report.json")
}
