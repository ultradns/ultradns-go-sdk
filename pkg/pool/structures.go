package pool

type Monitor struct {
	Method          string `json:"method,omitempty"`
	URL             string `json:"url,omitempty"`
	TransmittedData string `json:"transmittedData,omitempty"`
	SearchString    string `json:"searchString,omitempty"`
}
