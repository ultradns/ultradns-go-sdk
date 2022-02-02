package task_test

import (
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/task"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const serviceErrorString = "Task service is not properly configured"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := task.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := task.New(conf); err.Error() != "config error while creating Task service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := task.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestGetTaskStatusWithConfigError(t *testing.T) {
	taskService := task.Service{}

	if _, _, err := taskService.GetTaskStatus(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestGetTaskStatusWithInvalidTaskID(t *testing.T) {
	taskService, err := task.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := taskService.GetTaskStatus("a"); er.Error() != "error while getting task status : task id - a : error from api response - error code : 54001 - error message : Cannot find the task status for the input taskId" {
		t.Fatal(er)
	}
}

func TestTaskWaitError(t *testing.T) {
	taskService, err := task.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if er := taskService.TaskWait("a", 2, 1); er.Error() != "error while getting task status : task id - a : error from api response - error code : 54001 - error message : Cannot find the task status for the input taskId" {
		t.Fatal(er)
	}
}

func TestTaskWaitTimeoutError(t *testing.T) {
	taskService, err := task.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if er := taskService.TaskWait("a", 0, 0); er.Error() != "timeout for checking task status : last returned task status - <nil>" {
		t.Fatal(er)
	}
}

func TestFailedTaskWithSecondaryZone(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	nameServerIP := &zone.NameServer{
		IP: integration.TestPrimaryNameServer,
	}
	nameServerIPList := &zone.NameServerIPList{
		NameServerIP1: nameServerIP,
	}

	primaryNameServer := &zone.PrimaryNameServers{
		NameServerIPList: nameServerIPList,
	}

	secondaryZone := &zone.SecondaryZone{
		PrimaryNameServers: primaryNameServer,
	}

	zoneData := &zone.Zone{
		Properties:          integration.GetZoneProperties("non-existing-zone.com.", zone.Secondary),
		SecondaryCreateInfo: secondaryZone,
	}

	if _, er := zoneService.CreateZone(zoneData); !strings.Contains(er.Error(), "is not authoritative for zone 'non-existing-zone.com.'. Please provide correct name server.") {
		t.Fatal(er)
	}
}
