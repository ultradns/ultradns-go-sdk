package task_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/test"
	"github.com/ultradns/ultradns-go-sdk/pkg/task"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	secondaryZoneType = "SECONDARY"
)

var (
	testSondaryZoneName = test.GetRandomSecondaryZoneName()
)

func TestNewSuccess(t *testing.T) {
	conf := test.GetConfig()

	if _, err := task.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""

	if _, err := task.New(conf); err.Error() != "config error while creating Task service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetSuccess(t *testing.T) {
	if _, err := task.Get(test.TestClient); err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := task.Get(nil); err.Error() != "Task service is not properly configured" {
		t.Fatal(err)
	}
}

func TestGetTaskStatusWithConfigError(t *testing.T) {
	taskService := task.Service{}

	if _, _, err := taskService.GetTaskStatus(""); err.Error() != "Task service is not properly configured" {
		t.Fatal(err)
	}
}

func TestGetTaskStatusWithInvalidTaskID(t *testing.T) {
	taskService, err := task.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := taskService.GetTaskStatus("a"); er.Error() != "error while getting task status : task id - a : error code : 54001 - error message : Cannot find the task status for the input taskId" {
		t.Fatal(er)
	}
}

func TestTaskWaitError(t *testing.T) {
	taskService, err := task.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if er := taskService.TaskWait("a", 2, 1); er.Error() != "error while getting task status : task id - a : error code : 54001 - error message : Cannot find the task status for the input taskId" {
		t.Fatal(er)
	}
}

func TestTaskWaitTimeoutError(t *testing.T) {
	taskService, err := task.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if er := taskService.TaskWait("a", 0, 0); er.Error() != "timeout for checking task status : last returned task status - <nil>" {
		t.Fatal(er)
	}
}

func TestCreateZoneSuccessWithSecondaryZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	zone := getSecondaryZone()

	if _, er := zoneService.CreateZone(zone); er != nil {
		t.Fatal(er)
	}
}

func TestDeleteSecondaryZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone(testSondaryZoneName); er != nil {
		t.Fatal(er)
	}
}

func getSecondaryZone() *zone.Zone {
	nameServerIP := &zone.NameServer{
		IP: test.TestPrimaryNameServer,
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

	return &zone.Zone{
		Properties:          test.GetZoneProperties(testSondaryZoneName, secondaryZoneType),
		SecondaryCreateInfo: secondaryZone,
	}
}
