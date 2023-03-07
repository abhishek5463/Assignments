package zip

import (
	"archive/zip"
	"io"
	"log"
)

type ZipArchiver struct {
}

func CreateNewZipArchiver() *ZipArchiver {
	return &ZipArchiver{}
}
func (zipper *ZipArchiver) Archive(names []string, readers ...io.Reader) (io.Reader, error) {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		zipW := zip.NewWriter(pw)
		for i, reader := range readers {
			name := names[i]
			file, err := zipW.Create(name)
			if err != nil {
				log.Fatalf("Error in creating file %v", err)
			}
			_, err1 := io.Copy(file, reader)
			if err != nil {
				log.Fatalf("Error in copying file %v", err1)
			}
		}
		zipW.Close()
	}()

	return pr, nil
}
