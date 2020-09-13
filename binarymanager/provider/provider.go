package provider

type Provider interface {
	GetID() string
	Initialize() error
	FileExists(path string) bool
	GetFileInfo(path string) (error, string)
	Upload(stream string) (error, string) // string should be stream/bytes
	Download(path string) (error, string) // string should be stream/bytes
	Delete(path string) error
}
