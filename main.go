package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/awsgo"
	"gambituser/db"
	"gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()
	if !ValidateParameters() {
		fmt.Println("Error in Parameters. Should send 'SecretName'")
		err := errors.New("error in Parameters. Should send 'SecretName'")
		return event, err
	}
	var data models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error when reading Secret " + err.Error())
		return event, err
	}
	err = db.SignUp(data)
	return event, err
}

func ValidateParameters() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
