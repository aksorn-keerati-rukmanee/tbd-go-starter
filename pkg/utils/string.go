package utils

import (
	"encoding/json"
	"log"
)

// Function to check if a string slice contains a specific string.
func ContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func JSONPrettyString(val interface{}) string {
	jsonData, err := json.MarshalIndent(val, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}
