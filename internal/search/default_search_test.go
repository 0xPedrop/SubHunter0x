package search

import (
	"os"
	"testing"
)

func TestLoadWordList(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test_wordlist.txt")
	if err != nil {
		t.Fatalf("Erro ao criar arquivo temporário: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString("sub1\nsub2\nsub3\n")
	if err != nil {
		t.Fatalf("Erro ao escrever no arquivo: %v", err)
	}

	tmpfile.Close()

	wordlist, err := LoadWordList(tmpfile.Name())
	if err != nil {
		t.Fatalf("Erro ao carregar a wordlist: %v", err)
	}

	expected := []string{"sub1", "sub2", "sub3"}
	if len(wordlist) != len(expected) {
		t.Errorf("Esperado %d subdomínios, mas obteve %d", len(expected), len(wordlist))
	}
	for i, v := range expected {
		if wordlist[i] != v {
			t.Errorf("Esperado %s, mas obteve %s", v, wordlist[i])
		}
	}
}
