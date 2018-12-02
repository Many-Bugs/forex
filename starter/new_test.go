package starter

import (
	"reflect"
	"testing"
)

func TestDefault(t *testing.T) {
	tests := []struct {
		name string
		want *Content
	}{
		// TODO: Add test cases.
		{
			name: "test feature",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
