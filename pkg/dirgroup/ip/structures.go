package ip

import (
	"fmt"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	DirGroupType = "ip"
)

type DirGroupIP struct {
	AccountName string       `json:"account_name,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	IPs         []*IPAddress `json:"ips,omitempty"`
}

type Response struct {
	AccountName string       `json:"account_name,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	IPs         []*IPAddress `json:"ips,omitempty"`
}

type ResponseList struct {
	QueryInfo   *helper.QueryInfo  `json:"queryInfo,omitempty"`
	CursorInfo  *helper.CursorInfo `json:"cursorInfo,omitempty"`
	DirGroupIPs []*Response        `json:"dirgroupips,omitempty"`
}

type IPAddress struct {
	Start   string `json:"start,omitempty"`
	End     string `json:"end,omitempty"`
	Cidr    string `json:"cidr,omitempty"`
	Address string `json:"address,omitempty"`
}

func (d DirGroupIP) DirGroupIPID() string {
	return fmt.Sprintf("%s:%s", d.Name, d.AccountName)
}
