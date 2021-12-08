package task_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/task"
	"github.com/ultradns/ultradns-go-sdk/pkg/test"
)

func TestNewSuccess(t *testing.T) {
	conf := test.GetConfig()
	_, err := task.New(conf)

	if err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""
	_, err := task.New(conf)

	if err.Error() != "config error while creating Task service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetSuccess(t *testing.T) {
	_, err := task.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	_, err := task.Get(nil)

	if err.Error() != "Task service is not properly configured" {
		t.Fatal(err)
	}
}
