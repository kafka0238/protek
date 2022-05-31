package util

func InArray(needle string, haystack []string) bool {
	for _, value := range haystack {
		if needle == value {
			return true
		}
	}
	return false
}
