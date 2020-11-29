package filesystem_test

import (
	"bufio"
	"github.com/parmaanu/goutils/errorutils"
	"github.com/parmaanu/goutils/filesystem"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsFileSystem(t *testing.T) {
	var fs filesystem.FileSystem = filesystem.OsFileSystem{}
	fp, err := fs.Open("test.log")
	if errorutils.PrintOnErr("cannot open file", err) {
		os.Exit(1)
	}
	reader := bufio.NewReader(fp)

	data := []string{}
	for {
		buf, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(buf) == 0 {
			continue
		}

		data = append(data, string(buf))
	}
	expectedData := []string{"test,1", "test,2", "test,3"}
	assert.Equal(t, expectedData, data, "read data is different from file")
}

func TestMockFileSystem(t *testing.T) {
	mfs := filesystem.NewMockFileSystem()
	mfs.SetFileData("temp.log", []string{"test,1", "test,2", "test,3"})

	var fs filesystem.FileSystem = mfs
	fp, err := fs.Open("temp.log")
	if errorutils.PrintOnErr("cannot open file", err) {
		os.Exit(1)
	}
	reader := bufio.NewReader(fp)

	data := []string{}
	for {
		buf, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(buf) == 0 {
			continue
		}

		data = append(data, string(buf))
	}
	expectedData := []string{"test,1", "test,2", "test,3"}
	assert.Equal(t, expectedData, data, "read data is different from mock file")
}

func TestMockFileSystemStat(t *testing.T) {
	mfs := filesystem.NewMockFileSystem()
	{ // when file does not exist
		out, err := mfs.Stat("FileDoesNotExist.txt")
		assert.Nil(t, out)
		assert.EqualError(t, err, "FileDoesNotExist.txt file does not exist in MockFileSystem")
	}

	{ // when file exists in mock filesystem
		mfs.SetFileData("temp.log", []string{"test,1", "test,2", "test,3"})
		out, err := mfs.Stat("temp.log")
		assert.NotNil(t, out)
		assert.Nil(t, err)
	}
}
