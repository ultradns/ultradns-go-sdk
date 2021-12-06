package client

import (
	"fmt"
	"net/http"
)

//Config struct wraps the credential info for the Client
type Config struct {
	Username   string
	Password   string
	HostURL    string
	APIVersion string
	UserAgent  string
}

//client struct wraps the http client, config and ultradns api base url
type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
}

// Response
type Response struct {
	Data  interface{}
	Error []*ErrorResponse
}

// ErrorResponse
type ErrorResponse struct {
	ErrorCode        int    `json:"errorCode,omitempty"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
	ErrorString      string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// SuccessResponse
type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

func (e ErrorResponse) String() string {
	return fmt.Sprintf("error code : %v - error message : %v", e.ErrorCode, e.ErrorMessage)
}

func Target(i interface{}) *Response {
	return &Response{Data: i}
}
