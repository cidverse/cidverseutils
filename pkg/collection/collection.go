package collection

import "reflect"

// InArray checks if a slice contains a element
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// MapGetValueOrDefault returns a value in a map or a default value
func MapGetValueOrDefault(entity map[string]string, key string, defaultValue string) (val string) {
	value, found := entity[key]

	if found {
		return value
	}

	return defaultValue
}

// MapHasKey checks if a map contains a key
func MapHasKey(inputMap map[string]interface{}, key string) bool {
	_, hasKey := inputMap[key]
	return hasKey
}