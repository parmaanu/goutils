package fileutils

import (
	"bufio"
	"errors"
	filesystem "github.com/parmaanu/goutils/filesystem"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

// GetCurrentUsername returns the current username
func GetCurrentUsername() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return usr.Username
}

// FileExist returns true if the file exists
func FileExist(fname string) bool {
	if fileinfo, err := filesystem.Stat(fname); err != nil || fileinfo.Mode().IsRegular() == false {
		return false
	}
	return true
}

// DirExist returns true if the directory exists
func DirExist(dir string) bool {
	if fileinfo, err := os.Stat(dir); err != nil || fileinfo.Mode().IsDir() == false {
		return false
	}
	return true
}

// GetBaseDir returns the base directory
func GetBaseDir(fpath string) string {
	return filepath.Base(filepath.Dir(fpath))
}

// ExpandIfExists expands the filename
func ExpandIfExists(path string) (string, error) {
	fullpath, err := tilde.Expand(path)
	if err != nil {
		return "", err
	}
	fullpath, err = filepath.Abs(fullpath)
	if err != nil {
		return "", err
	}
	if FileExist(fullpath) == false && DirExist(fullpath) == false {
		return "", errors.New("file or dir not found: " + fullpath)
	}
	return fullpath, nil
}

// WriteBufferToFile writes the given buf to a file.
// The file is created or truncated (if it already exists). It uses os.Create internally
func WriteBufferToFile(fname string, buf []byte) error {
	fo, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = fo.Write(buf)
	if err != nil {
		return err
	}
	fo.Close()
	return nil
}

// GetPidfile return ~/.local/<appname>-<username>.pid file
func GetPidfile(appname string) (string, error) {
	localpath, err := tilde.Expand("~/.local")
	if err != nil {
		return "", err
	}
	if DirExist(localpath) == false {
		err := os.Mkdir(localpath, 0755)
		if err != nil {
			return "", err
		}
	}
	return localpath + "/" + appname + "-" + GetCurrentUsername() + ".pid", nil
}

// ReadFullFileAsBytes reads the file and retruns the data as byte slice
func ReadFullFileAsBytes(fname string) []byte {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return []byte{}
	}
	return buf
}

// Readln returns a single line (without the ending \n) from the input buffered reader.
// An error is returned if there is an error with the buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// ReadStdin reads a single char from stdin
func ReadStdin() rune {
	ch, _, _ := bufio.NewReader(os.Stdin).ReadRune()
	return ch
}
