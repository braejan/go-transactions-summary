package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadFileToS3(content string, fileName string) error {
	// Create a new AWS session using environment variables
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %v", err)
	}

	// Create a S3 client using the session
	svc := s3.New(sess)
	file := strings.NewReader(content)
	// S3 configuration options
	params := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   file,
	}

	// Load the file into the S3 bucket
	_, err = svc.PutObject(params)
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %v", err)
	}

	return nil
}

func handleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// extract file from request body
	body := event.Body
	cleanBody, fileName, err := cleanStringBody(body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: err.Error()}, nil
	}
	err = uploadFileToS3(cleanBody, fileName)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: err.Error()}, nil
	}
	response := "👏👏👏 Tu archivo ha sido subido a S3 exitosamente. Un proceso interno lo estará ejecutando. 😉"
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: response}, nil
}

func main() {
	lambda.Start(handleRequest)
}

func cleanStringBody(body string) (cleanBody string, fileName string, err error) {
	// return strings.Replace(strings.Replace(body, "\n", "", -1), "\r", "", -1)
	splitted := strings.Split(body, "\n")
	// Validate Header Content-Disposition
	fmt.Println("splitted[4]", splitted[4])
	if !strings.Contains(splitted[1], "Content-Disposition: form-data;") {
		err = fmt.Errorf("invalid header Content-Disposition: missing form-data")
		return
	}
	// Content-Disposition: form-data; name="file"; filename="txns.csv"
	// Find the variable name:
	variables := strings.Split(splitted[1], ";")
	if len(variables) < 3 {
		err = fmt.Errorf("invalid header Content-Disposition: missing variables")
		return
	}
	fileName = strings.Split(variables[2], "=")[1]
	fileName = strings.ReplaceAll(fileName, "\"", "")
	fileName = strings.Trim(fileName, " ")
	if fileName == "" {
		err = fmt.Errorf("invalid header Content-Disposition: missing filename")
		return
	}
	for _, line := range splitted[4:] {
		cleanBody += line + "\n"
	}
	if strings.Trim(cleanBody, " ") == "" {
		err = fmt.Errorf("invalid body: empty file")
		return
	}
	return
}
