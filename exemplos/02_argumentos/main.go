package main

import (
	"fmt"
	"os"
	"strings"
)

// main imprime os argumentos passados pela linha de comando.
func main() {
	args := os.Args[1:] // Ignora o nome do programa
	if len(args) == 0 {
		fmt.Println("Uso: go run main.go <arg1> <arg2> ...")
		return
	}

	resultado := strings.Join(args, " ")
	fmt.Println("Argumentos recebidos:", resultado)
}
