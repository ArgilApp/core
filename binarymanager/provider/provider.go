package provider

type Provider interface {
	GetID() string
	Initialize() error
	FileExists(path string) bool
	GetFileInfo(path string) (FileInfo, error)
	CreateUploadHandle() UploadFile
	MoveFile(oldPath string, newPath string) error
	CreateDownloadHandle(path string) DownloadFile
	GetDirectDownloadLink(path string) (string, error)
	Delete(path string) error
	GetFullFilePath(path string) string
	GetInProgressFilePath(path string) string
	SupportedDownloadAccessTypes() []AccessType
}

// AccessType is not implemented yet
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

type DownloadFile interface {
	Initialize(path string)
	Read(buffer []byte) (int, error)
	Cleanup()
}

type Hashes struct {
	MD5    string
	SHA1   string
	SHA256 string
	SHA3   string
}

type FileInfo struct {
}
