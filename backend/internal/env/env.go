package env

import (
	"context"
	"encoding/json"
	"log"
	"os"

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
	viperViper := viper.New()
	viperViper.AutomaticEnv()
	viperViper.AddConfigPath("internal/env")
	viperViper.SetConfigType("env")

	switch os.Getenv("ENV") { // 環境情報のみ OS環境変数からを取得
	case LOCAL:
		viperViper.SetConfigName("lcl.env")
		//local環境ではlocalstackを使う

	case DEV:
		viperViper.SetConfigName("dev.env")

	case STG:
		viperViper.SetConfigName("stg.env")

	case PRD:
		viperViper.SetConfigName("prd.env")

	default:
		log.Fatalf("failed to serve: invalid environment")
	}

	if err := viperViper.ReadInConfig(); err != nil {
		log.Println("== == == == == == == == == == ")
		log.Printf("%#v\n", err)
		log.Println("== == == == == == == == == == ")
	}

	switch os.Getenv("ENV") { // 環境情報のみ OS環境変数からを取得

	case LOCAL:
		//local環境ではlocalstackを使う
		localstack(viperViper)

	case DEV:
		//dev環境ではaws環境

	case STG:
		//stg環境ではaws環境

	case PRD:
		//prd環境ではaws環境

	default:
		log.Fatalf("failed to serve: invalid environment")
	}

	TZ = viperViper.GetString("TZ")
	GoEchoPort = viperViper.GetString("GO_ECHO_PORT")
	GRPCPort = viperViper.GetString("GRPC_PORT")
	GRPCAddress = backendHost + ":" + GRPCPort

	PostgresDSN = "host=postgres" +
		" user=" + viperViper.GetString("POSTGRES_USER") +
		" password=" + viperViper.GetString("POSTGRES_PASSWORD") +
		" port=" + viperViper.GetString("POSTGRES_BACK_PORT") +
		" dbname=app" +
		" TimeZone=" + viperViper.GetString("TZ") +
		" sslmode=disable"

}

func localstack(
	viperViper *viper.Viper,
) {

	creds := credentials.NewStaticCredentialsProvider(
		viperViper.GetString("AWS_STATIC_CREDENTIAL_KEY"),
		viperViper.GetString("AWS_STATIC_CREDENTIAL_SECRET"),
		"",
	)

	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(viperViper.GetString("AWS_REGION")),
		config.WithBaseEndpoint(viperViper.GetString("AWS_ENDPOINT")),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(viperViper.GetString("AWS_SECRET_NAME")),
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

	var secretString2 SecretString
	if err := json.Unmarshal([]byte(secrets.MyLocalSecret.SecretString), &secretString2); err != nil {
		log.Printf("Failed to unmarshal secret string: %v", err)
	}

	viper.Set("db_user", secretString2.Username)
	viper.Set("db_password", secretString2.Password)
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
