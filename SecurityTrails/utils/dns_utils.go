package utils

import (
	"bufio"
	"context"
	"net"
	"os"
	"time"
)

// Carrega a wordlist de subdomínios a partir de um arquivo
func LoadWordList(filePath string) ([]string, error) {
	var wordlist []string

	// Abre o arquivo da wordlist
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Lê o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	// Verifica erros durante a leitura
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordlist, nil
}

// ResolveDNS tenta resolver o subdomínio usando o DNS
func ResolveDNS(subdomain string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resolver := net.Resolver{PreferGo: true}
	_, err := resolver.LookupHost(ctx, subdomain)

	return err == nil
}
