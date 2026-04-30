package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDoAllowsEmptyBodyForSuccessResponseTarget(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := c.Do(http.MethodGet, "", nil, Target(&SuccessResponse{}))
	if err != nil {
		t.Fatalf("expected nil error for empty body with SuccessResponse target, got %v", err)
	}
}

func TestDoRejectsEmptyBodyForNonSuccessResponseTarget(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := c.Do(http.MethodGet, "", nil, Target(&struct{}{}))
	if err == nil {
		t.Fatal("expected error for empty body with non-SuccessResponse target, got nil")
	}

	expected := "empty response body with status 200 (200 OK)"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}
}

func TestDoAllowsWhitespaceOnlyBodyForSuccessResponseTarget(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("  \n\t  "))
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := c.Do(http.MethodGet, "", nil, Target(&SuccessResponse{}))
	if err != nil {
		t.Fatalf("expected nil error for whitespace-only body with SuccessResponse target, got %v", err)
	}
}

func TestDoRejectsWhitespaceOnlyBodyForNonSuccessResponseTarget(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("  \n\t  "))
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := c.Do(http.MethodGet, "", nil, Target(&struct{}{}))
	if err == nil {
		t.Fatal("expected error for whitespace-only body with non-SuccessResponse target, got nil")
	}

	expected := "empty response body with status 200 (200 OK)"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}
}

func TestDoMalformedSuccessBodyNoPreviewByDefault(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<html>gateway error</html>"))
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := c.Do(http.MethodGet, "", nil, Target(&struct{}{}))
	if err == nil {
		t.Fatal("expected decode error for malformed success JSON body, got nil")
	}

	if !strings.Contains(err.Error(), "unable to decode success response (status 200)") {
		t.Fatalf("expected decode context in error, got %q", err.Error())
	}
	if strings.Contains(err.Error(), "body=") {
		t.Fatalf("did not expect body preview by default, got %q", err.Error())
	}
}

func TestDoMalformedSuccessBodyIncludesPreviewWhenDebugEnabled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<html>gateway error</html>"))
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}
	c.EnableDefaultDebugLogger()

	_, err := c.Do(http.MethodGet, "", nil, Target(&struct{}{}))
	if err == nil {
		t.Fatal("expected decode error for malformed success JSON body, got nil")
	}

	if !strings.Contains(err.Error(), "unable to decode success response (status 200)") {
		t.Fatalf("expected decode context in error, got %q", err.Error())
	}
	if !strings.Contains(err.Error(), "body=\"<html>gateway error</html>\"") {
		t.Fatalf("expected body preview in error when debug is enabled, got %q", err.Error())
	}
}

func TestDoMalformedSuccessBodyPreviewIsBoundedWhenDebugEnabled(t *testing.T) {
	largeBody := "<" + strings.Repeat("a", maxDecodePreviewBytes+100) + ">"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(largeBody))
	}))
	defer server.Close()

	c := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}
	c.EnableDefaultDebugLogger()

	_, err := c.Do(http.MethodGet, "", nil, Target(&struct{}{}))
	if err == nil {
		t.Fatal("expected decode error for malformed success JSON body, got nil")
	}

	if !strings.Contains(err.Error(), "body=") {
		t.Fatalf("expected bounded body preview in error when debug is enabled, got %q", err.Error())
	}
	if strings.Contains(err.Error(), strings.Repeat("a", maxDecodePreviewBytes+50)) {
		t.Fatalf("expected preview to be capped at %d bytes, got %q", maxDecodePreviewBytes, err.Error())
	}
}
