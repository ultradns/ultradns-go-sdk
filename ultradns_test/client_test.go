/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns_test

import (
	"fmt"
	"os"
	"testing"
	"ultradns-go-sdk/ultradns"
)

var (
	testUsername  = os.Getenv("ULTRADNS_USERNAME")
	testPassword  = os.Getenv("ULTRADNS_PASSWORD")
	testHost      = os.Getenv("ULTRADNS_HOST_URL")
	testVersion   = os.Getenv("ULTRADNS_API_VERSION")
	testUserAgent = os.Getenv("ULTRADNS_USER_AGENT")
	testZoneName  = "go_sdk_unit_testing.com"
)

func TestNewClientWithCredentials(t *testing.T) {
	_, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Error(err)
	}
}

func TestNewClientWithoutUsername(t *testing.T) {
	_, err := ultradns.NewClient("", testPassword, testHost, testVersion, testUserAgent)
	if err.Error() != "User Name is required to create a new http client for UltraDNS rest api" {
		t.Error(err)
	}
}

func TestNewClientWithoutPassword(t *testing.T) {
	_, err := ultradns.NewClient(testUsername, "", testHost, testVersion, testUserAgent)
	if err.Error() != "Password is required to create a new http client for UltraDNS rest api" {
		t.Error(err)
	}
}

func TestNewClientWithoutHost(t *testing.T) {
	_, err := ultradns.NewClient(testUsername, testPassword, "", testVersion, testUserAgent)
	if err.Error() != "Host Url is required to create a new http client for UltraDNS rest api" {
		t.Error(err)
	}
}

func TestNewClientWithoutApiVersion(t *testing.T) {
	_, err := ultradns.NewClient(testUsername, testPassword, testHost, "", testUserAgent)
	if err.Error() != "Api Version is required to create a new http client for UltraDNS rest api" {
		t.Error(err)
	}
}

func TestNewClientWithoutUserAgent(t *testing.T) {
	_, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, "")
	if err.Error() != "User Agent is required to create a new http client for UltraDNS rest api" {
		t.Error(err)
	}
}

func TestNewClientWithWrongSchema(t *testing.T) {
	_, err := ultradns.NewClient(testUsername, testPassword, ":", testVersion, testUserAgent)
	if err.Error() != fmt.Sprintf("parse \":%v\": missing protocol scheme", testVersion) {
		t.Error(err)
	}
}
