package ultradns_test

import (
	"testing"
	"ultradns-go-sdk/ultradns"
)

func TestDoSuccess(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	type result struct {
		QueryInfo  ultradns.QueryInfo  `json:"queryInfo,omitempty"`
		ResultInfo ultradns.ResultInfo `json:"resultInfo,omitempty"`
	}
	target := ultradns.Target(&result{})
	res, err := testClient.Do("GET", "zones", nil, target)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Not a Successful response : returned response code - %v", res.StatusCode)
	}
	resData := target.Data.(*result)
	if resData.QueryInfo.Limit != 100 {
		t.Errorf("Query limit mismatched : returned limit - %v", resData.QueryInfo.Limit)
	}

	if resData.ResultInfo.ReturnedCount != 100 {
		t.Errorf("Returned count mismatched : returned count - %v", resData.ResultInfo.ReturnedCount)
	}

}

func TestDoNonExistingZone(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}

	target := ultradns.Target(nil)
	res, err := testClient.Do("GET", "zones/errortestingzone", nil, target)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 404 {
		t.Errorf("Not a 404 response : returned response code - %v", res.StatusCode)
	}

	errDataArr := target.Error.(*[]ultradns.ErrorResponse)
	errData := *errDataArr
	if errData[0].ErrorCode != 1801 {
		t.Errorf("Error code mismatch : returned error code - %v", errData[0].ErrorCode)
	}

	if errData[0].ErrorMessage != "Zone does not exist in the system." {
		t.Errorf("Error message mismatch : returned error message - %v", errData[0].ErrorMessage)
	}

}

func TestDoInvalidMethod(t *testing.T) {
	testClient, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do("()", "zones", nil, nil)
	_ = res

	if err.Error() != "net/http: invalid method \"()\"" {
		t.Error(err)
	}
}
