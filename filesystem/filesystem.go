package filesystem

import (
	"io"
	"os"
)

// FileSystem implements the file system interface
type FileSystem interface {
	Open(name string) (File, error)
	Stat(name string) (os.FileInfo, error)
}

// File interface
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
}

// Represents the current filesystem being used in the program
var fs FileSystem = OsFileSystem{}

// SetFileSystem sets the file system to be used (MockFileSystem is used during testing)
func SetFileSystem(newFS FileSystem) {
	fs = newFS
}

// GetFileSystem returns the pointer to the current filesystem
func GetFileSystem() FileSystem {
	return fs
}

// Open returns the results of Open from the current filesystem in place
func Open(filename string) (File, error) {
	return fs.Open(filename)
}

// Stat returns os level information about the file
func Stat(filename string) (os.FileInfo, error) {
	return fs.Stat(filename)
}
