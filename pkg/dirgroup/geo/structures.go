package geo

import (
	"fmt"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	DirGroupType = "geo"
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

func (d DirGroupGeo) DirGroupGeoID() string {
	return fmt.Sprintf("%s:%s", d.Name, d.AccountName)
}
