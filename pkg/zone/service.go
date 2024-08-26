package zone

import (
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
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

func (s *Service) CreateZone(zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s create started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPost, basePath, zone, target)

	if err != nil {
		zoneName := ""
		if zone.Properties != nil {
			zoneName = zone.Properties.Name
		}

		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.CreateError(serviceName, zoneName, err)
	}

	if er := s.checkZoneTask(res); er != nil {
		s.c.Error("[%s] %s create failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return nil, er
	}

	s.c.Trace("[%s] %s create completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) ReadZone(zoneName string) (*http.Response, *Response, error) {
	target := client.Target(&Response{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s read started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, basePath+zoneName, nil, target)

	if err != nil {
		s.c.Error("[%s] %s read failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ReadError(serviceName, zoneName, err)
	}

	zoneResponse := target.Data.(*Response)

	s.c.Trace("[%s] %s read completed successfully", version.GetSDKVersion(), serviceName)

	return res, zoneResponse, nil
}

func (s *Service) UpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s update started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPut, basePath+zoneName, zone, target)

	if err != nil {
		s.c.Error("[%s] %s update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.UpdateError(serviceName, zoneName, err)
	}

	s.c.Trace("[%s] %s update completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) PartialUpdateZone(zoneName string, zone *Zone) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s partial update started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodPatch, basePath+zoneName, zone, target)

	if err != nil {
		s.c.Error("[%s] %s partial update failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.PartialUpdateError(serviceName, zoneName, err)
	}

	s.c.Trace("[%s] %s partial update completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) DeleteZone(zoneName string) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	zoneName = url.PathEscape(zoneName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s delete started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodDelete, basePath+zoneName, nil, target)

	if err != nil {
		s.c.Error("[%s] %s delete failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, errors.DeleteError(serviceName, zoneName, err)
	}

	s.c.Trace("[%s] %s delete completed successfully", version.GetSDKVersion(), serviceName)

	return res, nil
}

func (s *Service) ListZone(queryInfo *helper.QueryInfo) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("[%s] %s list started", version.GetSDKVersion(), serviceName)

	res, err := s.c.Do(http.MethodGet, basePathForList+queryInfo.URI(), nil, target)

	if err != nil {
		s.c.Error("[%s] %s list failed with error: %v", version.GetSDKVersion(), serviceName, err)
		return res, nil, errors.ListError(serviceName, basePathForList+queryInfo.URI(), err)
	}

	zoneListResponse := target.Data.(*ResponseList)

	s.c.Trace("[%s] %s list completed successfully", version.GetSDKVersion(), serviceName)

	return res, zoneListResponse, nil
}

func (s *Service) checkZoneTask(res *http.Response) error {
	s.c.Trace("[%s] check zone %s started", version.GetSDKVersion(), serviceName)
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

	s.c.Trace("[%s] check zone %s completed successfully", version.GetSDKVersion(), serviceName)

	return nil
}
