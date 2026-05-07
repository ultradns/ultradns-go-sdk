package resource_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

const serviceErrorString = "CDNResource service configuration failed"

func TestNewError(t *testing.T) {
	if _, err := resource.New(client.Config{}); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetError(t *testing.T) {
	if _, err := resource.Get(nil); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateWithConfigError(t *testing.T) {
	svc := resource.Service{}
	if _, err := svc.Create("acc1", "example.com.", &resource.Resource{}); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadWithConfigError(t *testing.T) {
	svc := resource.Service{}
	if _, _, err := svc.Read("acc1", "example.com."); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateWithConfigError(t *testing.T) {
	svc := resource.Service{}
	if _, err := svc.Update("acc1", "example.com.", &resource.Resource{}); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteWithConfigError(t *testing.T) {
	svc := resource.Service{}
	if _, err := svc.Delete("acc1", "example.com."); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListWithConfigError(t *testing.T) {
	svc := resource.Service{}
	if _, _, err := svc.List("acc1", nil); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}