package helper

import (
	"fmt"
	"net/url"
)

// QueryInfo
type QueryInfo struct {
	Query   string `json:"q,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	Reverse bool   `json:"reverse,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Offset  int    `json:"offset,omitempty"`
}

// ResultInfo
type ResultInfo struct {
	TotalCount    int `json:"totalCount,omitempty"`
	Offset        int `json:"offset,omitempty"`
	ReturnedCount int `json:"returnedCount,omitempty"`
}

// CursorInfo
type CursorInfo struct {
	Limit    int    `json:"limit,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
}

func (q *QueryInfo) String() string {
	if q.Limit == 0 {
		q.Limit = 100
	}
	queryInfo := fmt.Sprintf("&q=%v&offset=%v&cursor=%v&limit=%v&sort=%v&reverse=%v", q.Query, q.Offset, q.Cursor, q.Limit, q.Sort, q.Reverse)
	return url.PathEscape(queryInfo)
}
