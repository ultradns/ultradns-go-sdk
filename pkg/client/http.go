package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const contentType = "application/json"
const throttleSleep = 1 * time.Second
const maxThrottleRetry = 3
const maxDecodePreviewBytes = 512

var (
	defaultUserAgent = version.GetSDKVersion()
)

func (c *Client) Do(method, path string, payload interface{}, target *Response) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, path)
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

	userAgent := defaultUserAgent + ";" + c.userAgent

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Accept", contentType)
	req.Header.Add("User-Agent", userAgent)

	c.logHttpRequest(req)
	res, err := c.httpClient.Do(req)
	c.logHttpResponse(res)

	resp := &http.Response{}

	if res != nil {
		resp.Status = res.Status
		resp.StatusCode = res.StatusCode
		resp.Header = res.Header
	}

	if target == nil {
		return resp, errors.ResponseTargetError("<nil>")
	}

	if resp.StatusCode == http.StatusTooManyRequests && target.retry < maxThrottleRetry {
		c.Warn("Throttling the request: '%s %s': attempt=%v ", method, url, (target.retry + 1))
		target.retry += 1
		time.Sleep(throttleSleep)
		return c.Do(method, path, payload, target)
	}

	if err != nil {
		return resp, err
	}

	defer res.Body.Close()

	er := c.validateResponse(res, target)

	if er != nil {
		return resp, er
	}

	return resp, nil
}

func (c *Client) validateResponse(res *http.Response, target *Response) error {

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		if res.StatusCode == http.StatusNoContent {
			return nil
		}

		reader := bufio.NewReader(res.Body)
		if _, err := reader.Peek(1); err == io.EOF {
			if _, ok := target.Data.(*SuccessResponse); ok {
				return nil
			}
			return fmt.Errorf("empty response body with status %d (%s)", res.StatusCode, res.Status)
		} else if err != nil {
			return err
		}

		decodeReader := io.Reader(reader)
		previewBuf := &limitedPreviewBuffer{max: maxDecodePreviewBytes}
		if c.logger.logLevel >= LogDebug {
			decodeReader = io.TeeReader(reader, previewBuf)
		}

		err := json.NewDecoder(decodeReader).Decode(&target.Data)
		if err != nil {
			if err == io.EOF {
				if _, ok := target.Data.(*SuccessResponse); ok {
					return nil
				}
				return fmt.Errorf("empty response body with status %d (%s)", res.StatusCode, res.Status)
			}

			if c.logger.logLevel >= LogDebug {
				preview := previewBuf.String()
				preview = strings.ReplaceAll(preview, "\n", "\\n")
				return fmt.Errorf("unable to decode success response (status %d): %w; body=%q", res.StatusCode, err, preview)
			}

			return fmt.Errorf("unable to decode success response (status %d): %w", res.StatusCode, err)
		}
	} else {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&target.ErrorList)
		if err == nil {
			return errors.APIResponseError(target.ErrorList[0].String())
		}

		c.Warn("Unable to parse API error message: %s", err.Error())

		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&target.Error)
		if err == nil {
			return errors.APIResponseError(target.Error.String())
		}

		return err
	}

	return nil
}

type limitedPreviewBuffer struct {
	buf bytes.Buffer
	max int
}

func (b *limitedPreviewBuffer) Write(p []byte) (int, error) {
	originalLen := len(p)
	remaining := b.max - b.buf.Len()
	if remaining > 0 {
		if len(p) > remaining {
			p = p[:remaining]
		}
		_, _ = b.buf.Write(p)
	}

	return originalLen, nil
}

func (b *limitedPreviewBuffer) String() string {
	return b.buf.String()
}
