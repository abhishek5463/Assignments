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

	//sample URLs for videos
	url1 := "https://filesamples.com/samples/video/mp4/sample_960x400_ocean_with_audio.mp4"
	url2 := "https://filesamples.com/samples/video/mp4/sample_1280x720_surfing_with_audio.mp4"

	//creating oject for downloader
	downloader1 := downloader.NewDownlaoder(url1)
	downloader2 := downloader.NewDownlaoder(url2)

	//starting time stamp
	startTime := time.Now()

	//creating reader of first file
	r1, err := downloader1.Download()
	if err != nil {
		log.Fatal(err)
	}

	//creating reader of scond file
	r2, err := downloader2.Download()
	if err != nil {
		log.Fatal(err)
	}

	//new onject of zip
	zipper := zip.New()
	r3, err := zipper.Archive([]string{"f1.mp4", "f2.mp4"}, r1, r2)
	if err != nil {
		log.Fatal(err)
	}

	//creating zip file
	file, err := os.Create("result.zip")
	if err != nil {
		log.Fatal(err)
	}

	//copying data from zip reader
	size, err1 := io.Copy(file, r3)
	if err != nil {
		log.Fatalf("Error in file copy %v :", err1)
	}

	//printing amount of data we have copied and time taken
	fmt.Printf("Number of Mbs downloaded: %d Mb Time taken: %v\n", size/1000000, time.Since(startTime))
}
