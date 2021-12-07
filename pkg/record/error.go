package record

import (
	"fmt"

	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func CreateRecordError(r rrset.RRSetKey, err error) error {
	return fmt.Errorf("error while creating record - %s : %w", r, err)
}

func UpdateRecordError(r rrset.RRSetKey, err error) error {
	return fmt.Errorf("error while updating record - %s : %w", r, err)
}

func PartialUpdateRecordError(r rrset.RRSetKey, err error) error {
	return fmt.Errorf("error while partial updating record - %s : %w", r, err)
}

func ReadRecordError(r rrset.RRSetKey, err error) error {
	return fmt.Errorf("error while reading record - %s : %w", r, err)
}

func DeleteRecordError(r rrset.RRSetKey, err error) error {
	return fmt.Errorf("error while deleting record - %s : %w", r, err)
}
