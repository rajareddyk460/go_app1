package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func Test() (string, error) {

	return fmt.Sprintf("Welcome to AWS_Session"), nil
}

func main() {

	lambda.Start(Test)
}
