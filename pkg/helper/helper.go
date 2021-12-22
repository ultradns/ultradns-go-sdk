package helper

import (
	"fmt"
	"strings"
)

func AppendRootDot(zoneName string) string {
	return fmt.Sprintf("%s.", zoneName)
}

func GetZoneFQDN(zoneName string) string {
	if len(zoneName) > 0 {
		if lastChar := zoneName[len(zoneName)-1]; lastChar != '.' {
			return AppendRootDot(zoneName)
		}
	}

	return zoneName
}

func GetOwnerFQDN(name, zone string) string {
	if !strings.Contains(name, zone) {
		name = AppendRootDot(name) + GetZoneFQDN(zone)
	}

	return GetZoneFQDN(name)
}

func GetRRTypeFullString(key string) string {
	var rrTypes = map[string]string{
		"A":         "A (1)",
		"1":         "A (1)",
		"NS":        "NS (2)",
		"2":         "NS (2)",
		"CNAME":     "CNAME (5)",
		"5":         "CNAME (5)",
		"SOA":       "SOA (6)",
		"6":         "SOA (6)",
		"PTR":       "PTR (12)",
		"12":        "PTR (12)",
		"HINFO":     "HINFO (13)",
		"13":        "HINFO (13)",
		"MX":        "MX (15)",
		"15":        "MX (15)",
		"TXT":       "TXT (16)",
		"16":        "TXT (16)",
		"RP":        "RP (17)",
		"17":        "RP (17)",
		"AAAA":      "AAAA (28)",
		"28":        "AAAA (28)",
		"SRV":       "SRV (33)",
		"33":        "SRV (33)",
		"NAPTR":     "NAPTR (35)",
		"35":        "NAPTR (35)",
		"DS":        "DS (43)",
		"43":        "DS (43)",
		"SSHFP":     "SSHFP (44)",
		"44":        "SSHFP (44)",
		"TLSA":      "TLSA (52)",
		"52":        "TLSA (52)",
		"SPF":       "SPF (99)",
		"99":        "SPF (99)",
		"CAA":       "CAA (257)",
		"257":       "CAA (257)",
		"APEXALIAS": "APEXALIAS (65282)",
		"65282":     "APEXALIAS (65282)",
	}

	return rrTypes[key]
}
