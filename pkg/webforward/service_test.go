package webforward_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

const serviceErrorString = "WebForward service configuration failed"

const serverErrorString = "Server error Response - { code: '500', message: 'Internal Server Error' }"

func TestNewSuccess(t *testing.T) {
	if integration.TestUsername == "" {
		t.Skip("integration credentials not configured")
	}

	conf := integration.GetConfig()

	if _, err := webforward.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := webforward.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, _, err := webForwardService.Create("", &webforward.WebForward{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, _, err := webForwardService.Read("", ""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, err := webForwardService.Update("", "", &webforward.WebForward{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, err := webForwardService.Delete("", ""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, _, err := webForwardService.List("", &helper.QueryInfo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestWebForwardCRUD(t *testing.T) {
	const (
		zoneName = "example.com."
		guid     = "09084E85FE7773E2"
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/authorization/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
		case r.Method == http.MethodPost && r.URL.Path == "/zones/example.com./webforwards":
			var payload webforward.WebForward
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				t.Errorf("decode create payload: %v", err)
			}
			if payload.RequestTo != "www.example.com" ||
				payload.DefaultForwardType != webforward.HTTP301Redirect {
				t.Errorf("unexpected create payload: %+v", payload)
			}
			w.WriteHeader(http.StatusCreated)
			_, _ = io.WriteString(w, `{"guid":"`+guid+`"}`)
		case r.Method == http.MethodGet && r.URL.Path == "/zones/example.com./webforwards":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"webForwards":[{"guid":"`+guid+`","requestTo":"http://www.example.com","defaultRedirectTo":"https://target.example.com","defaultForwardType":"HTTP_301_REDIRECT","redirectType":"HTTP"}],"resultInfo":{"totalCount":1,"offset":0,"returnedCount":1}}`)
		case r.Method == http.MethodPut && r.URL.Path == "/zones/example.com./webforwards/"+guid:
			var payload webforward.WebForward
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				t.Errorf("decode update payload: %v", err)
			}
			if payload.GUID != guid {
				t.Errorf("update payload must carry the guid: %+v", payload)
			}
			if payload.DefaultRedirectTo != "https://updated.example.com" {
				t.Errorf("unexpected update payload: %+v", payload)
			}
			w.WriteHeader(http.StatusOK)
		case r.Method == http.MethodDelete && r.URL.Path == "/zones/example.com./webforwards/"+guid:
			w.WriteHeader(http.StatusNoContent)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	service, err := webforward.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	payload := &webforward.WebForward{
		RequestTo:          "www.example.com",
		DefaultRedirectTo:  "https://target.example.com",
		DefaultForwardType: webforward.HTTP301Redirect,
	}

	if _, created, err := service.Create(zoneName, payload); err != nil {
		t.Fatal(err)
	} else if created.GUID != guid {
		t.Fatalf("guid mismatched expected - %s : found - %s", guid, created.GUID)
	}

	if _, read, err := service.Read(zoneName, guid); err != nil {
		t.Fatal(err)
	} else if read.RedirectType != "HTTP" {
		t.Fatalf("redirectType mismatched expected - HTTP : found - %s", read.RedirectType)
	}

	updatePayload := &webforward.WebForward{
		RequestTo:          "www.example.com",
		DefaultRedirectTo:  "https://updated.example.com",
		DefaultForwardType: webforward.HTTP302Redirect,
	}

	if _, err := service.Update(zoneName, guid, updatePayload); err != nil {
		t.Fatal(err)
	}
	if updatePayload.GUID != "" {
		t.Fatalf("update should not mutate the caller payload: %+v", updatePayload)
	}

	if _, err := service.Delete(zoneName, guid); err != nil {
		t.Fatal(err)
	}
}

func TestReadFollowsResultInfoPages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/authorization/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
		case r.Method == http.MethodGet && r.URL.Path == "/zones/example.com./webforwards":
			w.Header().Set("Content-Type", "application/json")
			if r.URL.RawQuery == "&q=&offset=1&cursor=&limit=100&sort=&reverse=false" {
				_, _ = io.WriteString(w, `{"webForwards":[{"guid":"09084E85FE7773E2","requestTo":"api.example.com"}],"resultInfo":{"totalCount":2,"offset":1,"returnedCount":1}}`)
				return
			}
			_, _ = io.WriteString(w, `{"webForwards":[{"guid":"0908486E5BFC7DE1","requestTo":"www.example.com"}],"resultInfo":{"totalCount":2,"offset":0,"returnedCount":1}}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	service, err := webforward.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, read, err := service.Read("example.com.", "09084E85FE7773E2"); err != nil {
		t.Fatal(err)
	} else if read.RequestTo != "api.example.com" {
		t.Fatalf("requestTo mismatched expected - api.example.com : found - %s", read.RequestTo)
	}
}

func newServerErrorTestService(t *testing.T) *webforward.Service {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/authorization/token" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, `{"errorCode":500,"errorMessage":"Internal Server Error"}`)
	}))
	t.Cleanup(server.Close)

	service, err := webforward.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})

	if err != nil {
		t.Fatal(err)
	}

	return service
}

func TestCreateWithServerError(t *testing.T) {
	service := newServerErrorTestService(t)

	_, _, err := service.Create("example.com.", &webforward.WebForward{})

	expected := "Error while creating WebForward: " + serverErrorString + ": {key: 'example.com.'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("create error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestReadWithListServerError(t *testing.T) {
	service := newServerErrorTestService(t)

	_, _, err := service.Read("example.com.", "09084E85FE7773E2")

	expected := "Error while reading WebForward: Error while listing WebForward: " + serverErrorString + ": {key: 'zones/example.com./webforwards?&q=&offset=0&cursor=&limit=100&sort=&reverse=false'}: {key: '09084E85FE7773E2'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("read error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestReadWithZoneNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/authorization/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
		case r.Method == http.MethodGet && r.URL.Path == "/zones/example.com./webforwards":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.WriteString(w, `{"errorCode":404,"errorMessage":"Data Not Found"}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	service, err := webforward.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = service.Read("example.com.", "09084E85FE7773E2")

	expected := "Resource not found: { name: 'WebForward', type: 'webForward', key:'09084E85FE7773E2'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("read not found error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestReadWithWebForwardNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/authorization/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
		case r.Method == http.MethodGet && r.URL.Path == "/zones/example.com./webforwards":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"webForwards":[{"guid":"0908486E5BFC7DE1","requestTo":"www.example.com"}],"resultInfo":{"totalCount":1,"offset":0,"returnedCount":1}}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	service, err := webforward.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = service.Read("example.com.", "09084E85FE7773E2")

	expected := "Resource not found: { name: 'WebForward', type: 'webForward', key:'09084E85FE7773E2'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("read not found error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestUpdateWithServerError(t *testing.T) {
	service := newServerErrorTestService(t)

	_, err := service.Update("example.com.", "09084E85FE7773E2", &webforward.WebForward{})

	expected := "Error while updating WebForward: " + serverErrorString + ": {key: '09084E85FE7773E2'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("update error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestUpdateWithNilPayload(t *testing.T) {
	service := newServerErrorTestService(t)

	_, err := service.Update("example.com.", "09084E85FE7773E2", nil)

	expected := "Error while updating WebForward: " + serverErrorString + ": {key: '09084E85FE7773E2'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("nil update payload error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestDeleteWithServerError(t *testing.T) {
	service := newServerErrorTestService(t)

	_, err := service.Delete("example.com.", "09084E85FE7773E2")

	expected := "Error while deleting WebForward: " + serverErrorString + ": {key: '09084E85FE7773E2'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("delete error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestListWithServerError(t *testing.T) {
	service := newServerErrorTestService(t)

	_, _, err := service.List("example.com.", nil)

	expected := "Error while listing WebForward: " + serverErrorString + ": {key: 'zones/example.com./webforwards'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("list error mismatched expected - %v : found - %v", expected, err)
	}
}

func TestListWithQueryInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/authorization/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/zones/example.com./webforwards"):
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"webForwards":[{"guid":"0908486E5BFC7DE1","requestTo":"www.example.com"}],"resultInfo":{"totalCount":1,"offset":0,"returnedCount":1}}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	service, err := webforward.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, list, err := service.List("example.com.", &helper.QueryInfo{Sort: "REQUEST_TO"})

	if err != nil {
		t.Fatal(err)
	}

	if len(list.WebForwards) != 1 {
		t.Fatalf("web forward count mismatched expected - 1 : found - %d", len(list.WebForwards))
	}
}
