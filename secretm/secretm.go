package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/Josh2604/literaluser/awsgo"
	"github.com/Josh2604/literaluser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > getting secret: ", secretData)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return secretData, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &secretData)
	fmt.Println(" > read secret ok")
	return secretData, nil
}
