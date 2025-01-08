package utils

import (
	"encoding/json"
	"fmt"
)

// ParseJSON parses a JSON string into a struct
func ParseJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	return nil
}

// ToJSON converts a struct to a JSON string
func ToJSON(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to JSON: %w", err)
	}
	return data, nil
}

// JSONDiff compares two JSON objects and returns the differences
func JSONDiff(json1, json2 []byte) (differences []string, err error) {
	var obj1, obj2 map[string]interface{}
	if err := json.Unmarshal(json1, &obj1); err != nil {
		return nil, fmt.Errorf("failed to parse JSON1: %w", err)
	}
	if err := json.Unmarshal(json2, &obj2); err != nil {
		return nil, fmt.Errorf("failed to parse JSON2: %w", err)
	}

	for key, value1 := range obj1 {
		if value2, exists := obj2[key]; exists {
			if value1 != value2 {
				differences = append(differences, fmt.Sprintf("Key '%s' differs: %v vs %v", key, value1, value2))
			}
		} else {
			differences = append(differences, fmt.Sprintf("Key '%s' missing in second JSON", key))
		}
	}

	for key := range obj2 {
		if _, exists := obj1[key]; !exists {
			differences = append(differences, fmt.Sprintf("Key '%s' missing in first JSON", key))
		}
	}

	return differences, nil
}
