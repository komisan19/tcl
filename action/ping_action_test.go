package action

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPingAction(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "success",
			args:    args{"https://example.com"},
			want:    200,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PingAction(tt.args.domain)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("missmatch = (-got, +want):\n%s", diff)
			}
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
