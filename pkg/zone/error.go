package zone

import "fmt"

func CreateZoneError(zonename string, err error) error {
	return fmt.Errorf("error while creating zone - %s : %w", zonename, err)
}

func UpdateZoneError(zonename string, err error) error {
	return fmt.Errorf("error while updating zone - %s : %w", zonename, err)
}

func PartialUpdateZoneError(zonename string, err error) error {
	return fmt.Errorf("error while partial updating zone - %s : %w", zonename, err)
}

func ReadZoneError(zonename string, err error) error {
	return fmt.Errorf("error while reading zone - %s : %w", zonename, err)
}

func DeleteZoneError(zonename string, err error) error {
	return fmt.Errorf("error while deleting zone - %s : %w", zonename, err)
}

func ListZoneError(uri string, err error) error {
	return fmt.Errorf("error while listing zone : path and query params - %s : %w", uri, err)
}
