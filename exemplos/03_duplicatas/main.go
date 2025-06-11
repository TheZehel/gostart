package main

import (
	"bufio"
	"fmt"
	"os"
)

// Este programa lê linhas da entrada padrão (stdin) e exibe as duplicadas com suas contagens.
func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite linhas de texto (Ctrl+D para encerrar no Linux/macOS ou Ctrl+Z no Windows):")

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	fmt.Println("\nLinhas duplicadas:")
	for linha, contagem := range counts {
		if contagem > 1 {
			fmt.Printf("%d\t%s\n", contagem, linha)
		}
	}
}
