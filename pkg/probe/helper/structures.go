package helper

type LimitsInfo struct {
	Connect      *Limit        `json:"connect,omitempty"`
	AvgConnect   *Limit        `json:"avgConnect,omitempty"`
	Run          *Limit        `json:"run,omitempty"`
	AvgRun       *Limit        `json:"avgRun,omitempty"`
	Total        *Limit        `json:"total,omitempty"`
	Average      *Limit        `json:"average,omitempty"`
	LossPercent  *Limit        `json:"lossPercent,omitempty"`
	SearchString *SearchString `json:"searchString,omitempty"`
	Response     *SearchString `json:"response,omitempty"`
}

type Limit struct {
	Warning  int `json:"warning,omitempty"`
	Critical int `json:"critical,omitempty"`
	Fail     int `json:"fail,omitempty"`
}

type SearchString struct {
	Warning  string `json:"warning,omitempty"`
	Critical string `json:"critical,omitempty"`
	Fail     string `json:"fail,omitempty"`
}
