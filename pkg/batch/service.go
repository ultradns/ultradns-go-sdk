package batch

import (
	"encoding/json"
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const (
	serviceName = "Batch"
	// basePath follows the API requirement that batch calls only run under
	// the /v1/ path. Per the REST API User Guide, Batch and Batch Query calls
	// can ONLY be run when using /v1/ in the API call.
	basePath = "v1/batch"
)

// Action wraps a single operation submitted to the batch API. The request
// body is submitted as a bare JSON array of Action objects.
type Action struct {
	Method string      `json:"method,omitempty"`
	URI    string      `json:"uri,omitempty"`
	Body   interface{} `json:"body,omitempty"`
}

// Response wraps the result of a single action in a batch response. The
// response body is a bare JSON array of Response objects, each carrying an
// individual HTTP status for the corresponding Action.
type Response struct {
	Status   int                   `json:"status,omitempty"`
	Response json.RawMessage       `json:"response,omitempty"`
	Error    *client.ErrorResponse `json:"error,omitempty"`
}

// Service wraps the Ultradns batch service.
type Service struct {
	c *client.Client
}

// New creates a batch service with a newly initialized client.
func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

// Get returns a batch service backed by the given client.
func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

// Execute submits a list of actions to the batch API as a single unit. The
// REST API processes the actions sequentially and rolls back the whole batch
// if any action fails.
//
// The batch API is async-capable (async=true with task polling) and, in the
// published API guide, is documented against a separate API host; neither is
// handled here. Verify reachability of the configured client's host for the
// /v1/batch path before relying on this in production.
func (s *Service) Execute(actions []*Action) (*http.Response, []*Response, error) {
	target := client.Target(&[]*Response{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s execute started", serviceName)

	res, err := s.c.Do(http.MethodPost, basePath, actions, target)

	if err != nil {
		s.c.Error("%s execute failed with error: %v", serviceName, err)
		return res, nil, errors.CreateError(serviceName, "batch", err)
	}

	batchResponse := target.Data.(*[]*Response)

	s.c.Trace("%s execute completed successfully", serviceName)

	return res, *batchResponse, nil
}
