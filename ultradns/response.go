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

type QueryInfo struct {
	Sort    string `json:"sort,omitempty"`
	Reverse bool   `json:"reverse,omitempty"`
	Limit   int    `json:"limit,omitempty"`
}

type ResultInfo struct {
	TotalCount    int `json:"totalCount,omitempty"`
	Offset        int `json:"offset,omitempty"`
	ReturnedCount int `json:"returnedCount,omitempty"`
}

func Target(i interface{}) *Response {
	return &Response{
		Data:  i,
		Error: &[]ErrorResponse{},
	}
}
