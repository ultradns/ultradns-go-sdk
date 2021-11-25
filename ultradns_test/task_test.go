/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns_test

import (
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func TestGetTaskStatusFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, "testHost", testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, _, er := testClient.GetTaskStatus("")

	if er.Error() != "Get \"testHostv2/tasks/\": Post \"testHostv2/authorization/token\": unsupported protocol scheme \"\"" {
		t.Error(er)
	}
}

func TestGetTaskStatusFailureResponse(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, _, er := testClient.GetTaskStatus("a")

	if er.Error() != "error while getting task status - error code : 54001 - error message : Cannot find the task status for the input taskId" {
		t.Error(er)
	}
}

func TestZoneTaskWaitFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}

	er := testClient.TaskWait("a", 1, 10)

	if er.Error() != "error while getting task status - error code : 54001 - error message : Cannot find the task status for the input taskId" {
		t.Error(er)
	}
}

func TestZoneTaskWaitCreatingSecondaryZoneFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}

	defer testClient.DeleteZone(testZoneName)
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
		AccountName: testUsername,
		Type:        "SECONDARY",
	}

	nameserverIp := ultradns.NameServerIp{
		Ip: "ultradn.net.",
	}
	nameserverIpList := ultradns.NameServerIpList{
		NameServerIp1: &nameserverIp,
	}
	primaryNameServer := ultradns.PrimaryNameServers{
		NameServerIpList: &nameserverIpList,
	}
	secondaryZone := ultradns.SecondaryZone{
		PrimaryNameServers:  &primaryNameServer,
		AllowUnResponsiveNs: true,
	}
	zone := ultradns.Zone{
		Properties:          &zoneProp,
		SecondaryCreateInfo: &secondaryZone,
	}

	_, er := testClient.CreateZone(zone)

	if !strings.Contains(er.Error(), "code : ERROR") {
		t.Fatal(er)
	}

}

func testZoneTaskWaitCreatingSecondaryZoneSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}

	defer testClient.DeleteZone(testZoneName)
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
		AccountName: testUsername,
		Type:        "SECONDARY",
	}

	nameserverIp := ultradns.NameServerIp{
		Ip: "ultradn.net.",
	}
	nameserverIpList := ultradns.NameServerIpList{
		NameServerIp1: &nameserverIp,
	}
	primaryNameServer := ultradns.PrimaryNameServers{
		NameServerIpList: &nameserverIpList,
	}
	secondaryZone := ultradns.SecondaryZone{
		PrimaryNameServers:  &primaryNameServer,
		AllowUnResponsiveNs: true,
	}
	zone := ultradns.Zone{
		Properties:          &zoneProp,
		SecondaryCreateInfo: &secondaryZone,
	}

	res, er := testClient.CreateZone(zone)

	if er != nil {
		t.Fatal(er)
	}

	if res.StatusCode != 202 {
		t.Errorf("Zone not created : returned response code - %v", res.StatusCode)
	}

}
