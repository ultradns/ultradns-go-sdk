package geo

import (
	"fmt"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

type DirGroupGeo struct {
	AccountName string   `json:"account_name,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Codes       []string `json:"codes,omitempty"`
}

type Response struct {
	AccountName string   `json:"account_name,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Codes       []string `json:"codes,omitempty"`
}

type ResponseList struct {
	QueryInfo    *helper.QueryInfo  `json:"queryInfo,omitempty"`
	CursorInfo   *helper.CursorInfo `json:"cursorInfo,omitempty"`
	DirGroupGeos []*Response        `json:"dirgroupgeos,omitempty"`
}

func (d *DirGroupGeo) DirGroupGeoURI() string {
	d.AccountName = url.PathEscape(d.AccountName)
	d.Name = url.PathEscape(d.Name)

	return fmt.Sprintf("accounts/%s/dirgroups/geo/%s", d.AccountName, d.Name)
}

func (d *DirGroupGeo) DirGroupGeoListURI() string {
	d.AccountName = url.PathEscape(d.AccountName)

	return fmt.Sprintf("accounts/%s/dirgroups/geo", d.AccountName)
}
