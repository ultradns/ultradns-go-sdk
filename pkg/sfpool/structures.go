package sfpool

import "github.com/ultradns/ultradns-go-sdk/pkg/pool"

const schema = "http://schemas.ultradns.com/SFPool.jsonschema"

type Profile struct {
	Context                  string             `json:"@context,omitempty"`
	PoolDescription          string             `json:"poolDescription,omitempty"`
	LiveRecordDescription    string             `json:"liveRecordDescription,omitempty"`
	LiveRecordState          string             `json:"liveRecordState,omitempty"`
	RegionFailureSensitivity string             `json:"regionFailureSensitivity,omitempty"`
	Status                   string             `json:"status,omitempty"`
	BackupRecord             *pool.BackupRecord `json:"backupRecord,omitempty"`
	Monitor                  *pool.Monitor      `json:"monitor,omitempty"`
}

func (profile *Profile) SetContext() {
	profile.Context = schema
}
