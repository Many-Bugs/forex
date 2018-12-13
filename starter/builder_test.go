package starter

import (
	"reflect"
	"testing"
)

func TestDefaultBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *Content
	}{
		// TODO: Add test cases.
		{
			name: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
