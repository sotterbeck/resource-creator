package domain

import "testing"

func TestPatternImage_IsValidRes(t *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{w: 0, h: 0}, want: false},
		{name: "half empty", args: args{w: 0, h: 1}, want: false},
		{name: "negative", args: args{w: -1, h: -1}, want: false},
		{name: "not power of 2", args: args{w: 3, h: 3}, want: false},
		{name: "power of 2", args: args{w: 2, h: 2}, want: true},
		{name: "power of 2 not square", args: args{w: 2, h: 4}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PatternImage{}
			if got := p.IsValidRes(tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("isValidRes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTileImage_IsValidRes(t1 *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{w: 0, h: 0}, want: false},
		{name: "not power of 2", args: args{w: 15, h: 15}, want: false},
		{name: "power of 2 not square", args: args{w: 2, h: 4}, want: false},
		{name: "power of 2 square", args: args{w: 2, h: 2}, want: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TileImage{}
			if got := t.IsValidRes(tt.args.w, tt.args.h); got != tt.want {
				t1.Errorf("IsValidRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
