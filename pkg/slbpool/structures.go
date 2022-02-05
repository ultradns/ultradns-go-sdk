package slbpool

import "github.com/ultradns/ultradns-go-sdk/pkg/pool"

const Schema = "http://schemas.ultradns.com/SLBPool.jsonschema"

type Profile struct {
	Context                  string         `json:"@context,omitempty"`
	ResponseMethod           string         `json:"responseMethod,omitempty"`
	RegionFailureSensitivity string         `json:"regionFailureSensitivity,omitempty"`
	ServingPreference        string         `json:"servingPreference,omitempty"`
	Description              string         `json:"description,omitempty"`
	Status                   string         `json:"status,omitempty"`
	RDataInfo                []*RDataInfo   `json:"rdataInfo,omitempty"`
	AllFailRecord            *AllFailRecord `json:"allFailRecord,omitempty"`
	Monitor                  *pool.Monitor  `json:"monitor,omitempty"`
}

type RDataInfo struct {
	Description      string `json:"description,omitempty"`
	ForcedState      string `json:"forcedState,omitempty"`
	ProbingEnabled   bool   `json:"probingEnabled"`
	AvailableToServe bool   `json:"availableToServe"`
}

type AllFailRecord struct {
	Description string `json:"description,omitempty"`
	RData       string `json:"rdata,omitempty"`
	Serving     bool   `json:"serving"`
}

func (profile *Profile) SetContext() {
	profile.Context = Schema
}

func (profile *Profile) GetContext() string {
	return profile.Context
}
