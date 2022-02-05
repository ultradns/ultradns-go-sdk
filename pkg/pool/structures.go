package pool

type Monitor struct {
	Method          string `json:"method,omitempty"`
	URL             string `json:"url,omitempty"`
	TransmittedData string `json:"transmittedData"`
	SearchString    string `json:"searchString"`
}
