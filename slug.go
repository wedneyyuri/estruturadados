package estruturadados

import "strings"

// Slugfy returns the canonical value.
func Slugfy(v string) string {
	v = strings.TrimSpace(v)
	v = strings.ToLower(v)
	return strings.ReplaceAll(v, " ", "-")
}
