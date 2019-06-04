package estruturadados

import "strconv"

// Contains reports whether value is within values.
func Contains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func genTerms(size int) []string {
	terms := make([]string, size)
	for i := 0; i < size; i++ {
		terms[i] = strconv.Itoa(i)
	}
	return terms
}
