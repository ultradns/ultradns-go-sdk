package sfpool

const schema = "http://schemas.ultradns.com/SFPool.jsonschema"

type Profile struct {
	Context                  string        `json:"@context,omitempty"`
	PoolDescription          string        `json:"poolDescription,omitempty"`
	LiveRecordDescription    string        `json:"liveRecordDescription,omitempty"`
	LiveRecordState          string        `json:"liveRecordState,omitempty"`
	RegionFailureSensitivity string        `json:"regionFailureSensitivity,omitempty"`
	Status                   string        `json:"status,omitempty"`
	BackupRecord             *BackupRecord `json:"backupRecord,omitempty"`
	Monitor                  *Monitor      `json:"monitor,omitempty"`
}

type BackupRecord struct {
	RData       string `json:"rdata,omitempty"`
	Description string `json:"description,omitempty"`
}

type Monitor struct {
	Method          string `json:"method,omitempty"`
	URL             string `json:"url,omitempty"`
	TransmittedData string `json:"transmittedData,omitempty"`
	SearchString    string `json:"searchString,omitempty"`
}

func (profile *Profile) SetContext() {
	profile.Context = schema
}
