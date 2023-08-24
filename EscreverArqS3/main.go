package main

import (
	"bytes"
	"fmt"
	"log"

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

func main() {

	bucket := aws.String("bucket-teste-ale1234")
	keyname := aws.String("test123-parte3.txt")

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

	svc := s3.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	payload := "Alessandro Gabriel Marques"
	var tamanho int = len(payload)
	t := int64(tamanho)

	params := &s3.PutObjectInput{
		Bucket: bucket,  // Required
		Key:    keyname, // Required
		ACL:    aws.String("bucket-owner-full-control"),
		Body:   bytes.NewReader([]byte(payload)),
		//ContentLength: aws.Int64(27),
		ContentLength: aws.Int64(t),
	}
	resp, err := svc.PutObject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
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
