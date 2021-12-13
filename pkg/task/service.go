package task

import (
	"net/http"
	"time"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

const (
	serviceName = "Task"
	basePath    = "tasks/"
)

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)
	if err != nil {
		return nil, client.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, client.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) GetTaskStatus(taskID string) (*http.Response, *Task, error) {
	target := client.Target(&Task{})

	if s.c == nil {
		return nil, nil, client.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, basePath+taskID, nil, target)
	if err != nil {
		return nil, nil, StatusError(taskID, err)
	}

	task := target.Data.(*Task)

	return res, task, nil
}

func (s *Service) TaskWait(taskID string, retries, timegap int) error {
	var taskStatus *Task

	for i := 0; i < retries; i++ {
		time.Sleep(time.Duration(timegap) * time.Second)

		_, task, err := s.GetTaskStatus(taskID)
		if err != nil {
			return err
		}

		if task != nil {
			switch task.Code {
			case "COMPLETE":
				return nil
			case "ERROR":
				return FailedTaskError(task)
			}
		}

		taskStatus = task
	}

	return TimeoutError(taskStatus)
}
