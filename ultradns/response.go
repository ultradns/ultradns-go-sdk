package ultradns

type Response struct {
	Data  interface{}
	Error interface{}
}

type ErrorResponse struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorString      string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func Target(i interface{}) *Response {
	return &Response{
		Data:  i,
		Error: &[]ErrorResponse{},
	}
}
