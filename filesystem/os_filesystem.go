package filesystem

import "os"

// OsFileSystem implements fileSystem using the local disk.
type OsFileSystem struct{}

// Open file
func (OsFileSystem) Open(name string) (File, error) { return os.Open(name) }

// Stat give fileinfo of the file
func (OsFileSystem) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
