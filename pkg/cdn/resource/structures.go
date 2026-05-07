package resource

import "encoding/json"

const (
	TypeBYOD      = "BYOD"
	TypeSynthetic = "SYNTHETIC"
)

// CdnConfig describes a single CDN provider entry within a CDN resource.
type CdnConfig struct {
	ClientCdnID string `json:"clientCdnId,omitempty"`
	CdnName     string `json:"cdnName,omitempty"`
	Description string `json:"description,omitempty"`
	FQDN        string `json:"fqdn,omitempty"`
}

// Configs wraps the configs block of a CDN resource.
// The API uses @JsonAnyGetter/@JsonAnySetter on the Java side, meaning
// AdditionalProperties are serialized as inline fields alongside "cdns".
type Configs struct {
	CDNs                 []*CdnConfig           `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MarshalJSON serializes Configs so that CDNs appears as "cdns" and all
// additional properties are inlined at the same level.
func (c Configs) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range c.AdditionalProperties {
		m[k] = v
	}
	if len(c.CDNs) > 0 {
		m["cdns"] = c.CDNs
	}
	return json.Marshal(m)
}

// UnmarshalJSON deserializes Configs, pulling out "cdns" and leaving
// everything else in AdditionalProperties.
func (c *Configs) UnmarshalJSON(data []byte) error {
	c.CDNs = nil
	c.AdditionalProperties = nil

	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if cdns, ok := m["cdns"]; ok {
		if err := json.Unmarshal(cdns, &c.CDNs); err != nil {
			return err
		}
		delete(m, "cdns")
	}
	c.AdditionalProperties = make(map[string]interface{})
	for k, v := range m {
		var val interface{}
		if err := json.Unmarshal(v, &val); err != nil {
			return err
		}
		c.AdditionalProperties[k] = val
	}
	return nil
}

// Preferences wraps the preferences block of a CDN resource.
// All fields are free-form and serialized inline (same @JsonAnyGetter pattern).
type Preferences struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MarshalJSON serializes Preferences by inlining all additional properties.
func (p Preferences) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range p.AdditionalProperties {
		m[k] = v
	}
	return json.Marshal(m)
}

// UnmarshalJSON deserializes Preferences into AdditionalProperties.
func (p *Preferences) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	p.AdditionalProperties = make(map[string]interface{})
	for k, v := range m {
		var val interface{}
		if err := json.Unmarshal(v, &val); err != nil {
			return err
		}
		p.AdditionalProperties[k] = val
	}
	return nil
}

// Resource wraps the CDN resource payload handled by zone-level APIs.
// Maps to Java MultiCdnResource DTO.
type Resource struct {
	// Writeable fields
	FQDN        string       `json:"fqdn,omitempty"`
	Type        string       `json:"type,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	TTL         int          `json:"ttl,omitempty"`
	ContentType string       `json:"contentType,omitempty"`
	Configs     *Configs     `json:"configs,omitempty"`
	Preferences *Preferences `json:"preferences,omitempty"`
	// Read-only / computed fields
	ResourceID  int    `json:"resourceId,omitempty"`
	Version     string `json:"version,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	OwnerName   string `json:"ownerName,omitempty"`
}

// ListOptions wraps list query options for CDN listing APIs.
type ListOptions struct {
	Page int
	Size int
}

// Item wraps a list item returned in list response.
type Item struct {
	FQDN       string `json:"fqdn,omitempty"`
	Type       string `json:"type,omitempty"`
	ResourceID int    `json:"resourceId,omitempty"`
	Name       string `json:"name,omitempty"`
}

// ResponseList wraps list response payload returned by CDN list API.
type ResponseList struct {
	Content       []*Item `json:"content,omitempty"`
	Page          int     `json:"page,omitempty"`
	Size          int     `json:"size,omitempty"`
	TotalPages    int     `json:"totalPages,omitempty"`
	TotalElements int     `json:"totalElements,omitempty"`
}