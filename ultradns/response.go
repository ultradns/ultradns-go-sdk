/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
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
