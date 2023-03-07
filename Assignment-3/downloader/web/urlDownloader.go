package web

import (
	"io"
	"log"
	"net/http"
)

type UrlDownloader struct {
	uri string
}

func CreateNewUrlDownloader(uriPath string) *UrlDownloader {
	return &UrlDownloader{uri: uriPath}
}
func (downloader *UrlDownloader) Download() (io.Reader, error) {
	resp, err := http.Get(downloader.uri)
	if err != nil {
		log.Fatalf("file not downloaded %v", err)
	}
	r, w := io.Pipe()
	go func() {
		defer resp.Body.Close()
		defer w.Close()
		_, err := io.Copy(w, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return r, nil
}
