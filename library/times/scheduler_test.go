package times

import "testing"

func TestParseAny(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		// TODO: Add test cases.
		{name: "A", value: "20160203"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseAny(tt.value)
		})
	}
}
