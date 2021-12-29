package client

import (
	"errors"
	"fmt"
)

var (
	errConfig         = errors.New("config validation failure")
	errService        = errors.New("service is not properly configured")
	errResponseTarget = errors.New("response target type mismatched : returned type")
)

func ConfigError(key string) error {
	return fmt.Errorf("%w: %s is missing", errConfig, key)
}

func ServiceError(service string) error {
	return fmt.Errorf("%s %w", service, errService)
}

func ServiceConfigError(service string, err error) error {
	return fmt.Errorf("config error while creating %s service : %w", service, err)
}

func ResponseTargetError(key string) error {
	return fmt.Errorf("%w - %s", errResponseTarget, key)
}

func ResponseError(errResponseList []*ErrorResponse) error {
	return fmt.Errorf("%s", errResponseList[0])
}
