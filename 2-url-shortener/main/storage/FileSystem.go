package storage

import (
	"sync"
	"os"
	"io/ioutil"
	"strconv"
	"path/filepath"
)

type FileSystem struct {
	Root string
	sync.RWMutex
}

func (fs *FileSystem) Init(root string) error {
	fs.Root = root

	return os.MkdirAll(root, 0744)
}

func (fs *FileSystem) Code() string {
	fs.Lock()

	files, _ := ioutil.ReadDir(fs.Root)
	fs.Unlock()

	return strconv.FormatUint(uint64(len(files)+1), 36)
}

func (fs *FileSystem) Save(url string) string {
	code := fs.Code()

	fs.Lock()
	ioutil.WriteFile(filepath.Join(fs.Root, code), []byte(url), 0744)
	fs.Unlock()

	return code
}

func (fs *FileSystem) Load(code string) (string, error) {
	fs.Lock()
	urlBytes, err := ioutil.ReadFile(filepath.Join(fs.Root, code))
	fs.Unlock()

	return string(urlBytes), err
}
