package task

import "fmt"

// Task
type Task struct {
	TaskID    string `json:"taskId,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	ResultURI string `json:"resultUri,omitempty"`
}

func (t Task) String() string {
	return fmt.Sprintf("taskId : %v - code : %v - message : %v", t.TaskID, t.Code, t.Message)
}
