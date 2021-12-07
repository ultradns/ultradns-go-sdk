package record

import (
	"fmt"
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

type RecordService struct {
	c *client.Client
}

func New(config client.Config) (*RecordService, error) {
	client, err := client.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &RecordService{c: client}, nil
}

func Get(client *client.Client) (*RecordService, error) {
	if client == nil {
		return nil, fmt.Errorf("record service is not properly configured")
	}

	return &RecordService{c: client}, nil
}

// Create record with provided rrset.
func (rs *RecordService) CreateRecord(rrSetKey rrset.RRSetKey, rrSet rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, fmt.Errorf("record service is not properly configured")
	}

	res, err := rs.c.Do(http.MethodPost, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rs *RecordService) UpdateRecord(rrSetKey rrset.RRSetKey, rrSet rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, fmt.Errorf("record service is not properly configured")
	}

	res, err := rs.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rs *RecordService) PartialUpdateRecord(rrSetKey rrset.RRSetKey, rrSet rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, fmt.Errorf("record service is not properly configured")
	}

	res, err := rs.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rs *RecordService) ReadRecord(rrSetKey rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, fmt.Errorf("record service is not properly configured")
	}

	res, err := rs.c.Do(http.MethodGet, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rs *RecordService) DeleteRecord(rrSetKey rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, fmt.Errorf("record service is not properly configured")
	}

	res, err := rs.c.Do(http.MethodDelete, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, err
	}

	return res, nil
}
