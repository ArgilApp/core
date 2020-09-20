package main

import (
	"log"
	"os"
	"time"

	"github.com/argilapp/core/binarymanager"
	"github.com/argilapp/core/binarymanager/provider"
)

func main() {
	var prov provider.Provider
	prov = &provider.LocalFileSystem{
		ID:               "localfs-test",
		StorageDirectory: "./tmp/localfs",
	}
	// prov = &provider.GoogleCloudStorage{
	// 	ID:     "gcs-test",
	// 	Bucket: "argil-user-content-test",
	// }
	binarymanager.AddProvider(prov)

	stream, err := os.Open("./tmp/example_file")
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	hashes, err := binarymanager.Upload(stream)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(hashes)

	var exists = binarymanager.HashExists(hashes)

	if exists {
		log.Println("Found hash", hashes.SHA256)

		tempDownloadFile, _ := os.Create("./tmp/downloaded")
		err = binarymanager.Download(hashes, tempDownloadFile)

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(3 * time.Second)
		delErr := binarymanager.Delete(hashes)
		if delErr != nil {
			log.Fatal(delErr)
		}
	} else {
		log.Fatalln("Could not find hash", hashes)
	}
}
