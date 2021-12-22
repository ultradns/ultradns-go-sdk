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
