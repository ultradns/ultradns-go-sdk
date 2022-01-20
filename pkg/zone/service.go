package zone

import (
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/task"
)

const (
	serviceName     = "Zone"
	basePath        = "zones/"
	basePathForList = "v3/zones/"
	zoneTaskRetries = 5
	zoneTaskTimeGap = 10
	taskHeader      = "X-Task-Id"
)

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, helper.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) CreateZone(zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, basePath, zone, target)

	if err != nil {
		zoneName := ""
		if zone.Properties != nil {
			zoneName = zone.Properties.Name
		}

		return nil, helper.CreateError(serviceName, zoneName, err)
	}

	if er := s.checkZoneTask(res); er != nil {
		return nil, er
	}

	return res, nil
}

func (s *Service) ReadZone(zoneName string) (*http.Response, *Response, error) {
	target := client.Target(&Response{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, basePath+zoneName, nil, target)

	if err != nil {
		return nil, nil, helper.ReadError(serviceName, zoneName, err)
	}

	zoneResponse := target.Data.(*Response)

	return res, zoneResponse, nil
}

func (s *Service) UpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, basePath+zoneName, zone, target)

	if err != nil {
		return nil, helper.UpdateError(serviceName, zoneName, err)
	}

	return res, nil
}

func (s *Service) PartialUpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, basePath+zoneName, zone, target)

	if err != nil {
		return nil, helper.PartialUpdateError(serviceName, zoneName, err)
	}

	return res, nil
}

func (s *Service) DeleteZone(zoneName string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, basePath+zoneName, nil, target)

	if err != nil {
		return nil, helper.DeleteError(serviceName, zoneName, err)
	}

	return res, nil
}

func (s *Service) ListZone(queryInfo *helper.QueryInfo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, basePathForList+queryInfo.URI(), nil, target)

	if err != nil {
		return nil, nil, helper.ListError(serviceName, basePathForList+queryInfo.URI(), err)
	}

	zoneListResponse := target.Data.(*ResponseList)

	return res, zoneListResponse, nil
}

func (s *Service) checkZoneTask(res *http.Response) error {
	if res.StatusCode == http.StatusAccepted {
		taskID := res.Header.Get(taskHeader)
		taskService, err := task.Get(s.c)

		if err != nil {
			return err
		}

		er := taskService.TaskWait(taskID, zoneTaskRetries, zoneTaskTimeGap)

		if er != nil {
			return er
		}
	}

	return nil
}
