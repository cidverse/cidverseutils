package hash

import (
	"testing"
)

func TestUUIDNoDash(t *testing.T) {
	uuidWithDash := "550e8400-e29b-41d4-a716-446655440000"
	expectedUUID := "550e8400e29b41d4a716446655440000"
	result := UUIDNoDash(uuidWithDash)
	if result != expectedUUID {
		t.Errorf("Expected UUIDNoDash to return '%s', but got '%s'", expectedUUID, result)
	}
}
