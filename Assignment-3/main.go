package main

import (
	"assignment-3/downloader"
	"assignment-3/zip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	url1 := "/Users/abhishek.ak/Downloads/untitled folder/i-card.pdf"
	url2 := "https://filesamples.com/samples/video/mp4/sample_1280x720_surfing_with_audio.mp4"
	downloader1 := downloader.NewDownlaoder(url1)
	downloader2 := downloader.NewDownlaoder(url2)

	startTime := time.Now()

	r1, err := downloader1.Download()
	if err != nil {
		log.Fatal(err)
	}
	r2, err := downloader2.Download()
	if err != nil {
		log.Fatal(err)
	}
	zipper := zip.New()
	r3, err := zipper.Archive([]string{"f1.pdf", "f2.mp4"}, r1, r2)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("result.zip")
	if err != nil {
		log.Fatal(err)
	}
	size, err1 := io.Copy(file, r3)
	if err != nil {
		log.Fatalf("Error in file copy %v :", err1)
	}
	fmt.Printf("Number of Mbs downloaded: %d Mb Time taken: %v\n", size/1000000, time.Since(startTime))
}
