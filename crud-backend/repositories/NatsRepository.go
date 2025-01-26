package repositories

import (
	"encoding/json"

	"crud-backend/models"

	"github.com/nats-io/nats.go"
)

type NatsRepository struct {
	Conn *nats.Conn
}

func (r *NatsRepository) PublishTask(task *models.Task) error {
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = r.Conn.Publish("new_task", taskData)

	return err
}
