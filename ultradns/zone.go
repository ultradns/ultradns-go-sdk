/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns

import (
	"fmt"
	"net/http"
)

type Zone struct {
	Properties          *ZoneProperties `json:"properties"`
	PrimaryCreateInfo   *PrimaryZone    `json:"primaryCreateInfo,omitempty"`
	SecondaryCreateInfo *SecondaryZone  `json:"secondaryCreateInfo,omitempty"`
	AliasCreateInfo     *AliasZone      `json:"aliasCreateInfo,omitempty"`
	ChangeComment       string          `json:"changeComment,omitempty"`
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
	ChangeComment        string `json:"changeComment,omitempty"`
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
	TsigKey       string `json:"tsigKey,omitempty"`
	TsigKeyValue  string `json:"tsigKeyValue,omitempty"`
	TsigAlgorithm string `json:"tsigAlgorithm,omitempty"`
}

type NameServerIpList struct {
	NameServerIp1 *NameServerIp `json:"nameServerIp1"`
	NameServerIp2 *NameServerIp `json:"nameServerIp2"`
	NameServerIp3 *NameServerIp `json:"nameServerIp3"`
}

type PrimaryNameServers struct {
	NameServerIpList *NameServerIpList `json:"nameServerIpList"`
}

type PrimaryZone struct {
	ForceImport      bool             `json:"forceImport" default:"true"`
	CreateType       string           `json:"createType"`
	NameServer       *NameServerIp    `json:"nameServer,omitempty"`
	Tsig             *Tsig            `json:"tsig,omitempty"`
	OriginalZoneName string           `json:"originalZoneName,omitempty"`
	RestrictIPList   *[]RestrictIp    `json:"restrictIPList,omitempty"`
	NotifyAddresses  *[]NotifyAddress `json:"notifyAddresses,omitempty"`
	Inherit          string           `json:"inherit,omitempty"`
}

type SecondaryZone struct {
	PrimaryNameServers       *PrimaryNameServers `json:"primaryNameServers"`
	NotificationEmailAddress string              `json:"notificationEmailAddress,omitempty"`
	AllowUnResponsiveNs      bool                `json:"allowUnresponsiveNS,omitempty"`
}

type AliasZone struct {
	OriginalZoneName string `json:"originalZoneName"`
}

type NameServersList struct {
	Ok        []string `json:"ok,omitempty"`
	Unknown   []string `json:"unknown,omitempty"`
	Missing   []string `json:"missing,omitempty"`
	Incorrect []string `json:"incorrect,omitempty"`
}

type RegistrarInfo struct {
	Registrar       string           `json:"registrar,omitempty"`
	WhoIsExpiration string           `json:"whoisExpiration,omitempty"`
	NameServers     *NameServersList `json:"nameServers,omitempty"`
}

type TransferStatusDetails struct {
	LastRefresh              string `json:"lastRefresh"`
	NextRefresh              string `json:"nextRefresh"`
	LastRefreshStatus        string `json:"lastRefreshStatus"`
	LastRefreshStatusMessage string `json:"lastRefreshStatusMessage"`
}

type ZoneResponse struct {
	Properties *ZoneProperties `json:"properties"`

	//Primary Zone Response
	RegistrarInfo   *RegistrarInfo   `json:"registrarInfo,omitempty"`
	Tsig            *Tsig            `json:"tsig,omitempty"`
	RestrictIPList  *[]RestrictIp    `json:"restrictIpList,omitempty"`
	NotifyAddresses *[]NotifyAddress `json:"notifyAddresses,omitempty"`

	//Secondary Zone Response
	PrimaryNameServers    *PrimaryNameServers    `json:"primaryNameServers,omitempty"`
	TransferStatusDetails *TransferStatusDetails `json:"transferStatusDetails,omitempty"`

	//Alias Zone Response
	OriginalZoneName string `json:"originalZoneName,omitempty"`
}

type ZoneListResponse struct {
	QueryInfo  *QueryInfo      `json:"queryInfo"`
	ResultInfo *ResultInfo     `json:"resultInfo"`
	Zones      *[]ZoneResponse `json:"zones"`
}

//create zone
func (c *Client) CreateZone(zone Zone) (*http.Response, error) {
	target := Target(&SuccessResponse{})
	res, err := c.Do("POST", "zones", zone, target)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errDataListPtr := target.Error.(*[]ErrorResponse)
		errDataList := *errDataListPtr
		return res, fmt.Errorf("error while creating a zone (%v) - %s", zone.Properties.Name, errDataList[0])
	}

	if res.StatusCode == 202 {
		taskId := res.Header.Get("X-Task-Id")
		er := c.ZoneTaskWait(taskId)

		if er != nil {
			return res, er
		}
	}

	return res, nil
}

//read zone
func (c *Client) ReadZone(zoneName string) (*http.Response, string, *ZoneResponse, error) {
	target := Target(&ZoneResponse{})
	res, err := c.Do("GET", "zones/"+zoneName, nil, target)

	if err != nil {
		return nil, "", nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errDataListPtr := target.Error.(*[]ErrorResponse)
		errDataList := *errDataListPtr
		return res, "", nil, fmt.Errorf("error while reading a zone (%v) - %s", zoneName, errDataList[0])
	}
	zoneResponse := target.Data.(*ZoneResponse)
	zoneType := zoneResponse.Properties.Type

	return res, zoneType, zoneResponse, nil
}

//update zone
func (c *Client) UpdateZone(zoneName string, zone Zone) (*http.Response, error) {
	target := Target(&SuccessResponse{})
	res, err := c.Do("PUT", "zones/"+zoneName, zone, target)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errDataListPtr := target.Error.(*[]ErrorResponse)
		errDataList := *errDataListPtr
		return res, fmt.Errorf("error while updating a zone (%v) - %s", zoneName, errDataList[0])
	}

	return res, nil
}

//delete zone
func (c *Client) DeleteZone(zoneName string) (*http.Response, error) {
	target := Target(&SuccessResponse{})
	res, err := c.Do("DELETE", "zones/"+zoneName, nil, target)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errDataListPtr := target.Error.(*[]ErrorResponse)
		errDataList := *errDataListPtr
		return res, fmt.Errorf("error while Deleting a zone (%v) - %s", zoneName, errDataList[0])
	}

	return res, nil
}

//list Zones
func (c *Client) ListZone(query string) (*http.Response, *ZoneListResponse, error) {
	target := Target(&ZoneListResponse{})
	res, err := c.Do("GET", "zones/"+query, nil, target)

	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errDataListPtr := target.Error.(*[]ErrorResponse)
		errDataList := *errDataListPtr
		return res, nil, fmt.Errorf("error while listing zones - %s", errDataList[0])
	}
	zoneListResponse := target.Data.(*ZoneListResponse)

	return res, zoneListResponse, nil
}
