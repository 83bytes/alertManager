package utils

import (
	"testing"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func TestStrictUnmarshal(t *testing.T) {
	validJSON := []byte(`{
		"name": "John Doe",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "Anytown"
		}
	}`)

	invalidJSON := []byte(`{
		"name": "John Doe",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "Anytown"
		},
		"unknownField": "value"
	}`)

	// this json has an unknown field in the inner struct
	invalidJSON_2 := []byte(`{
		"name": "John Doe",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "Anytown",
			"unknownInnerField": "value"
		}
	}`)

	var person Person

	err := StrictUnmarshal(validJSON, &person)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	err = StrictUnmarshal(invalidJSON, &person)
	if err == nil {
		t.Errorf("expected error, got none")
	}

	err = StrictUnmarshal(invalidJSON_2, &person)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}
