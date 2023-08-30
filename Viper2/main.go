package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	println("Valor 1: ", DotEnv("Variavel_1"))
	println("Valor 2: ", DotEnv("Variavel_2"))
	println("Valor 3: ", DotEnv("Variavel_3"))
	println("Valor 4: ", DotEnv("Variavel_4"))
}
