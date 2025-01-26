package nats

import (
	"os"

	"github.com/nats-io/nats.go"
)

func InitNats() (*nats.Conn, error) {
	url := os.Getenv("NATS_URL")

	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
