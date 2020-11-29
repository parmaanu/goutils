package filesystem

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

func expandFilename(fname string) string {
	fname, err := tilde.Expand(fname)
	fname, err1 := filepath.Abs(fname)
	if err != nil || err1 != nil {
		fmt.Println("problem with filename, cannot get absolute filepath")
		os.Exit(1)
	}
	return fname
}

// MockFileSystem implements the mock filesystem which reads from []string{}
type MockFileSystem struct {
	mockFileData map[string]string
}

// SetFileData sets the data for correspoding file and data in that file
func (m *MockFileSystem) SetFileData(fname string, fdata []string) {
	data := strings.Join(fdata, "\n")
	m.mockFileData[fname] = data

	absFileName := expandFilename(fname)
	if absFileName != fname {
		m.mockFileData[absFileName] = data
	}
}

// Open file
func (m *MockFileSystem) Open(fname string) (File, error) {
	filedata, exists := m.mockFileData[fname]
	if !exists {
		absFileName := expandFilename(fname)
		filedata, exists = m.mockFileData[absFileName]
		if !exists {
			return nil, errors.New("file not found in MockFileSystem")
		}
	}
	return mockFile{name: fname, Reader: strings.NewReader(filedata)}, nil
}

// Stat give fileinfo of the mock file
func (m *MockFileSystem) Stat(fname string) (os.FileInfo, error) {
	_, exists := m.mockFileData[fname]
	if !exists {
		return nil, errors.New(fname + " file does not exist in MockFileSystem")
	}
	return mockFileInfo{}, nil
}

// NewMockFileSystem returns the objects of type MockFileSystem
func NewMockFileSystem() *MockFileSystem {
	var mfs MockFileSystem
	mfs.mockFileData = make(map[string]string)
	return &mfs
}
