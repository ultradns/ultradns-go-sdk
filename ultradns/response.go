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
	ErrorString      string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ZoneResponse struct {
	Properties *ZoneProperties `json:"properties"`

	//Primary Zone Response
	RegistrarInfo   *RegistrarInfo   `json:"registrarInfo,omitempty"`
	Tsig            *Tsig            `json:"tsig,omitempty"`
	RestrictIPList  *[]RestrictIp    `json:"restrictIPList,omitempty"`
	NotifyAddresses *[]NotifyAddress `json:"notifyAddresses,omitempty"`

	//Secondary Zone Response
	PrimaryNameServers    *PrimaryNameServers    `json:"primaryNameServers,omitempty"`
	TransferStatusDetails *TransferStatusDetails `json:"transferStatusDetails,omitempty"`

	//Alias Zone Response
	OriginalZoneName string `json:"originalZoneName,omitempty"`
}

type QueryInfo struct {
	Query   string `json:"q,omitempty"`
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
