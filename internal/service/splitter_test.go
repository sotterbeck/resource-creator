package service

import (
	"image"
	"reflect"
	"resource-creator/internal/domain"
	"testing"
)

func TestImageSplitter_SplitImage(t *testing.T) {
	type fields struct {
		validator domain.Image
	}
	type args struct {
		img image.Image
		res int
	}
	f := fields{validator: &domain.PatternImage{}}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []image.Image
		wantErr bool
	}{
		{
			name:    "invalid image resolution",
			fields:  f,
			args:    args{img: image.NewRGBA(image.Rect(0, 0, 6, 6)), res: 4},
			wantErr: true,
		},
		{
			name:    "split resolution negative",
			fields:  f,
			args:    args{img: image.NewRGBA(image.Rect(0, 0, 8, 8)), res: -4},
			wantErr: true,
		},
		{
			name:    "split resolution zero",
			fields:  f,
			args:    args{img: image.NewRGBA(image.Rect(0, 0, 8, 8)), res: 0},
			wantErr: true,
		},
		{
			name:    "split resolution larger than image",
			fields:  f,
			args:    args{img: image.NewRGBA(image.Rect(0, 0, 8, 8)), res: 16},
			wantErr: true,
		},
		{
			name:    "split resolution not divisible",
			fields:  f,
			args:    args{img: image.NewRGBA(image.Rect(0, 0, 8, 8)), res: 3},
			wantErr: true,
		},
		{
			name:   "one split",
			fields: f,
			args:   args{img: image.NewRGBA(image.Rect(0, 0, 8, 8)), res: 8},
			want:   []image.Image{image.NewRGBA(image.Rect(0, 0, 8, 8))},
		},
		{
			name:   "two splits",
			fields: f,
			args:   args{img: image.NewRGBA(image.Rect(0, 0, 8, 4)), res: 4},
			want: func() []image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 8, 4))
				return []image.Image{
					img.SubImage(image.Rect(0, 0, 4, 4)).(image.Image),
					img.SubImage(image.Rect(4, 0, 8, 4)).(image.Image),
				}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ImageSplitter{
				validator: tt.fields.validator,
			}
			got, err := s.SplitImage(tt.args.img, tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
