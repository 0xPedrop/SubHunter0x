package search

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// LoadWordList carrega a lista de subdomínios a partir de um arquivo
func LoadWordList(filePath string) ([]string, error) {
	var wordlist []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordlist, nil
}

// ResolveDNS tenta resolver o DNS para um subdomínio
func ResolveDNS(subdomain string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resolver := net.Resolver{PreferGo: true}
	_, err := resolver.LookupHost(ctx, subdomain)

	return err == nil
}

// FindSubdomains encontra os subdomínios válidos para o domínio usando a wordlist
func FindSubdomains(domain string, wordlist []string) ([]string, error) {
	var foundSubdomains []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, sub := range wordlist {
		wg.Add(1)
		go func(sub string) {
			defer wg.Done()
			subdomain := fmt.Sprintf("%s.%s", sub, domain)
			if ResolveDNS(subdomain) {
				mu.Lock()
				foundSubdomains = append(foundSubdomains, subdomain)
				mu.Unlock()
			}
		}(sub)
	}

	wg.Wait()
	return foundSubdomains, nil
}
