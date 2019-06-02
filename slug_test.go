package estruturadados

import "testing"

func TestSlugfy(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "with spaces", input: "celular moto g", expected: "celular-moto-g"},
		{name: "with leading/trailling spaces", input: " celular ", expected: "celular"},
		{name: "with mixed caps", input: "Celular", expected: "celular"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slugfy(tt.input); got != tt.expected {
				t.Errorf("Slugfy() = %v, want %v", got, tt.expected)
			}
		})
	}
}
