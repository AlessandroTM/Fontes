package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Registro struct {
	Codigo    int    `csv:"codigo"`
	Nome      string `csv:"nome"`
	Sobrenome string `csv:"sobrenome"`
	Idade     int    `csv:"idade"`
}

type Config struct {
	Address string
	Region  string
	Profile string
	ID      string
	Secret  string
}

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

func main() {

	sess, err := Newc(Config{
		Address: "http://localhost:4566",
		Region:  "us-east-1",
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	s3Client := s3.New(sess)
	bucket := "bucket-teste-ale1234"
	key := "teste.csv"

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	result, err := s3Client.GetObject(requestInput)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()
	body1, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString1 := string(body1)

	fmt.Println(bodyString1)

	var reader = csv.NewReader(strings.NewReader(bodyString1))
	reader.Comma = ','
	var registro Registro
	for {
		err := Unmarshal(reader, &registro)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(registro)
		fmt.Printf("%s %s has age of %d\n", registro.Nome, registro.Sobrenome, registro.Idade)
	}
}

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}

func Newc(config Config) (*session.Session, error) {
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
