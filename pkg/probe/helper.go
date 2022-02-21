package probe

import (
	"fmt"

	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/dns"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ftp"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/http"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ping"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/smtp"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/smtpsend"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/tcp"
)

func ValidateProbeDetails(probeData *Probe) error {
	if expected, returned := getProbeDetailsType(probeData.Type), fmt.Sprintf("%T", probeData.Details); expected != returned {
		return errors.TypeMismatchError(expected, returned)
	}

	return nil
}

func getProbeDetailsType(probeType string) string {
	var probeDetailsType = map[string]string{
		HTTP:     "*http.Details",
		TCP:      "*tcp.Details",
		FTP:      "*ftp.Details",
		SMTP:     "*smtp.Details",
		SMTPSend: "*smtpsend.Details",
		PING:     "*ping.Details",
		DNS:      "*dns.Details",
	}

	return probeDetailsType[probeType]
}

func getProbeDetails(probeType string) interface{} {
	switch probeType {
	case HTTP:
		return &http.Details{}
	case TCP:
		return &tcp.Details{}
	case FTP:
		return &ftp.Details{}
	case SMTP:
		return &smtp.Details{}
	case PING:
		return &ping.Details{}
	case DNS:
		return &dns.Details{}
	case SMTPSend:
		return &smtpsend.Details{}
	}

	return nil
}
