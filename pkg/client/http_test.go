package client_test

import (
	"net/http"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func TestDoSuccess(t *testing.T) {
	target := client.Target(&zone.ResponseList{})
	res, err := integration.TestClient.Do(http.MethodGet, "zones", nil, target)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(target.Error[0].String())
	}
}

func TestDoNilTarget(t *testing.T) {
	_, err := integration.TestClient.Do(http.MethodGet, "zones", nil, nil)

	if err.Error() != "response target type mismatched : returned type - <nil>" {
		t.Fatal(err)
	}
}

func TestDoWrongTarget(t *testing.T) {
	_, err := integration.TestClient.Do(http.MethodGet, "zones", nil, &zone.Zone{})

	if err.Error() != "response target type mismatched : returned type - *client.Response" {
		t.Fatal(err)
	}
}

func TestDoNonExistingZone(t *testing.T) {
	target := client.Target(&zone.Response{})
	_, err := integration.TestClient.Do(http.MethodGet, "zones/unit-test-non-existing-zone.com", nil, target)

	if err.Error() != "error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Error(err)
	}
}

func TestDoInvalidMethod(t *testing.T) {
	_, err := integration.TestClient.Do("()", "zones", nil, nil)

	if err.Error() != "net/http: invalid method \"()\"" {
		t.Error(err)
	}
}
