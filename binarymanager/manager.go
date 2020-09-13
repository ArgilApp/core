package binarymanager

import (
	"log"
	"strings"

	"github.com/argilapp/core/binarymanager/provider"
)

var providers []provider.Provider

func AddProvider(provider provider.Provider) {
	log.Println("Loading providers")
	err := provider.Initialize()

	if err == nil {
		log.Println("Registered provider", provider.GetID())
		providers = append(providers, provider)
	} else {
		log.Println("Failed to register provider:", err.Error())
	}
}

func ListProviders() []provider.Provider {
	return providers
}

func HashExists(hash string) bool {
	var chunkedPath = chunkHashPath(hash)

	for _, p := range providers {
		var exists = p.FileExists(chunkedPath)

		if exists {
			return true
		}
	}
	return false
}

func chunkHashPath(hash string) string {
	var chunks []string
	runes := []rune(hash) // split the string into a slice of each individual character

	const chunkSize = 2

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}

	// only need the first half of the chunks to make a good enough path
	var halfChunks = chunks[:(len(chunks) / 2)]

	// join all the chunks together and then add the original hash back to the end of it
	//     do we only want the second half of the hash appended?
	var path = strings.Join(halfChunks, "/") + "/" + hash

	return path
}
