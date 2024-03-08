package utils

// Function to check if a string slice contains a specific string.
func ContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
