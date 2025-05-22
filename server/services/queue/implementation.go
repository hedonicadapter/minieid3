package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hedonicadapter/gopher/models"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Rdb  redis.UniversalClient
	List string
}

func (s Service) Enqueue(ctx context.Context, task models.Task) error {
	taskJson, err := json.Marshal(task)
	if err != nil {
		return err
	}

	errr := s.Rdb.LPush(ctx, s.List, string(taskJson)).Err()
	return errr
}

func (s Service) Dequeue(ctx context.Context) ([]string, error) {
	res, err := s.Rdb.BRPop(ctx, 5*time.Second, s.List).Result()
	return res, err
}

func (s Service) Peek() {
	fmt.Println("not implemented")
}

func (s Service) Poll(ctx context.Context, taskHandler func(task models.Task) any) {
	for {
		res, err := s.Dequeue(ctx)
		if err != nil {
			fmt.Println("error dequeueing something: ", err.Error())
			continue
		}
		if len(res) < 2 {
			fmt.Println("Unexpected Dequeue result length")
			continue
		}

		var task models.Task
		if err := json.Unmarshal([]byte(res[1]), &task); err != nil {
			fmt.Println("Error unmarshalling task: ", err.Error())
			continue
		}

		taskHandler(task)
	}
}

func InitService(rdb redis.UniversalClient, list string) *Service {
	return &Service{Rdb: rdb, List: list}
}

// RPUSH queue:${roomId} ${username}: Add user to queue or create queue
// LRANGE queue:${roomId} 0 -1: Get all users in the queue
// LREM queue:${roomId} -1, ${username}: Remove user from queue
// LMOVE queue:${roomId} queue:${roomId} LEFT RIGHT: Cycle through queue (move user from front of queue to back of queue)

// Enqueue
// push something to tail

// Dequeue
// remove something from head

// Peek
// read some element in queue

// Poll
// read head of queue
// Dequeue head, execute head action
