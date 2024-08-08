package task

import (
	"errors"
	"fmt"
)

var (
	errTaskTimeout = errors.New("Timeout task status check")
	errTaskFailed  = errors.New("Task failed")
)

func TimeoutError(t *Task) error {
	return fmt.Errorf("%w: { status: '%s' }", errTaskTimeout, t)
}

func FailedTaskError(t *Task) error {
	return fmt.Errorf("%w: %s", errTaskFailed, t)
}
