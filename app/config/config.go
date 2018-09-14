package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	ENV        string
	PORT       string
	JWT_SECRET string
	CORS       string

	DB_TYPE     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string

	RABBITMQ_USER     string
	RABBITMQ_PASSWORD string
	RABBITMQ_HOST     string
	RABBITMQ_PORT     string
	RABBIT_COMMENT_COMMENT_QUEUE string
	RABBIT_POST_COMMENT_QUEUE string
	RABBIT_POST_VOTE_QUEUE string
	RABBIT_NOTIFICATION_EXCHANGE string

	WORKER_AMOUNT string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		err := readEnv()
		if err != nil {
			panic(err)
		}
		config := newConfig()
		instance = &config
	}
	return instance
}

func newConfig() Config {
	return Config{
		ENV:  GetEnv("ENV", "develop"),
		PORT: GetEnv("PORT", "5004"),

		DB_TYPE:     GetEnv("DB_TYPE", "mysql"),
		DB_USERNAME: GetEnv("DB_USERNAME", "root"),
		DB_PASSWORD: GetEnv("DB_PASSWORD", "root"),
		DB_HOST:     GetEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:     GetEnv("DB_PORT", "3306"),
		DB_NAME:     GetEnv("DB_NAME", "amazing-code-database"),

		RABBITMQ_HOST:     GetEnv("RABBITMQ_HOST", "localhost"),
		RABBITMQ_PORT:     GetEnv("RABBITMQ_PORT", "5672"),
		RABBITMQ_USER:     GetEnv("RABBITMQ_USER", "guest"),
		RABBITMQ_PASSWORD: GetEnv("RABBITMQ_PASSWORD", "guest"),
		RABBIT_COMMENT_COMMENT_QUEUE: GetEnv("RABBIT_COMMENT_COMMENT_QUEUE", "comment_comment_queue"),
		RABBIT_POST_VOTE_QUEUE: GetEnv("RABBIT_POST_VOTE_QUEUE", "post_vote_queue"),
		RABBIT_POST_COMMENT_QUEUE: GetEnv("RABBIT_POST_COMMENT_QUEUE", "post_comment_queue"),
		RABBIT_NOTIFICATION_EXCHANGE: GetEnv("RABBIT_NOTIFICATION_EXCHANGE", "notifications"),

		WORKER_AMOUNT: GetEnv("WORKER_AMOUNT", "3"),
	}
}

func GetEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}

func readEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "=")
		if len(values) == 2 {
			err = os.Setenv(values[0], values[1])
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
