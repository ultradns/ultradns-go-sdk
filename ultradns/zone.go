package ultradns

type Zone struct {
	Properties          ZoneProperties `json:"properties"`
	PrimaryCreateInfo   PrimaryZone    `json:"primaryCreateInfo,omitempty"`
	SecondaryCreateInfo SecondaryZone  `json:"secondaryCreateInfo,omitempty"`
	AliasCreateInfo     AliasZone      `json:"aliasCreateInfo,omitempty"`
	ChangeComment       string         `json:"changeComment,omitempty"`
}

type ZoneProperties struct {
	Name                 string `json:"name"`
	AccountName          string `json:"accountName"`
	Type                 string `json:"type"`
	Owner                string `json:"owner,omitempty"`
	Status               string `json:"status,omitempty"`
	DnsSecStatus         string `json:"dnssecStatus,omitempty"`
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`
	ResourceRecordCount  int    `json:"resourceRecordCount,omitempty"`
}

type Tsig struct {
	TsigKeyName   string `json:"tsigKeyName"`
	TsigKeyValue  string `json:"tsigKeyValue"`
	TsigAlgorithm string `json:"tsigAlgorithm"`
	Description   string `json:"description,omitempty"`
}

type RestrictIp struct {
	StartIp  string `json:"startIP,omitempty"`
	EndIp    string `json:"endIP,omitempty"`
	Cidr     string `json:"cidr,omitempty"`
	SingleIp string `json:"singleIP,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

type NotifyAddress struct {
	NotifyAddress string `json:"notifyAddress"`
	Description   string `json:"description,omitempty"`
}

type NameServerIp struct {
	Ip            string `json:"ip"`
	TsigKey       string `json:"tsigKey"`
	TsigKeyValue  string `json:"tsigKeyValue"`
	TsigAlgorithm string `json:"tsigAlgorithm"`
}

type NameServerIpList struct {
	NameServerIp1 NameServerIp `json:"nameServerIp1"`
	NameServerIp2 NameServerIp `json:"nameServerIp2"`
	NameServerIp3 NameServerIp `json:"nameServerIp3"`
}

type PrimaryZone struct {
	ForceImport      bool            `json:"forceImport" default:"true"`
	CreateType       string          `json:"createType"`
	NameServer       NameServerIp    `json:"nameServer,omitempty"`
	Tsig             Tsig            `json:"tsig,omitempty"`
	OriginalZoneName string          `json:"originalZoneName,omitempty"`
	RestrictIPList   []RestrictIp    `json:"restrictIPList,omitempty"`
	NotifyAddresses  []NotifyAddress `json:"notifyAddresses,omitempty"`
	Inherit          string          `json:"inherit,omitempty"`
}

type SecondaryZone struct {
	PrimaryNameServers       NameServerIpList `json:"primaryNameServers"`
	NotificationEmailAddress string           `json:"notificationEmailAddress,omitempty"`
}

type AliasZone struct {
	OriginalZoneName string `json:"originalZoneName"`
}
