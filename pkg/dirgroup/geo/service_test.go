package geo_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const serviceErrorString = "DirGroupGeo service configuration failed"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := geo.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := geo.New(conf); err.Error() != "DirGroupGeo service configuration failed: Missing required parameters: [ password ]" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := geo.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateDirGeoGroupWithConfigError(t *testing.T) {
	geoService := geo.Service{}
	if _, err := geoService.Create(&geo.DirGroupGeo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadDirGeoGroupWithConfigError(t *testing.T) {
	geoService := geo.Service{}
	if _, _, _, err := geoService.Read(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateDirGeoGroupWithConfigError(t *testing.T) {
	geoService := geo.Service{}
	if _, err := geoService.Update(&geo.DirGroupGeo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateDirGeoGroupWithConfigError(t *testing.T) {
	geoService := geo.Service{}
	if _, err := geoService.PartialUpdate(&geo.DirGroupGeo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteDirGeoGroupWithConfigError(t *testing.T) {
	geoService := geo.Service{}
	if _, err := geoService.Delete(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListDirGeoGroupWithConfigError(t *testing.T) {
	geoService := geo.Service{}
	if _, _, err := geoService.List(&helper.QueryInfo{}, &geo.DirGroupGeo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateDirGeoGroupFailure(t *testing.T) {
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := geoService.Create(&geo.DirGroupGeo{}); err.Error() != "Error while creating DirGroupGeo: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestReadDirGeoGroupFailure(t *testing.T) {
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, _, err := geoService.Read(""); err.Error() != "Error while reading DirGroupGeo: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestUpdateDirGeoGroupFailure(t *testing.T) {
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := geoService.Update(&geo.DirGroupGeo{}); err.Error() != "Error while updating DirGroupGeo: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestPartialUpdateDirGeoGroupFailure(t *testing.T) {
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := geoService.PartialUpdate(&geo.DirGroupGeo{}); err.Error() != "Error while partial updating DirGroupGeo: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestDeleteDirGeoGroupFailure(t *testing.T) {
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := geoService.Delete(""); err.Error() != "Error while deleting DirGroupGeo: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestListDirGeoGroupFailure(t *testing.T) {
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, err := geoService.List(&helper.QueryInfo{}, &geo.DirGroupGeo{}); err.Error() != "Error while listing DirGroupGeo: Server error Response - { code: '404', message: 'Status Code 404' }: {key: 'accounts//dirgroups/geo'}" {
		t.Fatal(err)
	}
}
