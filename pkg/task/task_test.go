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

	if err.Error() != "password is required to create a client" {
		t.Fatal(err)
	}
}
