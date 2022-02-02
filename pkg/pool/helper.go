package pool

import (
	"fmt"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

var (
	monitorMethod = map[string]bool{
		"GET":  true,
		"POST": true,
	}
	regionFailureSensitivity = map[string]bool{
		"HIGH": true,
		"LOW":  true,
	}
)

func ValidatePoolProfile(expProfileType string, rrSet *rrset.RRSet) error {
	rrsetProfileType := fmt.Sprintf("%T", rrSet.Profile)

	if expProfileType != rrsetProfileType {
		return helper.TypeMismatchError(expProfileType, rrsetProfileType)
	}

	rrSet.Profile.SetContext()

	return nil
}

func IsRegionFailureSensitivityValid(val string) bool {
	return IsValidField(val, regionFailureSensitivity)
}

func IsMonitorMethodValid(val string) bool {
	return IsValidField(val, monitorMethod)
}

func IsValidField(val string, dataMap map[string]bool) bool {
	_, ok := dataMap[val]

	return ok
}
