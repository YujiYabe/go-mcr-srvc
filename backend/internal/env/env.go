package env

import (
	"context"
	"encoding/json"
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
	STG   = "STG"
	PRD   = "prd"
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
		v.SetConfigName("local.env")
		//local環境ではlocalstackを使う
		localstack(*v)

	case DEV:
		v.SetConfigName("dev.env")
		//dev環境ではaws環境

	case STG:
		v.SetConfigName("stg.env")
		//stg環境ではaws環境

	case PRD:
		v.SetConfigName("prd.env")
		//prd環境ではaws環境

	default:
		log.Fatalf("failed to serve: invalid environment")
	}

	if err := v.ReadInConfig(); err != nil {
		log.Println("== == == == == == == == == == ")
		log.Printf("%#v\n", err)
		log.Println("== == == == == == == == == == ")
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

func localstack(viper viper.Viper) {

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

	var secretString = aws.ToString(result.SecretString)
	var secrets LocalstackSecrets
	if err := json.Unmarshal([]byte(secretString), &secrets); err != nil {
		log.Printf("Failed to unmarshal secret string: %v", err)
	}
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", secrets.MyLocalSecret.SecretString)
	log.Println("== == == == == == == == == == ")

	var secretString2 SecretString
	if err := json.Unmarshal([]byte(secrets.MyLocalSecret.SecretString), &secretString2); err != nil {
		log.Printf("Failed to unmarshal secret string: %v", err)
	}
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", secretString2.Username)
	log.Printf("%#v\n", secretString2.Password)
	log.Println("== == == == == == == == == == ")

}

type LocalstackSecrets struct {
	MyLocalSecret struct {
		SecretString string `json:"SecretString"`
	} `json:"my-local-secret"`
}

type SecretString struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
