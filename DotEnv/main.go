package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	os.Setenv("NOME_DEV", "Alessandro")

	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("Nome do Desenvolvedor:", os.Getenv("NOME_DEV"))

	// Exemplo de uma variavel não declarada
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Lista todas as variaveis do ambiente
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
