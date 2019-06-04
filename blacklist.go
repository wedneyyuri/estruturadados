package estruturadados

import "strings"

// ContentOnBlackList returns whether a term is blacklisted or not.
func ContentOnBlackList(searchBlacklist []string, content string) bool {
	contentWords := strings.Split(Slugfy(content), "-")
	_ = contentWords

	return false
}
