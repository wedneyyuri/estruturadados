package estruturadados

import (
	"testing"
)

func TestContains(t *testing.T) {
	values := []string{"a", "b", "c", "d", "e"}

	for _, v := range values {
		if !Contains(values, v) {
			t.Fatalf("Contains returned false for %v", v)
		}
	}

	if Contains(values, "z") {
		t.Fatalf("Contains returned true for %v", "z")
	}
}

func BenchmarkContains1000000(b *testing.B) {
	values := genTerms(1000000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range values {
			if !Contains(values, v) {
				panic("invalid result")
			}
		}
	}
}
