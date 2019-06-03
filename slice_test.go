package estruturadados

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		listA    []string
		listB    []string
		expected []string
	}{
		{
			name:     "Interseção",
			listA:    []string{"a", "b", "c"},
			listB:    []string{"b", "d", "c"},
			expected: []string{"b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection(tt.listA, tt.listB)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v got %v", tt.expected, got)
			}
		})
	}
}
