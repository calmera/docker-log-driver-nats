package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
)

type Config struct {
	Name          string
	Url           string
	Options       []nats.Option
	SubjectPrefix string
}

func ConfigFromEnv() (*Config, error) {
	result := &Config{}

	result.Name = os.Getenv(NatsClientIdVar)
	if result.Name == "" {
		return nil, fmt.Errorf("%s is required", NatsClientIdVar)
	}
	result.Options = append(result.Options, nats.Name(result.Name))

	result.Url = os.Getenv(NatsUrlVar)
	if result.Url == "" {
		return nil, fmt.Errorf("%s is required", NatsUrlVar)
	}

	jwt := os.Getenv(UserJwtVar)
	seed := os.Getenv(UserSeedVar)
	if jwt != "" && seed != "" {
		result.Options = append(result.Options, nats.UserJWTAndSeed(jwt, seed))
	}

	result.SubjectPrefix = os.Getenv(SubjectPrefixVar)

	return result, nil
}
