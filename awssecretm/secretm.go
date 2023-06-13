package awsecretm

import (
	"backend/lamda-golang-backend-gambit/awsgo"
	"backend/lamda-golang-backend-gambit/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

/*
Obtener Secrets de AWS SecretManager
*/
func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > Pedir Secreto " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println(" > Lectura de Secret OK " + secretName)
	return secretData, nil
}
