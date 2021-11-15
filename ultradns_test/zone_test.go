/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns_test

import (
	"testing"
	"ultradns-go-sdk/ultradns"
)

func TestCreateZone(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        "go_sdk_unit_testing.com",
		AccountName: testUsername,
		Type:        "PRIMARY",
	}
	PrimaryZone := ultradns.PrimaryZone{
		CreateType: "NEW",
	}
	zone := ultradns.Zone{
		Properties:        &zoneProp,
		PrimaryCreateInfo: &PrimaryZone,
	}

	res, err := testClient.CreateZone(zone, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 201 {
		t.Errorf("Zone not created : returned response code - %v", res.StatusCode)
	}
}

func TestReadZone(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	type result struct {
		Properties ultradns.ZoneProperties `json:"properties,omitempty"`
	}
	target := ultradns.Target(&result{})
	res, err := testClient.ReadZone("go_sdk_unit_testing.com", target)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Not a Successful response : returned response code - %v", res.StatusCode)
	}
	resData := target.Data.(*result)
	if resData.Properties.Name != "go_sdk_unit_testing.com." {
		t.Errorf("Zone name mismatched : returned zone name - %v", resData.Properties.Name)
	}

	if resData.Properties.Type != "PRIMARY" {
		t.Errorf("Zone type mismatched : returned zone type - %v", resData.Properties.Type)
	}

	if resData.Properties.Status != "ACTIVE" {
		t.Errorf("Zone status not active : returned zone status - %v", resData.Properties.Status)
	}

	if resData.Properties.AccountName != testUsername {
		t.Errorf("Zone account name mismatched : returned account name - %v", resData.Properties.AccountName)
	}

}

func TestUpdateZone(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        "go_sdk_unit_testing.com",
		AccountName: testUsername,
		Type:        "PRIMARY",
	}
	PrimaryZone := ultradns.PrimaryZone{
		CreateType: "NEW",
	}
	zone := ultradns.Zone{
		Properties:        &zoneProp,
		PrimaryCreateInfo: &PrimaryZone,
	}

	res, err := testClient.UpdateZone("go_sdk_unit_testing.com", zone, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Zone not updated : returned response code - %v", res.StatusCode)
	}
}

func TestDeleteZone(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.DeleteZone("go_sdk_unit_testing.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 204 {
		t.Errorf("Zone not Deleted : returned response code - %v", res.StatusCode)
	}

}
