package resource_test

import (
	"encoding/json"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

func TestConfigsMarshalJSONIncludesCDNsAndInlineProperties(t *testing.T) {
	cfg := resource.Configs{
		CDNs: []*resource.CdnConfig{
			{
				ClientCdnID: "TEST-PROVIDER-1",
				CdnName:     "Provider One",
				Description: "primary",
				FQDN:        "provider.example.com.",
			},
		},
		AdditionalProperties: map[string]interface{}{
			"checksConfig": map[string]interface{}{
				"protocol": "HTTP",
				"path":     "/",
			},
			"trafficDistribution": map[string]interface{}{
				"worldDefault": map[string]interface{}{
					"strategy": "performance",
				},
			},
		},
	}

	payload, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal configs failed: %v", err)
	}

	var got map[string]json.RawMessage
	if err = json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal marshaled payload failed: %v", err)
	}

	if _, ok := got["cdns"]; !ok {
		t.Fatal("expected key 'cdns' in marshaled payload")
	}
	if _, ok := got["checksConfig"]; !ok {
		t.Fatal("expected key 'checksConfig' in marshaled payload")
	}
	if _, ok := got["trafficDistribution"]; !ok {
		t.Fatal("expected key 'trafficDistribution' in marshaled payload")
	}

	var cdns []resource.CdnConfig
	if err = json.Unmarshal(got["cdns"], &cdns); err != nil {
		t.Fatalf("unmarshal cdns failed: %v", err)
	}

	if len(cdns) != 1 {
		t.Fatalf("expected 1 cdn entry, got %d", len(cdns))
	}
	if cdns[0].ClientCdnID != "TEST-PROVIDER-1" {
		t.Fatalf("expected clientCdnId TEST-PROVIDER-1, got %s", cdns[0].ClientCdnID)
	}
}

func TestConfigsUnmarshalJSONExtractsCDNsAndAdditionalProperties(t *testing.T) {
	input := `{
		"cdns": [
			{
				"clientCdnId": "TEST-PROVIDER-1",
				"cdnName": "Provider One",
				"description": "primary",
				"fqdn": "provider.example.com."
			}
		],
		"checksConfig": {"protocol": "HTTP", "path": "/"},
		"trafficDistribution": {"worldDefault": {"strategy": "performance"}}
	}`

	var cfg resource.Configs
	if err := json.Unmarshal([]byte(input), &cfg); err != nil {
		t.Fatalf("unmarshal configs failed: %v", err)
	}

	if len(cfg.CDNs) != 1 {
		t.Fatalf("expected 1 cdn entry, got %d", len(cfg.CDNs))
	}
	if cfg.CDNs[0].ClientCdnID != "TEST-PROVIDER-1" {
		t.Fatalf("expected clientCdnId TEST-PROVIDER-1, got %s", cfg.CDNs[0].ClientCdnID)
	}

	if cfg.AdditionalProperties == nil {
		t.Fatal("expected AdditionalProperties to be initialized")
	}
	if _, ok := cfg.AdditionalProperties["cdns"]; ok {
		t.Fatal("did not expect 'cdns' in AdditionalProperties")
	}
	if _, ok := cfg.AdditionalProperties["checksConfig"]; !ok {
		t.Fatal("expected 'checksConfig' in AdditionalProperties")
	}
	if _, ok := cfg.AdditionalProperties["trafficDistribution"]; !ok {
		t.Fatal("expected 'trafficDistribution' in AdditionalProperties")
	}
}

func TestConfigsRoundTripJSON(t *testing.T) {
	original := resource.Configs{
		CDNs: []*resource.CdnConfig{
			{ClientCdnID: "TEST-PROVIDER-1", CdnName: "Provider One"},
		},
		AdditionalProperties: map[string]interface{}{
			"cdnEnablementMap": map[string]interface{}{
				"worldDefault": []interface{}{"TEST-PROVIDER-1"},
				"asnOverrides": []interface{}{},
				"continents":   map[string]interface{}{},
			},
		},
	}

	payload, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal configs failed: %v", err)
	}

	var decoded resource.Configs
	if err = json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal configs failed: %v", err)
	}

	if len(decoded.CDNs) != 1 {
		t.Fatalf("expected 1 cdn after round-trip, got %d", len(decoded.CDNs))
	}
	if decoded.CDNs[0].ClientCdnID != original.CDNs[0].ClientCdnID {
		t.Fatalf("expected clientCdnId %s, got %s", original.CDNs[0].ClientCdnID, decoded.CDNs[0].ClientCdnID)
	}
	if _, ok := decoded.AdditionalProperties["cdnEnablementMap"]; !ok {
		t.Fatal("expected 'cdnEnablementMap' after round-trip")
	}
}

func TestConfigsMarshalJSONTypedCDNsWinsOverAdditionalPropertiesKey(t *testing.T) {
	cfg := resource.Configs{
		CDNs: []*resource.CdnConfig{
			{ClientCdnID: "TYPED-ID", CdnName: "Typed CDN"},
		},
		AdditionalProperties: map[string]interface{}{
			"cdns": []map[string]interface{}{{
				"clientCdnId": "INLINE-ID",
				"cdnName":     "Inline CDN",
			}},
		},
	}

	payload, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal configs failed: %v", err)
	}

	var got map[string]json.RawMessage
	if err = json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal marshaled payload failed: %v", err)
	}

	var cdns []resource.CdnConfig
	if err = json.Unmarshal(got["cdns"], &cdns); err != nil {
		t.Fatalf("unmarshal cdns failed: %v", err)
	}

	if len(cdns) != 1 {
		t.Fatalf("expected 1 cdn entry, got %d", len(cdns))
	}
	if cdns[0].ClientCdnID != "TYPED-ID" {
		t.Fatalf("expected typed clientCdnId TYPED-ID to win, got %s", cdns[0].ClientCdnID)
	}
}

func TestConfigsUnmarshalJSONClearsStaleCDNsOnReuse(t *testing.T) {
	first := `{
		"cdns": [{"clientCdnId": "FIRST-ID", "cdnName": "First CDN"}],
		"checksConfig": {"protocol": "HTTP"}
	}`
	second := `{
		"trafficDistribution": {"worldDefault": {"strategy": "performance"}}
	}`

	var cfg resource.Configs
	if err := json.Unmarshal([]byte(first), &cfg); err != nil {
		t.Fatalf("first unmarshal failed: %v", err)
	}
	if len(cfg.CDNs) != 1 {
		t.Fatalf("expected 1 cdn after first unmarshal, got %d", len(cfg.CDNs))
	}

	if err := json.Unmarshal([]byte(second), &cfg); err != nil {
		t.Fatalf("second unmarshal failed: %v", err)
	}
	if cfg.CDNs != nil {
		t.Fatalf("expected CDNs to be nil when second payload omits cdns, got %d entries", len(cfg.CDNs))
	}
	if _, ok := cfg.AdditionalProperties["checksConfig"]; ok {
		t.Fatal("expected stale checksConfig to be removed on reuse")
	}
	if _, ok := cfg.AdditionalProperties["trafficDistribution"]; !ok {
		t.Fatal("expected trafficDistribution from second payload")
	}
}

func TestPreferencesRoundTripJSON(t *testing.T) {
	original := resource.Preferences{
		AdditionalProperties: map[string]interface{}{
			"availabilityThresholds": map[string]interface{}{
				"world": 85,
			},
			"performanceFiltering": map[string]interface{}{
				"world": map[string]interface{}{"mode": "absolute"},
			},
		},
	}

	payload, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal preferences failed: %v", err)
	}

	var decoded resource.Preferences
	if err = json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal preferences failed: %v", err)
	}

	if _, ok := decoded.AdditionalProperties["availabilityThresholds"]; !ok {
		t.Fatal("expected 'availabilityThresholds' after round-trip")
	}
	if _, ok := decoded.AdditionalProperties["performanceFiltering"]; !ok {
		t.Fatal("expected 'performanceFiltering' after round-trip")
	}
}
