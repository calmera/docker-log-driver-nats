package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"

	"github.com/docker/go-plugins-helpers/sdk"
	"github.com/sirupsen/logrus"
)

var logLevels = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

const (
	NatsClientIdVar  = "NATS_CLIENT_ID"
	NatsUrlVar       = "NATS_URL"
	UserJwtVar       = "NATS_USER_JWT"
	UserSeedVar      = "NATS_USER_SEED"
	SubjectPrefixVar = "NATS_SUBJECT_PREFIX"
)

func main() {
	levelVal := os.Getenv("LOG_LEVEL")
	if levelVal == "" {
		levelVal = "info"
	}
	if level, exists := logLevels[levelVal]; exists {
		logrus.SetLevel(level)
	} else {
		fmt.Fprintln(os.Stderr, "invalid log level: ", levelVal)
		os.Exit(1)
	}
	cfg, err := ConfigFromEnv()
	if err != nil {
		logrus.Fatal(err)
	}

	nc, err := nats.Connect(cfg.Url, cfg.Options...)
	if err != nil {
		logrus.Fatal(err)
	}

	h := sdk.NewHandler(`{"Implements": ["LogDriver"]}`)
	handlers(&h, newDriver(nc, cfg))
	if err := h.ServeUnix("nats-logger", 0); err != nil {
		panic(err)
	}
}
