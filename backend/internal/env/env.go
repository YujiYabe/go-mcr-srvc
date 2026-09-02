package env

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/viper"
)

type serverConfig struct {
	BackendHost string
	GoEchoPort  string
	GRPCPort    string
	GRPCAddress string
}

type databaseConfig struct {
	DSN string
}

type auth0Config struct {
	Domain       string
	TokenURL     string
	Audience     string
	GrantType    string
	ClientSecret string
}

type pubSubConfig struct {
	BootstrapServers string
	ConsumerGroupID  string
	TestTopic        string
	OtherTopic       string
	FlushTimeoutMS   int
	SampleUserName   string
}

type redisConfig struct {
	Addr     string
	Password string
	DB       int
}

type Config struct {
	Server   serverConfig
	Database databaseConfig
	Auth0    auth0Config
	PubSub   pubSubConfig
	Redis    redisConfig
}

func Load() (
	config Config,
	err error,
) {
	// OS環境変数で環境を切り替える
	// 機密情報以外はXXX.envに記載。secret managerのキーはgithub secretsに保存?
	// 機密情報はsecret managerに保存

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	configName := env
	if env == "lcl" {
		configName = "local"
	}
	viperViper := initViper()

	viperViper.SetConfigName(configName + ".env")
	if returnedErr := viperViper.ReadInConfig(); returnedErr != nil {
		config, err = Config{}, fmt.Errorf("load environment file: %w", returnedErr)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	if env == "lcl" {
		if returnedErr := setupLocalstack(viperViper); returnedErr != nil {
			config, err = Config{}, fmt.Errorf("setup localstack: %w", returnedErr)
			return //nolint:nakedret // Use the project-wide named return convention.
		}
	}

	config, err = Config{
		Server:   newServerConfig(viperViper),
		Database: newDatabaseConfig(viperViper),
		Auth0:    newAuth0Config(viperViper),
		PubSub:   newPubSubConfig(viperViper),
		Redis:    newRedisConfig(viperViper),
	}, nil
	return //nolint:nakedret // Use the project-wide named return convention.
}

func initViper() (
	viperConfig *viper.Viper,
) {
	viperConfig = viper.New()
	viperConfig.AutomaticEnv()
	viperConfig.AddConfigPath("internal/env")
	viperConfig.SetConfigType("env")
	return
}

func setupLocalstack(
	viperViper *viper.Viper,
) (
	err error,
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
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(viperViper.GetString("AWS_SECRET_NAME")),
	}

	result, returnedErr := svc.GetSecretValue(context.TODO(), input)
	if returnedErr != nil {
		err = returnedErr
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	resultSecretString := aws.ToString(result.SecretString)
	localstackSecrets := &LocalstackSecrets{}
	if returnedErr := json.Unmarshal(
		[]byte(resultSecretString),
		localstackSecrets,
	); returnedErr != nil {
		err = returnedErr
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	var secretString SecretString
	if returnedErr := json.Unmarshal(
		[]byte(localstackSecrets.MyLocalSecret.SecretString),
		&secretString,
	); returnedErr != nil {
		err = returnedErr
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	viperViper.Set("POSTGRES_USER", secretString.Username)
	viperViper.Set("POSTGRES_PASSWORD", secretString.Password)

	err = nil
	return //nolint:nakedret // Use the project-wide named return convention.
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

func newServerConfig(
	viperViper *viper.Viper,
) (
	config serverConfig,
) {
	config = serverConfig{
		BackendHost: viperViper.GetString("BACKEND_HOST"),
		GoEchoPort:  viperViper.GetString("GO_ECHO_PORT"),
		GRPCPort:    viperViper.GetString("GRPC_PORT"),
		GRPCAddress: fmt.Sprintf(
			"%s:%s",
			viperViper.GetString("BACKEND_HOST"),
			viperViper.GetString("GRPC_PORT"),
		),
	}
	return
}

func newDatabaseConfig(
	viperViper *viper.Viper,
) (
	config databaseConfig,
) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s TimeZone=%s dbname=%s sslmode=disable",
		viperViper.GetString("POSTGRES_HOST"),
		viperViper.GetString("POSTGRES_USER"),
		viperViper.GetString("POSTGRES_PASSWORD"),
		viperViper.GetString("POSTGRES_BACK_PORT"),
		viperViper.GetString("TZ"),
		viperViper.GetString("POSTGRES_DB"),
	)

	config = databaseConfig{
		DSN: dsn,
	}
	return
}

func newAuth0Config(
	viperViper *viper.Viper,
) (
	config auth0Config,
) {
	domain := viperViper.GetString("AUTH0_DOMAIN")
	tokenURL := viperViper.GetString("AUTH0_TOKEN_URL")
	if tokenURL == "" && domain != "" {
		tokenURL = fmt.Sprintf("https://%s/oauth/token", domain)
	}

	config = auth0Config{
		Domain:       domain,
		TokenURL:     tokenURL,
		Audience:     viperViper.GetString("AUTH0_AUDIENCE"),
		GrantType:    viperViper.GetString("AUTH0_GRANT_TYPE"),
		ClientSecret: viperViper.GetString("AUTH0_CLIENT_SECRET"),
	}
	return
}

func newPubSubConfig(
	viperViper *viper.Viper,
) (
	config pubSubConfig,
) {
	config = pubSubConfig{
		BootstrapServers: viperViper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
		ConsumerGroupID:  viperViper.GetString("KAFKA_CONSUMER_GROUP_ID"),
		TestTopic:        viperViper.GetString("PUBSUB_TEST_TOPIC"),
		OtherTopic:       viperViper.GetString("PUBSUB_OTHER_TOPIC"),
		FlushTimeoutMS:   viperViper.GetInt("PUBSUB_FLUSH_TIMEOUT_MS"),
		SampleUserName:   viperViper.GetString("PUBSUB_SAMPLE_USER_NAME"),
	}
	return
}

func newRedisConfig(
	viperViper *viper.Viper,
) (
	config redisConfig,
) {
	config = redisConfig{
		Addr:     viperViper.GetString("REDIS_ADDR"),
		Password: viperViper.GetString("REDIS_PASSWORD"),
		DB:       viperViper.GetInt("REDIS_DB"),
	}
	return
}
