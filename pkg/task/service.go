package task

import (
	"net/http"
	"time"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
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
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) GetTaskStatus(taskID string) (*http.Response, *Task, error) {
	target := client.Target(&Task{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s read started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, basePath+taskID, nil, target)
	if err != nil {
		s.c.Error("[%s] %s read failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ReadError(serviceName, taskID, err)
	}

	task := target.Data.(*Task)

	s.c.Trace("[%s] %s read completed successfully", version.GetSDKVersion(), serviceName)

	return res, task, nil
}

func (s *Service) TaskWait(taskID string, retries, timegap int) error {
	var taskStatus *Task

	for i := 0; i < retries; i++ {
		s.c.Trace("[%s] sleeping for %d seconds", version.GetSDKVersion(), timegap)
		time.Sleep(time.Duration(timegap) * time.Second)

		_, task, err := s.GetTaskStatus(taskID)
		if err != nil {
			return err
		}

		if task != nil {
			switch task.Code {
			case "COMPLETE":
				s.c.Trace("[%s] %s completed with status: 'COMPLETE'", version.GetSDKVersion(), serviceName)
				return nil
			case "ERROR":
				s.c.Trace("[%s] %s completed with status: 'ERROR'", version.GetSDKVersion(), serviceName)
				return FailedTaskError(task)
			}
		}

		taskStatus = task
	}

	s.c.Error("[%s] %s check timed out", version.GetSDKVersion(), serviceName)

	return TimeoutError(taskStatus)
}
