package main

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/daemon/logger"
	"github.com/nats-io/nats.go"
)

const Name = "nats"

func NewNatsLogger(nc *nats.Conn, subjectPrefix string, info logger.Info) logger.Logger {
	return &NatsLogger{
		nc:      nc,
		subject: fmt.Sprintf("%s.daemons.%s.containers.%s", subjectPrefix, nc.Opts.Name, info.ContainerID),
	}
}

type NatsLogger struct {
	nc      *nats.Conn
	subject string
}

func (n NatsLogger) Log(message *logger.Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return n.nc.Publish(n.subject, b)
}

func (n NatsLogger) Name() string {
	return Name
}

func (n NatsLogger) Close() error {
	return nil
}
