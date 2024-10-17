package main

import (
	"SubHunter0x/internal/search"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go-subdomain-finder <domínio>")
		return
	}

	domain := os.Args[1]

	wordlist, err := search.LoadWordList("wordlist/common_subdomains.txt")
	if err != nil {
		fmt.Printf("Erro ao carregar a wordlist: %v\n", err)
		return
	}

	foundSubdomains, err := search.FindSubdomains(domain, wordlist)
	if err != nil {
		fmt.Printf("Erro ao encontrar subdomínios: %v\n", err)
		return
	}

	for _, subdomain := range foundSubdomains {
		fmt.Println(subdomain)
	}
}
