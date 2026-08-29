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

func Load() (Config, error) {
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
	if err := viperViper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("load environment file: %w", err)
	}

	if env == "lcl" {
		if err := setupLocalstack(viperViper); err != nil {
			return Config{}, fmt.Errorf("setup localstack: %w", err)
		}
	}

	return Config{
		Server:   newServerConfig(viperViper),
		Database: newDatabaseConfig(viperViper),
		Auth0:    newAuth0Config(viperViper),
		PubSub:   newPubSubConfig(viperViper),
		Redis:    newRedisConfig(viperViper),
	}, nil
}

func initViper() *viper.Viper {
	viperConfig := viper.New()
	viperConfig.AutomaticEnv()
	viperConfig.AddConfigPath("internal/env")
	viperConfig.SetConfigType("env")
	return viperConfig
}

func setupLocalstack(
	viperViper *viper.Viper,
) error {

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
		return err
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(viperViper.GetString("AWS_SECRET_NAME")),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return err
	}

	resultSecretString := aws.ToString(result.SecretString)
	localstackSecrets := &LocalstackSecrets{}
	if err := json.Unmarshal(
		[]byte(resultSecretString),
		localstackSecrets,
	); err != nil {
		return err
	}

	var secretString SecretString
	if err := json.Unmarshal(
		[]byte(localstackSecrets.MyLocalSecret.SecretString),
		&secretString,
	); err != nil {
		return err
	}

	viperViper.Set("POSTGRES_USER", secretString.Username)
	viperViper.Set("POSTGRES_PASSWORD", secretString.Password)

	return nil
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
) serverConfig {
	return serverConfig{
		BackendHost: viperViper.GetString("BACKEND_HOST"),
		GoEchoPort:  viperViper.GetString("GO_ECHO_PORT"),
		GRPCPort:    viperViper.GetString("GRPC_PORT"),
		GRPCAddress: fmt.Sprintf(
			"%s:%s",
			viperViper.GetString("BACKEND_HOST"),
			viperViper.GetString("GRPC_PORT"),
		),
	}
}

func newDatabaseConfig(
	viperViper *viper.Viper,
) databaseConfig {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s TimeZone=%s dbname=%s sslmode=disable",
		viperViper.GetString("POSTGRES_HOST"),
		viperViper.GetString("POSTGRES_USER"),
		viperViper.GetString("POSTGRES_PASSWORD"),
		viperViper.GetString("POSTGRES_BACK_PORT"),
		viperViper.GetString("TZ"),
		viperViper.GetString("POSTGRES_DB"),
	)

	return databaseConfig{
		DSN: dsn,
	}
}

func newAuth0Config(
	viperViper *viper.Viper,
) auth0Config {
	domain := viperViper.GetString("AUTH0_DOMAIN")
	tokenURL := viperViper.GetString("AUTH0_TOKEN_URL")
	if tokenURL == "" && domain != "" {
		tokenURL = fmt.Sprintf("https://%s/oauth/token", domain)
	}

	return auth0Config{
		Domain:       domain,
		TokenURL:     tokenURL,
		Audience:     viperViper.GetString("AUTH0_AUDIENCE"),
		GrantType:    viperViper.GetString("AUTH0_GRANT_TYPE"),
		ClientSecret: viperViper.GetString("AUTH0_CLIENT_SECRET"),
	}
}

func newPubSubConfig(
	viperViper *viper.Viper,
) pubSubConfig {
	return pubSubConfig{
		BootstrapServers: viperViper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
		ConsumerGroupID:  viperViper.GetString("KAFKA_CONSUMER_GROUP_ID"),
		TestTopic:        viperViper.GetString("PUBSUB_TEST_TOPIC"),
		OtherTopic:       viperViper.GetString("PUBSUB_OTHER_TOPIC"),
		FlushTimeoutMS:   viperViper.GetInt("PUBSUB_FLUSH_TIMEOUT_MS"),
		SampleUserName:   viperViper.GetString("PUBSUB_SAMPLE_USER_NAME"),
	}
}

func newRedisConfig(
	viperViper *viper.Viper,
) redisConfig {
	return redisConfig{
		Addr:     viperViper.GetString("REDIS_ADDR"),
		Password: viperViper.GetString("REDIS_PASSWORD"),
		DB:       viperViper.GetInt("REDIS_DB"),
	}
}
