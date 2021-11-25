/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns

import (
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	TaskId    string `json:"taskId,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	ResultUri string `json:"resultUri,omitempty"`
}

type TaskList struct {
	QueryInfo  *QueryInfo  `json:"queryInfo"`
	ResultInfo *ResultInfo `json:"resultInfo"`
	Tasks      *[]Task     `json:"tasks"`
}

func (t Task) String() string {
	return fmt.Sprintf("taskId : %v - code : %v - message : %v", t.TaskId, t.Code, t.Message)
}

func (c *Client) GetTaskStatus(taskId string) (*http.Response, *Task, error) {
	target := Target(&Task{})
	res, err := c.Do("GET", "tasks/"+taskId, nil, target)

	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errDataListPtr := target.Error.(*[]ErrorResponse)
		errDataList := *errDataListPtr
		return res, nil, fmt.Errorf("error while getting task status - %s", errDataList[0])
	}

	task := target.Data.(*Task)

	return res, task, nil
}

func (c *Client) TaskWait(taskId string, retries, timegap int) error {
	var taskStatus *Task
	for i := 0; i < retries; i++ {
		time.Sleep(time.Duration(timegap) * time.Second)
		_, task, err := c.GetTaskStatus(taskId)

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
