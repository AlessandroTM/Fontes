package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Config struct {
	Address string
	Region  string
	Profile string
	ID      string
	Secret  string
}

func New(config Config) (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(config.ID, config.Secret, ""),
				Region:           aws.String(config.Region),
				Endpoint:         aws.String(config.Address),
				S3ForcePathStyle: aws.Bool(true),
			},
			Profile: config.Profile,
		},
	)
}

func main() {
	// Configuração da sessão AWS para o S3 local (LocalStack)
	sess, err := New(Config{
		Address: "http://localhost:4566",
		Region:  "eu-west-1",
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Criar um novo serviço S3
	svc := s3.New(sess)

	// Nome dos arquivos no bucket S3 simulado pelo LocalStack
	nomesArquivo := "Nomes.txt"
	sobrenomesArquivo := "Sobrenomes.txt"
	resultadoArquivo := "Resultado.txt"

	// Lê o conteúdo dos arquivos locais
	nomes, err := lerArquivoS3(svc, nomesArquivo)
	if err != nil {
		fmt.Println("Erro ao ler arquivo de nomes:", err)
		return
	}

	sobrenomes, err := lerArquivoS3(svc, sobrenomesArquivo)
	if err != nil {
		fmt.Println("Erro ao ler arquivo de sobrenomes:", err)
		return
	}

	// Cria um arquivo local com o resultado
	resultado, err := os.Create(resultadoArquivo)
	if err != nil {
		fmt.Println("Erro ao criar arquivo de resultado:", err)
		return
	}
	defer resultado.Close()

	// Escreve no arquivo de resultado os nomes completos
	for m, nome := range nomes {
		for n, sobrenome := range sobrenomes {
			if m == n {
				nomeCompleto := fmt.Sprintf("%s %s\n", nome, sobrenome)
				_, err := resultado.WriteString(nomeCompleto)
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo de resultado:", err)
					return
				}
			}
		}
	}

	// Upload do arquivo de resultado para o bucket S3
	uploadResultadoParaS3(svc, resultadoArquivo, resultadoArquivo)
	fmt.Println("Arquivo de resultado foi enviado para o bucket S3 com sucesso!")
}

// Função para ler um arquivo de um bucket S3
func lerArquivoS3(svc *s3.S3, nomeArquivo string) ([]string, error) {
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("bucket-teste-ale1234"), // Nome do seu bucket S3 simulado
		Key:    aws.String(nomeArquivo),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	var linhas []string
	scanner := bufio.NewScanner(result.Body)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return linhas, nil
}

// Função para enviar um arquivo para um bucket S3
func uploadResultadoParaS3(svc *s3.S3, nomeArquivoLocal string, nomeArquivoS3 string) {
	arquivo, err := os.Open(nomeArquivoLocal)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de resultado:", err)
		return
	}
	defer arquivo.Close()

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("bucket-teste-ale1234"), // Nome do seu bucket S3 simulado
		Key:    aws.String(nomeArquivoS3),
		Body:   arquivo,
	})
	if err != nil {
		fmt.Println("Erro ao enviar arquivo para o S3:", err)
		return
	}
}
