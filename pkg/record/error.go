package record

import (
	"fmt"
)

func CreateRecordError(recordID string, err error) error {
	return fmt.Errorf("error while creating record - %s : %w", recordID, err)
}

func UpdateRecordError(recordID string, err error) error {
	return fmt.Errorf("error while updating record - %s : %w", recordID, err)
}

func PartialUpdateRecordError(recordID string, err error) error {
	return fmt.Errorf("error while partial updating record - %s : %w", recordID, err)
}

func ReadRecordError(recordID string, err error) error {
	return fmt.Errorf("error while reading record - %s : %w", recordID, err)
}

func DeleteRecordError(recordID string, err error) error {
	return fmt.Errorf("error while deleting record - %s : %w", recordID, err)
}
