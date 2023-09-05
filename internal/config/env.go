package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gitlab.com/4nd3rs0n/errorsx"
)

type EnvErrInfo struct {
	VarName string
}
type EnvErr = errorsx.CustomErr[EnvErrInfo]

var VarCannotBeEmpty = errors.New("Variable cannot be empty")

const (
	defaultPgUser = "postgres"
	defaultPgPort = 5432
)
const (
	defaultRedisUser = "default"
	defaultRedisPort = 6379
	defaultRedisDB   = 0
)

func getEnvWithHandler[T any](envName string, handler func(env string) (T, error)) (T, EnvErr) {
	env := os.Getenv(envName)
	out, err := handler(env)
	if err != nil {
		return out, errorsx.NewCustomErr(err.Error(), EnvErrInfo{
			VarName: envName,
		})
	}
	return out, nil
}

func getEnvNotEmptyStr(envName string) (string, EnvErr) {
	if env := os.Getenv(envName); env == "" {
		return "", errorsx.NewCustomErr(VarCannotBeEmpty.Error(), EnvErrInfo{
			VarName: envName,
		})
	} else {
		return env, nil
	}
}

func GetPgUrlFromEnv() (string, EnvErr) {
	user, errx := getEnvWithHandler("POSTGRES_USER", func(env string) (string, error) {
		if env == "" {
			return "postgres", nil
		}
		return env, nil
	})
	if errx != nil {
		return "", errx
	}
	password, errx := getEnvNotEmptyStr("POSTGRES_PASSWORD")
	if errx != nil {
		return "", errx
	}
	host, errx := getEnvNotEmptyStr("POSTGRES_HOST")
	if errx != nil {
		return "", errx
	}
	port, errx := getEnvWithHandler("POSTGRES_PORT", func(env string) (int, error) {
		if env == "" {
			return defaultPgPort, nil
		}
		return strconv.Atoi(env)
	})
	if errx != nil {
		return "", errx
	}
	db, errx := getEnvNotEmptyStr("POSTGRES_DB")
	if errx != nil {
		return "", errx
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		user, password, host, port, db), nil
}

func GetRedisOptionsFromEnv() (*redis.Options, EnvErr) {
	host, errx := getEnvNotEmptyStr("REDIS_HOST")
	if errx != nil {
		return &redis.Options{}, errx
	}
	port, errx := getEnvWithHandler("REDIS_PORT", func(env string) (int, error) {
		if env == "" {
			return defaultRedisPort, nil
		}
		return strconv.Atoi(env)
	})
	if errx != nil {
		return &redis.Options{}, errx
	}

	username, errx := getEnvWithHandler("REDIS_USER", func(env string) (string, error) {
		if env == "" {
			return defaultRedisUser, nil
		}
		return env, nil
	})
	password, errx := getEnvNotEmptyStr("REDIS_PASSWORD")
	if errx != nil {
		return &redis.Options{}, errx
	}
	db, errx := getEnvWithHandler("REDIS_DB", func(env string) (int, error) {
		if env == "" {
			return defaultRedisDB, nil
		}
		return strconv.Atoi(env)
	})

	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: username,
		Password: password,
		DB:       db,
	}, nil
}

func GetSmtpCfgFromEnv() (SMTPConfig, EnvErr) {
	from, errx := getEnvNotEmptyStr("APP_MAIL_ADDRESS")
	if errx != nil {
		return SMTPConfig{}, errx
	}
	server, errx := getEnvNotEmptyStr("APP_SMTP_SERV_ADDRESS")
	if errx != nil {
		return SMTPConfig{}, errx
	}
	password, errx := getEnvNotEmptyStr("APP_MAIL_PASSWORD")
	if errx != nil {
		return SMTPConfig{}, errx
	}

	return SMTPConfig{
		From:         from,
		SMTPServer:   server,
		SMTPPassword: password,
	}, nil
}
