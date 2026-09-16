package webforward

import "github.com/ultradns/ultradns-go-sdk/pkg/helper"

// WebForward wraps the structure of a web forward resource.
type WebForward struct {
	GUID                   string   `json:"guid,omitempty"`
	RequestTo              string   `json:"requestTo,omitempty"`
	DefaultRedirectTo      string   `json:"defaultRedirectTo,omitempty"`
	DefaultForwardType     string   `json:"defaultForwardType,omitempty"`
	RelativeForwardType    string   `json:"relativeForwardType,omitempty"`
	DefaultRedirectType    string   `json:"defaultRedirectType,omitempty"`
	CertificateID          string   `json:"certificateId,omitempty"`
	CertificateManagedType string   `json:"certificateManagedType,omitempty"`
	CertificateName        string   `json:"certificateName,omitempty"`
	State                  string   `json:"state,omitempty"`
	ErrorDescription       string   `json:"errorDescription,omitempty"`
	ExpirationDays         string   `json:"expirationDays,omitempty"`
	Records                []Record `json:"records,omitempty"`
}

// Record wraps a rule entry of an advanced web forward. Records specify where
// to forward based on custom request headers; when records are present they
// take precedence over the default redirect.
type Record struct {
	RedirectTo  string `json:"redirectTo,omitempty"`
	ForwardType string `json:"forwardType,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Rules       []Rule `json:"rules,omitempty"`
}

// Rule wraps a single header match of an advanced web forward record.
type Rule struct {
	Header          string `json:"header,omitempty"`
	MatchCriteria   string `json:"matchCriteria,omitempty"`
	Value           string `json:"value,omitempty"`
	CaseInsensitive bool   `json:"caseInsensitive,omitempty"`
}

// ResponseList wraps the result of a web forward list operation
// (GET /zones/{zoneName}/webforwards).
type ResponseList struct {
	WebForwards []*WebForward      `json:"webForwards,omitempty"`
	QueryInfo   *helper.QueryInfo  `json:"queryInfo,omitempty"`
	ResultInfo  *helper.ResultInfo `json:"resultInfo,omitempty"`
}

const (
	// HTTP301Redirect keeps the redirect permanent.
	HTTP301Redirect = "HTTP_301_REDIRECT"
	// HTTP302Redirect keeps the redirect temporary.
	HTTP302Redirect = "HTTP_302_REDIRECT"
	// HTTP303Redirect is the See Other redirect type.
	HTTP303Redirect = "HTTP_303_REDIRECT"
	// HTTP307Redirect is the Temporary Redirect type.
	HTTP307Redirect = "HTTP_307_REDIRECT"
	// Framed renders the target inside an invisible frame.
	Framed = "Framed"

	// Path preserves the request path on the redirect target.
	Path = "PATH"
	// Parameter preserves the request query string on the redirect target.
	Parameter = "PARAMETER"
	// ParameterAndPath preserves both the request path and query string.
	ParameterAndPath = "PARAMETER_AND_PATH"
)
