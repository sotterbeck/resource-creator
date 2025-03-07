package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_isResourcePack(t *testing.T) {
	type content struct {
		name  string
		isDir bool
	}

	tests := []struct {
		name     string
		contents []content
		want     bool
		wantErr  bool
	}{
		{
			name:     "empty",
			contents: []content{},
			want:     false,
			wantErr:  false,
		},
		{
			name: "only mcmeta file",
			contents: []content{
				{name: "pack.mcmeta"},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "dir",
			contents: []content{
				{name: "pack.mcmeta"},
				{name: "assets", isDir: true},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			for _, file := range tt.contents {
				path := filepath.Join(tmpDir, file.name)
				var err error
				var f *os.File

				if file.isDir {
					err = os.MkdirAll(path, 0755)
				} else {
					f, err = os.Create(path)
				}
				if err != nil {
					t.Fatal(err)
				}
				f.Close()
			}

			got, err := IsResourcePack(tmpDir)

			if (err != nil) != tt.wantErr {
				t.Errorf("isResourcePack() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("isResourcePack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetResourcePackDir(t *testing.T) {
	type args struct {
		typ       string
		namespace string
		material  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "model directory",
			args: args{
				typ:       AssetTypeModel,
				namespace: "minecraft",
				material:  "stone",
			},
			want: "assets/minecraft/models/block/stone",
		},
		{
			name: "blockstate directory",
			args: args{
				typ:       AssetTypeBlockState,
				namespace: "minecraft",
				material:  "stone",
			},
			want: "assets/minecraft/blockstates",
		},
		{
			name: "texture directory",
			args: args{
				typ:       AssetTypeTexture,
				namespace: "minecraft",
				material:  "stone",
			},
			want: "assets/minecraft/textures/block/stone",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			err := CreateResourcePackFiles(t, tmpDir)

			got, err := GetResourcePackDir(tmpDir, tt.args.typ, tt.args.namespace, tt.args.material)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResourcePackDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantAbs := filepath.Join(tmpDir, tt.want)

			if got != wantAbs {
				t.Errorf("GetResourcePackDir() got = %v, want %v", got, wantAbs)
			}
		})
	}
}
