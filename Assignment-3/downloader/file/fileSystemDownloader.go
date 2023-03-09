package file

import (
	"io"
	"log"
	"os"
)

type FileSystemDownloader struct {
	uri string
}

func CreateNewFileSyastemDownloader(uriPath string) *FileSystemDownloader {
	return &FileSystemDownloader{uri: uriPath}
}

//implementation of Downloader for downloading from local 

func (downlaoder *FileSystemDownloader) Download() (io.Reader, error) {
	reader, err := os.Open(downlaoder.uri)
	if err != nil {
		log.Fatalf("file not found %v", err)
	}
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		defer reader.Close()

		_, err := io.Copy(w, reader)
		if err != nil {
			log.Fatalf("file not copied %v", err)
		}
	}()
	return r, nil
}
