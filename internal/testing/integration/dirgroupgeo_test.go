package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
)

var geoID = ""

func (t *IntegrationTest) TestDirGroupGeoResources(geoData *geo.DirGroupGeo, geoName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreatetDirGroupGeo",
		func(st *testing.T) {
			it.Test = st
			it.CreatetDirGroupGeo(geoData, geoName)
		})
	t.Test.Run("TestListDirGroupGeoResource",
		func(st *testing.T) {
			it.Test = st
			it.ListDirGroupGeo(geoData, geoName)
		})
	t.Test.Run("TestUpdateDirGroupGeo",
		func(st *testing.T) {
			it.Test = st
			it.UpdateDirGroupGeo(geoData, geoName)
		})
	t.Test.Run("TestPartialUpdateDirGroupGeo",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateDirGroupGeo(geoData, geoName)
		})
	t.Test.Run("TestReadDirGroupGeo",
		func(st *testing.T) {
			it.Test = st
			it.ReadDirGroupGeo(geoData, geoName)
		})
	t.Test.Run("TestDeleteDirGroupGeo",
		func(st *testing.T) {
			it.Test = st
			it.DeleteDirGroupGeo(geoData, geoName)
		})
}

func (t *IntegrationTest) CreatetDirGroupGeo(geoData *geo.DirGroupGeo, geoName string) {
	geodata := getDirGroupGeo(geoName)
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := geoService.Create(geodata); err != nil {
		t.Test.Fatal(err)

	}
}

func (t *IntegrationTest) ListDirGroupGeo(geoData *geo.DirGroupGeo, geoName string) {
	geodata := getDirGroupGeo(geoName)
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	_, _, err = geoService.List(nil, geodata)

	if err != nil {
		t.Test.Fatal(err)
	}
}

func (t *IntegrationTest) UpdateDirGroupGeo(geoData *geo.DirGroupGeo, geoName string) {
	geodata := getDirGroupGeo(geoName)
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := geoService.Update(geodata); err != nil {
		t.Test.Fatal(err)
	}

}

func (t *IntegrationTest) PartialUpdateDirGroupGeo(geoData *geo.DirGroupGeo, geoName string) {
	geodata := getDirGroupGeo(geoName)
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := geoService.PartialUpdate(geodata); err != nil {
		t.Test.Fatal(err)
	}
}

func (t *IntegrationTest) ReadDirGroupGeo(geoData *geo.DirGroupGeo, geoName string) {
	geodata := getDirGroupGeo(geoName)
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, _, _, err := geoService.Read(geodata.Name + ":" + geodata.AccountName); err != nil {
		t.Test.Fatal(err)
	}
}

func (t *IntegrationTest) DeleteDirGroupGeo(geoData *geo.DirGroupGeo, geoName string) {
	geodata := getDirGroupGeo(geoName)
	geoService, err := geo.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := geoService.Delete(geodata.Name + ":" + geodata.AccountName); err != nil {
		t.Test.Fatal(err)
	}
}

func getDirGroupGeo(geoName string) *geo.DirGroupGeo {
	geodata := &geo.DirGroupGeo{
		Name:        geoName,
		AccountName: integration.TestAccount,
		Codes:       []string{"CA", "US", "MX"},
		Description: "Description of GEO",
	}

	return geodata
}
