package dirpool

const Schema = "http://schemas.ultradns.com/DirPool.jsonschema"

type Profile struct {
	Context         string       `json:"@context,omitempty"`
	Description     string       `json:"description"`
	ConflictResolve string       `json:"conflictResolve,omitempty"`
	IgnoreECS       bool         `json:"ignoreECS"`
	NoResponse      *RDataInfo   `json:"noResponse,omitempty"`
	RDataInfo       []*RDataInfo `json:"rdataInfo,omitempty"`
}

type RDataInfo struct {
	Type             string   `json:"type,omitempty"`
	Ttl              int      `json:"ttl,omitempty"`
	AllNonConfigured bool     `json:"allNonConfigured"`
	GeoInfo          *GEOInfo `json:"geoInfo,omitempty"`
	IPInfo           *IPInfo  `json:"ipInfo,omitempty"`
}

type GEOInfo struct {
	Name                    string   `json:"name,omitempty"`
	Codes                   []string `json:"codes,omitempty"`
	IsExistingGroupFromPool bool     `json:"isExistingGroupFromPool"`
	ForceOverlap            bool     `json:"forceOverlap"`
	IsAccountLevel          bool     `json:"isAccountLevel"`
}

type IPInfo struct {
	Name                    string       `json:"name,omitempty"`
	IsExistingGroupFromPool bool         `json:"isExistingGroupFromPool"`
	IsAccountLevel          bool         `json:"isAccountLevel"`
	IPs                     []*IPAddress `json:"ips,omitempty"`
}

type IPAddress struct {
	Start  string `json:"start,omitempty"`
	End    string `json:"end,omitempty"`
	Cidr   string `json:"cidr,omitempty"`
	Single string `json:"address,omitempty"`
}

func (profile *Profile) SetContext() {
	profile.Context = Schema
}

func (profile *Profile) GetContext() string {
	return profile.Context
}
