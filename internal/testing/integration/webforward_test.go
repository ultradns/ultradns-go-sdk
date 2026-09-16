package integration_test

import (
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/webforward"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func TestWebForwardLifecycle(t *testing.T) {
	if integration.TestClient == nil {
		t.Skip("integration credentials not configured")
	}

	zoneName := integration.GetRandomZoneName()
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := zoneService.CreateZone(integration.GetPrimaryZone(zoneName)); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if _, err := zoneService.DeleteZone(zoneName); err != nil {
			t.Logf("unable to delete zone - %s : error - %s", zoneName, err.Error())
		}
	}()

	webForwardService, err := webforward.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	requestTo := "www." + strings.TrimSuffix(zoneName, ".")
	webForwardData := &webforward.WebForward{
		RequestTo:           requestTo,
		DefaultRedirectTo:   "https://example.com/",
		DefaultForwardType:  webforward.HTTP301Redirect,
		RelativeForwardType: webforward.ParameterAndPath,
	}

	_, created, err := webForwardService.Create(zoneName, webForwardData)

	if err != nil {
		t.Fatal(err)
	}

	if created.GUID == "" {
		t.Fatal("create response missing guid")
	}

	_, read, err := webForwardService.Read(zoneName, created.GUID)

	if err != nil {
		t.Fatal(err)
	}

	if read.RequestTo != requestTo {
		t.Fatalf("requestTo mismatched expected - %v : found - %v", requestTo, read.RequestTo)
	}

	_, res, err := webForwardService.List(zoneName, &helper.QueryInfo{})

	if err != nil {
		t.Fatal(err)
	}

	if len(res.WebForwards) != 1 {
		t.Fatalf("webForward count mismatched expected - 1 : found - %v", len(res.WebForwards))
	}

	webForwardData.DefaultRedirectTo = "https://example.org/"

	if _, err := webForwardService.Update(zoneName, created.GUID, webForwardData); err != nil {
		t.Fatal(err)
	}

	if _, err := webForwardService.Delete(zoneName, created.GUID); err != nil {
		t.Fatal(err)
	}
}
