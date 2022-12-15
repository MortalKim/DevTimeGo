package task

import (
	"WakaTImeGo/service/duration"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)
import "github.com/hibiken/asynq"

/**
 * @Author: Kim
 * @Description:
 * @File:  durationGen
 * @Date: 12/13/2022 8:05 PM
 */

const (
	TypeDurationDelivery = "duration:deliver"
)

type DurationDeliveryPayload struct {
	UserID string
}

func NewDurationDeliveryTask(userID string) (*asynq.Task, error) {
	payload, err := json.Marshal(DurationDeliveryPayload{UserID: userID})
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(TypeDurationDelivery, payload)
	GetTaskClient().Enqueue(task)
	return task, nil
}

func HandleDurationDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p DurationDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Generic user's duration userid= %s", p.UserID)
	duration.GenericUserDuration(p.UserID)
	return nil
}
