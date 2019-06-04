package estruturadados

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func readFile(filename string) []string {
	path := filepath.Join("testdata", filename)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	blacklist := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		blacklist = append(blacklist, text)
	}
	return blacklist
}

func TestContentOnBlackList(t *testing.T) {
	blacklist := []string{"lula livre", "michel temer"}

	tests := []struct {
		term      string
		blacklist bool
	}{
		{term: "camisa lula livre", blacklist: true},
		{term: "lula light ao alho e óleo", blacklist: false},
		{term: "Michel Miguel Elias Temer Lulia - Biografia", blacklist: true},
		{term: "Michel Teló - Baile do Teló - DVD + CD", blacklist: false},
	}

	for _, tt := range tests {
		t.Run(tt.term, func(t *testing.T) {
			if got := ContentOnBlackList(blacklist, tt.term); got != tt.blacklist {
				t.Errorf("Slugfy() = %v, want %v", got, tt.blacklist)
			}
		})
	}
}

func BenchmarkContentOnBlackList(b *testing.B) {
	searchTerms := readFile("google-trends-2018.txt")
	blacklistTerms := readFile("full-list-of-bad-words_text-file_2018_07_30.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, term := range searchTerms {
			ContentOnBlackList(blacklistTerms, term)
		}
	}
}
