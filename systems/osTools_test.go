package systems

import "testing"

func TestGetMinVer(t *testing.T) {
	tests := []struct {
		name    string
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMinVer()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinVer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMinVer() = %v, want %v", got, tt.want)
			}
		})
	}
}
