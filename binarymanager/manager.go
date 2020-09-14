package binarymanager

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/argilapp/core/binarymanager/provider"
	"golang.org/x/crypto/sha3"
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

func Upload(stream *os.File) (provider.Hashes, error) {
	const BufferSize = 100
	buffer := make([]byte, BufferSize)

	// Currently we assume the first provider is the one to write to
	uploadProvider := providers[0]

	var uploadFile = uploadProvider.CreateUploadHandle()

	// We calculate 4 different hashes for the file
	md5hash := md5.New()
	sha1hash := sha1.New()
	sha256hash := sha256.New()
	sha3hash := sha3.New256()

	// Read the stream to write to the provider & calculate the hashes
	for {
		bytesread, err := stream.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		readBuffer := buffer[:bytesread]

		uploadFile.Write(readBuffer)

		// Write to the hash generators
		md5hash.Write(readBuffer)
		sha1hash.Write(readBuffer)
		sha256hash.Write(readBuffer)
		sha3hash.Write(readBuffer)
	}

	hashes := provider.Hashes{
		MD5:    hex.EncodeToString(md5hash.Sum(nil)),
		SHA1:   hex.EncodeToString(sha1hash.Sum(nil)),
		SHA256: hex.EncodeToString(sha256hash.Sum(nil)),
		SHA3:   hex.EncodeToString(sha3hash.Sum(nil)),
	}

	uploadFile.Cleanup()

	oldPath := uploadProvider.GetInProgressFilePath(uploadFile.GetIdentifier())
	newPath := uploadProvider.GetFullFilePath(chunkHashPath(hashes.SHA256))
	uploadErr := uploadProvider.MoveFile(oldPath, newPath)

	if uploadErr != nil {
		return provider.Hashes{}, uploadErr
	}

	return hashes, nil
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