package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

// TestCDNResources orchestrates CRUD subtests for the CDN resource service.
// It must be called after the zone for zoneName has been created, since the
// REST API validates that the zone exists. MultiCDN does not require the zone
// to be GTD-enabled; GTD-related handling is managed separately by the
// MultiCDN feature.
func (t *IntegrationTest) TestCDNResources(zoneName string) {
	fqdn := "www." + zoneName
	name := "cdn-integ-" + integration.GetRandomString()
	updatedName := "cdn-integ-" + integration.GetRandomString()

	t.Test.Run("TestCreateCDNResource", func(st *testing.T) {
		it := IntegrationTest{Test: st}
		it.CreateCDNResource(fqdn, name)
	})
	t.Test.Run("TestReadCDNResource", func(st *testing.T) {
		it := IntegrationTest{Test: st}
		it.ReadCDNResource(fqdn)
	})
	t.Test.Run("TestUpdateCDNResource", func(st *testing.T) {
		it := IntegrationTest{Test: st}
		it.UpdateCDNResource(fqdn, updatedName)
	})
	t.Test.Run("TestListCDNResources", func(st *testing.T) {
		it := IntegrationTest{Test: st}
		it.ListCDNResources()
	})
	t.Test.Run("TestDeleteCDNResource", func(st *testing.T) {
		it := IntegrationTest{Test: st}
		it.DeleteCDNResource(fqdn)
	})
}

func (t *IntegrationTest) CreateCDNResource(fqdn, name string) {
	svc, err := cdnresource.Get(integration.TestClientCDN)
	if err != nil {
		t.Test.Fatal(err)
	}

	payload := getBYODCDNResource(name)
	if _, er := svc.Create(integration.TestAccountCDN, fqdn, payload); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ReadCDNResource(fqdn string) {
	svc, err := cdnresource.Get(integration.TestClientCDN)
	if err != nil {
		t.Test.Fatal(err)
	}

	_, res, er := svc.Read(integration.TestAccountCDN, fqdn)
	if er != nil {
		t.Test.Fatal(er)
	}

	if res == nil {
		t.Test.Fatal("expected non-nil CDN resource response")
	}
}

func (t *IntegrationTest) UpdateCDNResource(fqdn, updatedName string) {
	svc, err := cdnresource.Get(integration.TestClientCDN)
	if err != nil {
		t.Test.Fatal(err)
	}

	payload := getBYODCDNResource(updatedName)
	if _, er := svc.Update(integration.TestAccountCDN, fqdn, payload); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ListCDNResources() {
	svc, err := cdnresource.Get(integration.TestClientCDN)
	if err != nil {
		t.Test.Fatal(err)
	}

	_, res, er := svc.List(integration.TestAccountCDN, &cdnresource.ListOptions{Page: 1, Size: 100})
	if er != nil {
		t.Test.Fatal(er)
	}

	if res == nil {
		t.Test.Fatal("expected non-nil CDN resource list response")
	}
}

func (t *IntegrationTest) DeleteCDNResource(fqdn string) {
	svc, err := cdnresource.Get(integration.TestClientCDN)
	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := svc.Delete(integration.TestAccountCDN, fqdn); er != nil {
		t.Test.Fatal(er)
	}
}

// getBYODCDNResource builds a minimal but complete BYOD CDN resource payload.
func getBYODCDNResource(name string) *cdnresource.Resource {
	return &cdnresource.Resource{
		Name: name,
		Type: cdnresource.TypeBYOD,
		TTL:  300,
		Configs: &cdnresource.Configs{
			CDNs: []*cdnresource.CdnConfig{
				{
					ClientCdnID: "INTEG-PROVIDER-1",
					CdnName:     "Integration Test CDN",
					FQDN:        "cdn-integ-provider.example.com.",
				},
			},
			AdditionalProperties: map[string]interface{}{
				"cdnEnablementMap": map[string]interface{}{
					"worldDefault": []interface{}{"INTEG-PROVIDER-1"},
					"asnOverrides": map[string]interface{}{},
					"continents":   map[string]interface{}{},
				},
				"trafficDistribution": map[string]interface{}{
					"worldDefault": map[string]interface{}{
						"options": []interface{}{
							map[string]interface{}{
								"name":        "equal_distribution",
								"equalWeight": true,
								"distribution": []interface{}{
									map[string]interface{}{
										"id": "INTEG-PROVIDER-1",
									},
								},
							},
						},
					},
				},
			},
		},
		Preferences: &cdnresource.Preferences{
			AdditionalProperties: map[string]interface{}{
				"availabilityThresholds": map[string]interface{}{
					"world": 92,
					"continents": map[string]interface{}{
						"NA": map[string]interface{}{
							"default": 79,
							"countries": map[string]interface{}{
								"US": 70,
							},
						},
					},
				},
				"performanceFiltering": map[string]interface{}{
					"world": map[string]interface{}{
						"mode":              "relative",
						"relativeThreshold": 0.2,
					},
					"continents": map[string]interface{}{},
				},
				"enabledSubdivisionCountries": map[string]interface{}{
					"continents": map[string]interface{}{
						"NA": map[string]interface{}{
							"countries": []interface{}{"US"},
						},
					},
				},
			},
		},
	}
}
