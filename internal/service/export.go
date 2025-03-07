package service

type Exporter interface {
	Export(dir string) error
}
