package fileutils_test

import (
	"github.com/parmaanu/goutils/fileutils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

func TestFileExist(t *testing.T) {
	assert.False(t, fileutils.FileExist("testfile"), "'testfile' file should not exist")
	assert.True(t, fileutils.FileExist("fileutils_test.go"), "'fileutils_test.go' should exist")
}

func TestDirExist(t *testing.T) {
	assert.False(t, fileutils.DirExist("testdir"), "'testdir' should not exist")
	assert.False(t, fileutils.DirExist("fileutils_test.go"), "fileutils_test.go is not a directory")
	assert.True(t, fileutils.DirExist("."), "current directory should exist")
	assert.True(t, fileutils.DirExist(".."), "one directory up should exist")
	assert.True(t, fileutils.DirExist("../fileutils"), "../fileutils directory should exist")
}

func TestGetBaseDir(t *testing.T) {
	assert.Equal(t, "basedir", fileutils.GetBaseDir("~/long/directory/name/basedir/file"), "base directory is not correct")
	assert.Equal(t, "basedir", fileutils.GetBaseDir("~/long/directory/name/basedir/file.go"), "base directory is not correct")
}

func TestExpandIfExists(t *testing.T) {
	{
		homeDir, err := tilde.Expand("~")
		assert.NoError(t, err, "error while expanding home directory")

		fname, err := fileutils.ExpandIfExists("~/")
		assert.NoError(t, err, "error while expanding home directory")
		assert.Equal(t, homeDir, fname, "expanded filename is not correct")
	}
	{
		currDir, err := os.Getwd()
		assert.NoError(t, err, "error while getting current directory")

		fname, err := fileutils.ExpandIfExists("../fileutils")
		assert.NoError(t, err, "error while expanding current directory")
		assert.Equal(t, currDir, fname, "expanded filename is not correct")
	}
}
