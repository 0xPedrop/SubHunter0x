package main

import (
	"SubHunter0x/SecurityTrails/utils" // Para utilitários de API e DNS
	"SubHunter0x/internal/search"      // Para busca local e com API
	"SubHunter0x/reports"              // Para salvar relatórios em JSON
	"flag"
	"fmt"
)

// Função que exibe o banner da ferramenta
func showBanner() {
	green := "\033[32m"
	reset := "\033[0m"
	banner := `
███████╗██╗   ██╗██████╗ ██╗  ██╗██╗   ██╗███╗   ██╗████████╗███████╗██████╗  ██████╗ ██╗  ██╗
██╔════╝██║   ██║██╔══██╗██║  ██║██║   ██║████╗  ██║╚══██╔══╝██╔════╝██╔══██╗██╔═████╗╚██╗██╔╝
███████╗██║   ██║██████╔╝███████║██║   ██║██╔██╗ ██║   ██║   █████╗  ██████╔╝██║██╔██║ ╚███╔╝ 
╚════██║██║   ██║██╔══██╗██╔══██║██║   ██║██║╚██╗██║   ██║   ██╔══╝  ██╔══██╗████╔╝██║ ██╔██╗ 
███████║╚██████╔╝██████╔╝██║  ██║╚██████╔╝██║ ╚████║   ██║   ███████╗██║  ██║╚██████╔╝██╔╝ ██╗
╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   
                	===========================================                               
    			  SubHunter0x - Subdomain Enumeration Tool
			  https://github.com/0xPedrop/SubHunter0x

					Developed by: 
				https://github.com/0xPedrop
				https://github.com/0xJotave
						  
				All honor and glory to God!
    			===========================================
`
	fmt.Println(green + banner + reset)
}

func main() {
	// Definindo as flags para a linha de comando
	apiKey := flag.String("a", "", "Chave da API para SecurityTrails (opcional)")
	domain := flag.String("d", "", "Domínio para realizar a busca de subdomínios (obrigatório)")

	flag.Parse()

	// Verifica se o domínio foi fornecido
	if *domain == "" {
		fmt.Println("Erro: o parâmetro '-d' (domínio) é obrigatório.")
		flag.Usage()
		return
	}

	// Exibe o banner
	showBanner()

	// Diretório de relatórios
	reportFilePath := "./reports/reports.json"

	// Verifica se a busca será via API ou via wordlist local
	if *apiKey != "" {
		// Busca via API SecurityTrails
		api := search.SecurityTrailsAPI{
			APIKey:  *apiKey,
			BaseURL: "https://api.securitytrails.com",
		}

		subdomains, err := api.FindSubdomains(*domain)
		if err != nil {
			fmt.Println("Erro ao buscar subdomínios pela API:", err)
			return
		}

		// Exibe e salva os subdomínios encontrados via API
		fmt.Println("Subdomínios encontrados via API:")
		for _, subdomain := range subdomains {
			fmt.Println(subdomain)
		}

		// Salva o relatório
		err = reports.SaveReport(*domain, subdomains, reportFilePath)
		if err != nil {
			fmt.Println("Erro ao salvar o relatório:", err)
		} else {
			fmt.Printf("Relatório salvo em %s\n", reportFilePath)
		}
	} else {
		// Busca via wordlist local
		wordlistPath := "./wordlist/common_subdomains.txt"
		wordlist, err := utils.LoadWordList(wordlistPath)
		if err != nil {
			fmt.Println("Erro ao carregar o wordlist:", err)
			return
		}

		// Busca DNS usando a wordlist
		subdomains, err := search.FindSubdomains(*domain, wordlist)
		if err != nil {
			fmt.Println("Erro ao buscar subdomínios:", err)
			return
		}

		// Exibe e salva os subdomínios encontrados
		fmt.Println("Subdomínios encontrados via wordlist:")
		for _, subdomain := range subdomains {
			fmt.Println(subdomain)
		}

		// Salva o relatório
		err = reports.SaveReport(*domain, subdomains, reportFilePath)
		if err != nil {
			fmt.Println("Erro ao salvar o relatório:", err)
		} else {
			fmt.Printf("Relatório salvo em %s\n", reportFilePath)
		}
	}
}
