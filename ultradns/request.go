/**
 * Copyright 2012-2013 NeuStar, Inc. All rights reserved. NeuStar, the Neustar logo and related names and logos are
 * registered trademarks, service marks or tradenames of NeuStar, Inc. All other product names, company names, marks,
 * logos and symbols may be trademarks of their respective owners.
 */
package ultradns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonPatch struct {
	Op    string `json:"op,omitempty"`
	Path  string `json:"path,omitempty"`
	Value string `json:"value,omitempty"`
	From  string `json:"from,omitempty"`
}

//sends http request to the provided path of ultradns api
//return http response
func (c *Client) Do(method, path string, payload, target interface{}) (*http.Response, error) {
	httpClient := c.httpClient
	url := fmt.Sprintf("%s/%s", c.baseUrl, path)

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.userAgent)

	if method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json-patch+json")
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if target != nil {
		target := target.(*Response)
		if res.StatusCode >= http.StatusOK && res.StatusCode <= http.StatusIMUsed {
			er := json.NewDecoder(res.Body).Decode(&target.Data)
			if er != nil && er.Error() == "EOF" {
				return res, nil
			} else if er != nil {
				return res, er
			}
		} else {
			er := json.NewDecoder(res.Body).Decode(&target.Error)
			if er != nil {
				return res, er
			}
		}
	}

	return res, nil
}
