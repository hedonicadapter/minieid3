package queue

import (
	"context"

	"github.com/hedonicadapter/gopher/models"
)

type QueueService interface {
	Enqueue(ctx context.Context, task models.Task) error
	Dequeue(ctx context.Context) ([]string, error)
	Peek()
	Poll(ctx context.Context, taskHandler func(task models.Task) any)
}
