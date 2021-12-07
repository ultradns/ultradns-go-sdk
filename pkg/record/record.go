package record

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceName = "Record"

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, client.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, client.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

// Create record with provided rrset.
func (rs *Service) CreateRecord(rrSetKey rrset.RRSetKey, rrSet rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, client.ServiceError(serviceName)
	}

	res, err := rs.c.Do(http.MethodPost, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, CreateRecordError(rrSetKey, err)
	}

	return res, nil
}

func (rs *Service) UpdateRecord(rrSetKey rrset.RRSetKey, rrSet rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, client.ServiceError(serviceName)
	}

	res, err := rs.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, UpdateRecordError(rrSetKey, err)
	}

	return res, nil
}

func (rs *Service) PartialUpdateRecord(rrSetKey rrset.RRSetKey, rrSet rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, client.ServiceError(serviceName)
	}

	res, err := rs.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, PartialUpdateRecordError(rrSetKey, err)
	}

	return res, nil
}

func (rs *Service) ReadRecord(rrSetKey rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, client.ServiceError(serviceName)
	}

	res, err := rs.c.Do(http.MethodGet, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, ReadRecordError(rrSetKey, err)
	}

	return res, nil
}

func (rs *Service) DeleteRecord(rrSetKey rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if rs.c == nil {
		return nil, client.ServiceError(serviceName)
	}

	res, err := rs.c.Do(http.MethodDelete, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, DeleteRecordError(rrSetKey, err)
	}

	return res, nil
}
