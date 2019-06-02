package estruturadados

import "strings"

// ContentOnBlackList returns whether a term is blacklisted or not.
func ContentOnBlackList(searchBlacklist []string, content string) bool {
	contentWords := strings.Split(Slugfy(content), "-")

	for _, term := range searchBlacklist {
		termWords := strings.Split(term, " ")
		if ContainsAllIn(contentWords, termWords) {
			return true
		}
	}
	return false
}

// Contains if contains string in array of string
func Contains(arr []string, s string) bool {
	for _, str := range arr {
		if str == s {
			return true
		}
	}

	return false
}

// ContainsAllIn - compare all words to contains on list
func ContainsAllIn(wordsList, wordsToContais []string) bool {
	if len(wordsToContais) < 1 {
		return false
	}

	contains := []bool{}
	wordsToContaisLen := len(wordsToContais)
	for i, word := range wordsToContais {
		if Contains(wordsList, word) {
			contains = append(contains, true)
			if len(wordsToContais) > i {
				wordsToContais = wordsToContais[i+1:]
			}
		}
	}

	if len(contains) == wordsToContaisLen {
		return true
	}

	return false
}
