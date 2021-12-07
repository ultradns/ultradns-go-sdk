package task

import (
	"errors"
	"fmt"
)

var (
	errTaskTimeout = errors.New("timeout for checking task status")
	errTaskFailed  = errors.New("task failed")
)

func TimeoutError(t *Task) error {
	return fmt.Errorf("%w : last returned task status - %s ", errTaskTimeout, t)
}

func FailedTaskError(t *Task) error {
	return fmt.Errorf("%w : %s", errTaskFailed, t)
}

func StatusError(taskID string, err error) error {
	return fmt.Errorf("error while getting task status : task id - %s : %w", taskID, err)
}
