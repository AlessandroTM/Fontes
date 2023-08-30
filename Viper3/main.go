package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func main() {

	viperenv1 := viperEnvVariable("Variavel_1")
	fmt.Printf("viper : %s = %s \n", "Variavel_1", viperenv1)

	viperenv2 := viperEnvVariable("Variavel_2")
	fmt.Printf("viper : %s = %s \n", "Variavel_2", viperenv2)

	viperenv3 := viperEnvVariable("Variavel_3")
	fmt.Printf("viper : %s = %s \n", "Variavel_3", viperenv3)

	viperenv4 := viperEnvVariable("Variavel_4")
	fmt.Printf("viper : %s = %s \n", "Variavel_4", viperenv4)
}
