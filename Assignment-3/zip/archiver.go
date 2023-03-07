package zip

import (
	"io"
)

type Archiver interface {
	Archive(name []string, readers ...io.Reader) (io.Reader, error)
}

func New() Archiver {
	return &ZipArchiver{}
}
