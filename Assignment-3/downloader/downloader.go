package downloader

import (
	"assignment-3/downloader/file"
	"assignment-3/downloader/web"
	"io"
	"strings"
)

//Downloader interface 
type Downloader interface {
	Download() (io.Reader, error)
}


// Recognising which downloader to call
func NewDownlaoder(urlPath string) Downloader {
	if strings.HasPrefix(urlPath, "http") {
		return web.CreateNewUrlDownloader(urlPath)
	}
	return file.CreateNewFileSyastemDownloader(urlPath)
}
