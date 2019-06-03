package estruturadados

import "testing"

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
