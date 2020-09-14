package provider

type Provider interface {
	GetID() string
	Initialize() error
	FileExists(path string) bool
	GetFileInfo(path string) (string, error)
	CreateUploadHandle() UploadFile
	MoveFile(oldPath string, newPath string) error
	Download(path string) ([]byte, error)
	Delete(path string) error
	GetFullFilePath(path string) string
	GetInProgressFilePath(path string) string
}

type AccessType int

const (
	Streamable = iota
	DirectLink
)

type UploadFile interface {
	Initialize()
	Write(bytes []byte) error
	GetIdentifier() string
	Cleanup()
}

type Hashes struct {
	MD5    string
	SHA1   string
	SHA256 string
	SHA3   string
}
