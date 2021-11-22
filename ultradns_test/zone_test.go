/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns_test

import (
	"fmt"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func TestCreateZoneSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
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

	res, err := testClient.CreateZone(zone)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 201 {
		t.Errorf("Zone not created : returned response code - %v", res.StatusCode)
	}
}

func TestCreateZoneFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, "testHost", testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
		AccountName: testUsername,
		Type:        "PRIMARY",
	}
	zone := ultradns.Zone{
		Properties: &zoneProp,
	}

	_, er := testClient.CreateZone(zone)

	if er.Error() != "Post \"testHostv2/zones\": Post \"testHostv2/authorization/token\": unsupported protocol scheme \"\"" {
		t.Error(er)
	}
}

func TestCreateZoneFailureResponse(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
		AccountName: testUsername,
		Type:        "PRIMARY",
	}
	zone := ultradns.Zone{
		Properties: &zoneProp,
	}

	_, er := testClient.CreateZone(zone)

	if er.Error() != fmt.Sprintf("error while creating a zone (%v) - error code : 55001 - error message : zone.primaryCreateInfo is required field.", testZoneName) {
		t.Error(er)
	}

}

func TestReadZoneSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	res, zoneType, zoneResponse, er := testClient.ReadZone(testZoneName)
	if er != nil {
		t.Fatal(er)
	}
	if res.StatusCode != 200 {
		t.Errorf("Not a Successful response : returned response code - %v", res.StatusCode)
	}

	if zoneResponse.Properties.Name != testZoneName {
		t.Errorf("Zone name mismatched expected - %v : returned zone name - %v", testZoneName, zoneResponse.Properties.Name)
	}

	if zoneType != "PRIMARY" {
		t.Errorf("Zone type mismatched expected - PRIMARY : returned zone type - %v", zoneType)
	}

	if zoneResponse.Properties.Status != "ACTIVE" {
		t.Errorf("Zone status not active : returned zone status - %v", zoneResponse.Properties.Status)
	}

	if zoneResponse.Properties.AccountName != testUsername {
		t.Errorf("Zone account name mismatched : returned account name - %v", zoneResponse.Properties.AccountName)
	}

}

func TestReadZoneFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, "testHost", testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _, er := testClient.ReadZone(testZoneName)

	if er.Error() != fmt.Sprintf("Get \"testHostv2/zones/%v\": Post \"testHostv2/authorization/token\": unsupported protocol scheme \"\"", testZoneName) {
		t.Error(er)
	}
}

func TestUpdateZoneSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
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

	res, err := testClient.UpdateZone(testZoneName, zone)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Zone not updated : returned response code - %v", res.StatusCode)
	}
}

func TestUpdateZoneFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, "testHost", testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
		AccountName: testUsername,
		Type:        "PRIMARY",
	}
	zone := ultradns.Zone{
		Properties: &zoneProp,
	}

	_, er := testClient.UpdateZone(testZoneName, zone)

	if er.Error() != fmt.Sprintf("Put \"testHostv2/zones/%v\": Post \"testHostv2/authorization/token\": unsupported protocol scheme \"\"", testZoneName) {
		t.Error(er)
	}
}

func TestUpdateZoneFailureResponse(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	zoneProp := ultradns.ZoneProperties{
		Name:        testZoneName,
		AccountName: testUsername,
		Type:        "PRIMARY",
	}
	zone := ultradns.Zone{
		Properties: &zoneProp,
	}

	_, er := testClient.UpdateZone(testZoneName, zone)

	if er.Error() != fmt.Sprintf("error while updating a zone (%v) - error code : 55001 - error message : zone.primaryCreateInfo is required field.", testZoneName) {
		t.Error(er)
	}
}

func TestDeleteZoneSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.DeleteZone("go_sdk_unit_testing.com")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 204 {
		t.Errorf("Zone not Deleted : returned response code - %v", res.StatusCode)
	}

}

func TestDeleteZoneFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, "testHost", testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, er := testClient.DeleteZone("errortestingzone")

	if er.Error() != "Delete \"testHostv2/zones/errortestingzone\": Post \"testHostv2/authorization/token\": unsupported protocol scheme \"\"" {
		t.Error(er)
	}

}

func TestDeleteZoneFailureResponse(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, er := testClient.DeleteZone("errortestingzone")

	if er.Error() != "error while Deleting a zone (errortestingzone) - error code : 1801 - error message : Zone does not exist in the system." {
		t.Error(er)
	}

}

func TestReadZoneFailureResponse(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _, er := testClient.ReadZone(testZoneName)
	if er.Error() != fmt.Sprintf("error while reading a zone (%v) - error code : 1801 - error message : Zone does not exist in the system.", testZoneName) {
		t.Error(er)
	}
}

func TestListZoneSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}

	_, zoneListResponse, err := testClient.ListZone("?&limit=1")

	if err != nil {
		t.Fatal(err)
	}

	if zoneListResponse.ResultInfo.ReturnedCount != 1 {
		t.Errorf("zone returned count mismatched expected : 1 - returned count : %v", zoneListResponse.ResultInfo.ReturnedCount)
	}
}

func TestListZoneFailure(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, "testHost", testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, _, er := testClient.ListZone("?&limit=1")

	if er.Error() != "Get \"testHostv2/zones/?&limit=1\": Post \"testHostv2/authorization/token\": unsupported protocol scheme \"\"" {
		t.Error(er)
	}
}

func TestListZoneFailureResponse(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	_, _, er := testClient.ListZone("?&limit=k")

	if er.Error() != "error while listing zones - error code : 400 - error message : Invalid value for query parameter 'limit'" {
		t.Error(er)
	}
}
