package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cloudquery/cloudquery/cmd"
	"github.com/cloudquery/cloudquery/deploy"
	"os"
)

func main() {
	if env := os.Getenv("AWS_LAMBDA_FUNCTION_NAME"); env != "" {
		lambda.Start(deploy.LambdaHandler)
	} else {
		cmd.Execute()
	}

}
