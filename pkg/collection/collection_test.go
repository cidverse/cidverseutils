package collection

import (
	"testing"
)

func TestInArray(t *testing.T) {
	// Test case 1: Check if the method returns true and the correct index when the value is present in the slice
	array := []int{1, 2, 3, 4, 5}
	exists, index := InArray(3, array)
	if exists != true || index != 2 {
		t.Errorf("Expected exists to be true and index to be 2, got exists=%v and index=%v", exists, index)
	}

	// Test case 2: Check if the method returns false and -1 when the value is not present in the slice
	exists, index = InArray(6, array)
	if exists != false || index != -1 {
		t.Errorf("Expected exists to be false and index to be -1, got exists=%v and index=%v", exists, index)
	}

	// Test case 3: Check if the method works with slices of other types (e.g. string)
	array2 := []string{"apple", "banana", "cherry"}
	exists, index = InArray("banana", array2)
	if exists != true || index != 1 {
		t.Errorf("Expected exists to be true and index to be 1, got exists=%v and index=%v", exists, index)
	}
}

func TestMapGetValueOrDefault(t *testing.T) {
	tests := []struct {
		entity       map[string]string
		key          string
		defaultValue string
		expected     string
	}{
		{map[string]string{"msg": "hello world"}, "a", "c", "c"},
		{map[string]string{"msg": "hello world"}, "msg", "c", "hello world"},
		{map[string]string{}, "b", "def", "def"},
	}

	for _, test := range tests {
		result := MapGetValueOrDefault(test.entity, test.key, test.defaultValue)
		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}
