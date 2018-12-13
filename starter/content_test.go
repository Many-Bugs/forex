package starter

import (
	"testing"
)

func TestGetFieldStructPointer(t *testing.T) {
	c := &Content{
		App: App{
			ModuleID: 100000,
		},
	}
	type args struct {
		obj  interface{}
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "A",
			args: args{
				name: "App",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.GetFieldStructPointer(tt.args.name)
		})
	}
}

func TestInspectStruct(t *testing.T) {
	type args struct {
		o interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "A",
			args: args{
				o: &Content{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InspectStruct(tt.args.o)
		})
	}
}
