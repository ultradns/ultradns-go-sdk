package zone

import "github.com/ultradns/ultradns-go-sdk/pkg/helper"

// Zone
type Zone struct {
	Properties          *ZoneProperties `json:"properties,omitempty"`
	PrimaryCreateInfo   *PrimaryZone    `json:"primaryCreateInfo,omitempty"`
	SecondaryCreateInfo *SecondaryZone  `json:"secondaryCreateInfo,omitempty"`
	AliasCreateInfo     *AliasZone      `json:"aliasCreateInfo,omitempty"`
	ChangeComment       string          `json:"changeComment,omitempty"`
}

// ZoneProperties
type ZoneProperties struct {
	Name                 string `json:"name,omitempty"`
	AccountName          string `json:"accountName,omitempty"`
	Type                 string `json:"type,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Status               string `json:"status,omitempty"`
	DNSSecStatus         string `json:"dnssecStatus,omitempty"`
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`
	ResourceRecordCount  int    `json:"resourceRecordCount,omitempty"`
	ChangeComment        string `json:"changeComment,omitempty"`
}

// PrimaryZone
type PrimaryZone struct {
	ForceImport      bool             `json:"forceImport,omitempty"`
	CreateType       string           `json:"createType,omitempty"`
	NameServer       *NameServer      `json:"nameServer,omitempty"`
	Tsig             *Tsig            `json:"tsig,omitempty"`
	OriginalZoneName string           `json:"originalZoneName,omitempty"`
	RestrictIPList   []*RestrictIP    `json:"restrictIPList,omitempty"`
	NotifyAddresses  []*NotifyAddress `json:"notifyAddresses,omitempty"`
	Inherit          string           `json:"inherit,omitempty"`
}

// SecondaryZone
type SecondaryZone struct {
	PrimaryNameServers       *PrimaryNameServers `json:"primaryNameServers,omitempty"`
	NotificationEmailAddress string              `json:"notificationEmailAddress,omitempty"`
	AllowUnResponsiveNs      bool                `json:"allowUnresponsiveNS,omitempty"`
}

// AliasZone
type AliasZone struct {
	OriginalZoneName string `json:"originalZoneName,omitempty"`
}

// NameServer
type NameServer struct {
	IP            string `json:"ip,omitempty"`
	TsigKey       string `json:"tsigKey,omitempty"`
	TsigKeyValue  string `json:"tsigKeyValue,omitempty"`
	TsigAlgorithm string `json:"tsigAlgorithm,omitempty"`
}

// Tsig
type Tsig struct {
	TsigKeyName   string `json:"tsigKeyName,omitempty"`
	TsigKeyValue  string `json:"tsigKeyValue,omitempty"`
	TsigAlgorithm string `json:"tsigAlgorithm,omitempty"`
	Description   string `json:"description,omitempty"`
}

// RestrictIP
type RestrictIP struct {
	StartIP  string `json:"startIP,omitempty"`
	EndIP    string `json:"endIP,omitempty"`
	Cidr     string `json:"cidr,omitempty"`
	SingleIP string `json:"singleIP,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

// NotifyAddress
type NotifyAddress struct {
	NotifyAddress string `json:"notifyAddress,omitempty"`
	Description   string `json:"description,omitempty"`
}

// NameServerIPList
type NameServerIPList struct {
	NameServerIP1 *NameServer `json:"nameServerIp1,omitempty"`
	NameServerIP2 *NameServer `json:"nameServerIp2,omitempty"`
	NameServerIP3 *NameServer `json:"nameServerIp3,omitempty"`
}

// PrimaryNameServers
type PrimaryNameServers struct {
	NameServerIPList *NameServerIPList `json:"nameServerIpList,omitempty"`
}

// NameServersList
type NameServersList struct {
	Ok        []string `json:"ok,omitempty"`
	Unknown   []string `json:"unknown,omitempty"`
	Missing   []string `json:"missing,omitempty"`
	Incorrect []string `json:"incorrect,omitempty"`
}

// RegistrarInfo
type RegistrarInfo struct {
	Registrar       string           `json:"registrar,omitempty"`
	WhoIsExpiration string           `json:"whoisExpiration,omitempty"`
	NameServers     *NameServersList `json:"nameServers,omitempty"`
}

// TransferStatusDetails
type TransferStatusDetails struct {
	LastRefresh              string `json:"lastRefresh,omitempty"`
	NextRefresh              string `json:"nextRefresh,omitempty"`
	LastRefreshStatus        string `json:"lastRefreshStatus,omitempty"`
	LastRefreshStatusMessage string `json:"lastRefreshStatusMessage,omitempty"`
}

// ZoneResponse
type ZoneResponse struct {
	Properties *ZoneProperties `json:"properties,omitempty"`

	//Primary Zone Response
	RegistrarInfo   *RegistrarInfo   `json:"registrarInfo,omitempty"`
	Tsig            *Tsig            `json:"tsig,omitempty"`
	RestrictIPList  []*RestrictIP    `json:"restrictIpList,omitempty"`
	NotifyAddresses []*NotifyAddress `json:"notifyAddresses,omitempty"`

	//Secondary Zone Response
	PrimaryNameServers       *PrimaryNameServers    `json:"primaryNameServers,omitempty"`
	TransferStatusDetails    *TransferStatusDetails `json:"transferStatusDetails,omitempty"`
	NotificationEmailAddress string                 `json:"notificationEmailAddress,omitempty"`

	//Alias Zone Response
	OriginalZoneName string `json:"originalZoneName,omitempty"`
}

// ZoneListResponse
type ZoneListResponse struct {
	QueryInfo  *helper.QueryInfo  `json:"queryInfo,omitempty"`
	ResultInfo *helper.ResultInfo `json:"resultInfo,omitempty"`
	CursorInfo *helper.CursorInfo `json:"cursorInfo,omitempty"`
	Zones      []*ZoneResponse    `json:"zones,omitempty"`
}
