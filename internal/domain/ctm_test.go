package domain

import (
	"reflect"
	"testing"
)

func TestNewCTMProps(t *testing.T) {
	type args struct {
		method     string
		startTile  int
		endTile    int
		additional map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "empty",
			args: args{method: "ctm", startTile: 0, endTile: 0, additional: map[string]string{}},
			want: map[string]string{"method": "ctm", "tiles": "0-0"},
		},
		{
			name:    "start tile greater than end tile",
			args:    args{method: "ctm", startTile: 1, endTile: 0, additional: map[string]string{}},
			wantErr: true,
		},
		{
			name: "repeat method",
			args: args{method: "repeat", startTile: 0, endTile: 15, additional: map[string]string{"width": "4", "height": "4"}},
			want: map[string]string{"method": "repeat", "tiles": "0-15", "width": "4", "height": "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCTMProps(tt.args.method, tt.args.startTile, tt.args.endTile, tt.args.additional)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCTMProps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetProps(), tt.want) {
				t.Errorf("NewCTMProps() got = %v, want %v", got.GetProps(), tt.want)
			}
		})
	}
}
