package estruturadados

import (
	"math"
)

// Intersection between two slices.
func Intersection(a []string, b []string) []string {
	cap := int(math.Min(float64(len(a)), float64(len(b))))
	result := make([]string, 0, cap)
	seen := make(map[string]bool, cap)

	for _, v := range a {
		seen[v] = true
	}

	for _, v := range b {
		if !seen[v] {
			continue
		}
		result = append(result, v)
	}

	return result
}
