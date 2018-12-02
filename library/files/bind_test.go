package files

import (
	"fmt"
	"testing"
)

func TestBindFileToObjINI(t *testing.T) {

	type A struct {
		A1 string
		A2 int
	}

	type B struct {
		B1 string
		B2 int
	}

	type C struct {
		C1 string
		C2 int
	}

	type Test struct {
		A A
		B B
		C C
	}

	var testObj Test

	type args struct {
		path string
		obj  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test feature",
			args: args{
				path: "testbind.ini",
				obj:  &testObj,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BindFileToObj(tt.args.path, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("BindFileToObj() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("obj value: ", testObj)
		})
	}
}
