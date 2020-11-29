package filesystem

import (
	"os"
	"strings"
)

type mockFile struct {
	*strings.Reader
	name string
}

func (f mockFile) Name() string { return f.name }

func (f mockFile) Close() error { return nil }

type mockFileInfo struct {
	os.FileInfo
}

func (m mockFileInfo) Mode() os.FileMode {
	return 0
}
