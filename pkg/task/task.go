package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

type TaskService struct {
	c *client.Client
}

func New(config client.Config) (*TaskService, error) {
	client, err := client.NewClient(config)

	if err != nil {
		return nil, err
	}
	return &TaskService{c: client}, nil
}

func Get(client *client.Client) (*TaskService, error) {
	if client == nil {
		return nil, fmt.Errorf("task service is not properly configured")
	}
	return &TaskService{c: client}, nil
}

func (ts *TaskService) GetTaskStatus(taskID string) (*http.Response, *Task, error) {
	target := client.Target(&Task{})

	if ts.c == nil {
		return nil, nil, fmt.Errorf("task service is not properly configured")
	}

	res, err := ts.c.Do(http.MethodGet, "tasks/"+taskID, nil, target)

	if err != nil {
		return nil, nil, err
	}

	task := target.Data.(*Task)

	return res, task, nil
}

func (ts *TaskService) TaskWait(taskID string, retries, timegap int) error {
	var taskStatus *Task
	for i := 0; i < retries; i++ {
		time.Sleep(time.Duration(timegap) * time.Second)
		_, task, err := ts.GetTaskStatus(taskID)

		if err != nil {
			return err
		}

		if task != nil {
			switch task.Code {
			case "COMPLETE":
				return nil
			case "ERROR":
				return fmt.Errorf("error - %s", task)
			}
		}
		taskStatus = task
	}
	return fmt.Errorf("timeout for checking task status - last returned task status : %s", taskStatus)
}
