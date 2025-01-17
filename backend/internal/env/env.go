package env

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/viper"
)

const (
	LOCAL = "local"
	DEV   = "dev"
	PROD  = "prod"
)

var (
	backendHost = "backend"
	GoEchoPort  string
	GRPCPort    string
	GRPCAddress string
	TZ          string
	PostgresDSN string
)

func init() {
	v := viper.New()
	v.AutomaticEnv()
	v.AddConfigPath("internal/env")
	v.SetConfigType("env")

	switch v.GetString("ENV") {
	case LOCAL:
		v.SetConfigName(".localenv")
		//local環境ではlocalstackを使う
		localstack(*v)

	case DEV:
		v.SetConfigName(".devenv")

	case PROD:
		v.SetConfigName(".prodenv")

	default:
		log.Fatalf("failed to serve: invalid environment")
	}

	if err := v.ReadInConfig(); err != nil {
		log.Println("== == == == == == == == == == ")
		log.Printf("%#v\n", err)
		log.Println("== == == == == == == == == == ")

		// Continue even if file doesn't exist
	}

	TZ = v.GetString("TZ")
	GoEchoPort = v.GetString("GO_ECHO_PORT")
	GRPCPort = v.GetString("GRPC_PORT")
	GRPCAddress = backendHost + ":" + GRPCPort

	PostgresDSN = "host=postgres" +
		" user=" + v.GetString("POSTGRES_USER") +
		" password=" + v.GetString("POSTGRES_PASSWORD") +
		" port=" + v.GetString("POSTGRES_BACK_PORT") +
		" dbname=app" +
		" TimeZone=" + v.GetString("TZ") +
		" sslmode=disable"

}

func localstack(viper.Viper) {
	// customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	if service == secretsmanager.ServiceID {
	// 		return aws.Endpoint{
	// 			// URL: "http://localhost:4566/", // LocalStack のエンドポイント
	// 			URL: "http://localstack:4566", // LocalStack のエンドポイント
	// 		}, nil
	// 	}
	// 	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	// })

	creds := credentials.NewStaticCredentialsProvider("test", "test", "")

	secretName := "my-local-secret"
	region := "ap-northeast-1"

	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region),
		config.WithBaseEndpoint("http://localstack:4566"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {

		log.Fatal(err.Error())
	}

	var secretString string = aws.ToString(result.SecretString)
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", secretString)
	log.Println("== == == == == == == == == == ")

}

// Secrets Managerからシークレットを取得する関数
// func getSecretValue(client *secretsmanager.Client, secretName string) (string, error) {
// 	input := &secretsmanager.GetSecretValueInput{
// 		SecretId: aws.String(secretName),
// 	}

// 	result, err := client.GetSecretValue(context.TODO(), input)
// 	if err != nil {
// 		return "", err
// 	}

// 	return aws.ToString(result.SecretString), nil
// }
