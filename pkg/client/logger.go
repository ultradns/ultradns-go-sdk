package client

import (
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
)

type logLevelType int

const (
	LogOff logLevelType = iota
	LogError
	LogDebug
	LogTrace
)

type logger struct {
	logLevel logLevelType
	logger   *log.Logger
}

func (l logger) getLogPrefix(logLevel logLevelType) string {
	switch logLevel {
	case LogError:
		return "[ERROR] "
	case LogDebug:
		return "[DEBUG] "
	case LogTrace:
		return "[TRACE] "
	}

	return ""
}

func (l logger) logHttpRequest(req *http.Request) {
	if l.logLevel > LogOff && l.logger != nil && req != nil {
		l.logger.SetPrefix(l.getLogPrefix(LogDebug))
		data, _ := httputil.DumpRequest(req, true)
		re := regexp.MustCompile(`\r?\n`)
		str := re.ReplaceAllString(string(data), "\r\n\t")
		l.logger.Printf("HTTP Request Sent:\n\t%s", str)
	}
}

func (l logger) logHttpResponse(res *http.Response) {
	if l.logLevel > LogOff && l.logger != nil && res != nil {
		l.logger.SetPrefix(l.getLogPrefix(LogDebug))
		data, _ := httputil.DumpResponse(res, true)
		re := regexp.MustCompile(`\r?\n`)
		str := re.ReplaceAllString(string(data), "\r\n\t")
		l.logger.Printf("HTTP Response Received:\n\t%s", str)
	}
}

func (l logger) logf(logType logLevelType, format string, v ...any) {
	if l.logLevel >= logType && l.logger != nil {
		l.logger.SetPrefix(l.getLogPrefix(logType))
		l.logger.Printf(format+"\n", v...)
	}
}

func (c *Client) Error(format string, v ...any) {
	c.logger.logf(LogError, format, v...)
}

func (c *Client) Debug(format string, v ...any) {
	c.logger.logf(LogDebug, format, v...)
}

func (c *Client) Trace(format string, v ...any) {
	c.logger.logf(LogTrace, format, v...)
}

func (c *Client) EnableDefaultDebugLogger() {
	c.EnableLogger(LogDebug, log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
}

func (c *Client) EnableDefaultTraceLogger() {
	c.EnableLogger(LogTrace, log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
}

func (c *Client) EnableLogger(logLevel logLevelType, flags int) {
	c.logger.logLevel = logLevel
	c.logger.logger = log.Default()
	c.logger.logger.SetFlags(flags)
}
