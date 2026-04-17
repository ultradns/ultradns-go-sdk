package provider

// Provider wraps the account-level CDN provider payload.
type Provider struct {
	ClientCdnID string `json:"clientCdnId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        int    `json:"type,omitempty"`
}

// Response wraps single CDN provider read payload.
type Response struct {
	Provider *Provider `json:"provider,omitempty"`
}

// ResponseList wraps CDN provider list payload.
type ResponseList struct {
	Providers []*Provider `json:"providers,omitempty"`
}