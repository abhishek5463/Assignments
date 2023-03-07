package downloader

import (
	"assignment-3/downloader/file"
	"assignment-3/downloader/web"
	"io"
	"strings"
)

type Downloader interface {
	Download() (io.Reader, error)
}

func NewDownlaoder(urlPath string) Downloader {
	if strings.HasPrefix(urlPath, "http") {
		return web.CreateNewUrlDownloader(urlPath)
	}
	return file.CreateNewFileSyastemDownloader(urlPath)
}
