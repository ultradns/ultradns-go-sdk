package ip

import (
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	serviceName = "DirGroupIP"
)

type Service struct {
	c *client.Client
}

func New(cfn client.Config) (*Service, error) {
	c, err := client.NewClient(cfn)
	if err != nil {
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) Create(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, helper.GetDirGroupURI(dirGroupIP.DirGroupIPID(), DirGroupType), dirGroupIP, target)

	if err != nil {
		ipGroupName := dirGroupIP.Name

		return res, errors.CreateError(serviceName, ipGroupName, err)
	}

	return res, nil
}

func (s *Service) Read(dirGroupID string) (*http.Response, *Response, string, error) {
	target := client.Target(&Response{})

	dirGroupURI := helper.GetDirGroupURI(dirGroupID, DirGroupType)
	if s.c == nil {
		return nil, nil, dirGroupURI, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, helper.GetDirGroupURI(dirGroupID, DirGroupType), nil, target)
	if err != nil {
		return res, nil, dirGroupURI, errors.ReadError(serviceName, dirGroupID, err)
	}

	dirGroupIPResponse := target.Data.(*Response)

	return res, dirGroupIPResponse, dirGroupURI, nil
}

func (s *Service) Update(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupIPName := url.PathEscape(dirGroupIP.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, helper.GetDirGroupURI(dirGroupIP.DirGroupIPID(), DirGroupType), dirGroupIP, target)

	if err != nil {
		return res, errors.UpdateError(serviceName, dirGroupIPName, err)
	}

	return res, nil
}

func (s *Service) PartialUpdate(dirGroupIP *DirGroupIP) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupIPName := url.PathEscape(dirGroupIP.Name)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, helper.GetDirGroupURI(dirGroupIP.DirGroupIPID(), DirGroupType), dirGroupIP, target)

	if err != nil {
		return res, errors.PartialUpdateError(serviceName, dirGroupIPName, err)
	}

	return res, nil
}

func (s *Service) Delete(dirGroupID string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, helper.GetDirGroupURI(dirGroupID, DirGroupType), nil, target)

	if err != nil {
		return res, errors.DeleteError(serviceName, dirGroupID, err)
	}

	return res, nil
}

func (s *Service) List(queryInfo *helper.QueryInfo, dirGroupIP *DirGroupIP) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, helper.GetDirGroupListURI(dirGroupIP.AccountName, DirGroupType), nil, target)

	if err != nil {
		return res, nil, errors.ListError(serviceName, helper.GetDirGroupListURI(dirGroupIP.AccountName, DirGroupType), err)
	}

	dirGroupIPListResponse := target.Data.(*ResponseList)

	return res, dirGroupIPListResponse, nil
}
