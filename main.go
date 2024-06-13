package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Josh2604/literaluser/awsgo"
	"github.com/Josh2604/literaluser/db"
	"github.com/Josh2604/literaluser/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()
	if !ValidParams() {
		fmt.Println("Error on params, it must send SecretName")
		err := errors.New("error on params, it must send SecretName")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email: ", data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("UUID: ", data.UserUUID)
		}
	}
	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error on read secret: ", err)
		return event, err
	}

	err = db.SignUp(data)
	if err != nil {
		fmt.Println("error on sign up: ", err)
		return event, err
	}

	return event, err
}

func ValidParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
