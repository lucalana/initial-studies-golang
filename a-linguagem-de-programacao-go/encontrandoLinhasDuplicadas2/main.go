package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exibe a contagem e o texto das linhas que aparecem mais de uma vez na entrada. Ele lê de stdin ou
// de uma lista de arquivos nomeados.
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Nome arquivo: %s %d\t%s\n", os.Args[0], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Nota: ignorando erros em potencial do input.Err()
}
