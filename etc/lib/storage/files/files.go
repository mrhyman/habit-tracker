package files

import (
	"encoding/gob"
	"errors"
	"etc/lib/e"
	"etc/lib/storage"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type FileStorage struct {
	basePath string
}

const defaultPermissions = 0774

func New(basePath string) FileStorage {
	return FileStorage{basePath}
}

func (fs FileStorage) Save(page *storage.Page) (err error) {
	defer func() { err = e.Wrap("Can't save file", err) }()

	fPath := filepath.Join(fs.basePath, page.UserName)
	if err := os.MkdirAll(fPath, defaultPermissions); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (fs FileStorage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.Wrap("Can't pick random page", err) }()

	path := filepath.Join(fs.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, storage.ErrNoSavedPages
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(len(files))

	file := files[n]
	return fs.decodePage(filepath.Join(path, file.Name()))
}

func (fs FileStorage) Remove(p *storage.Page) error {
	fName, err := fileName(p)
	if err != nil {
		return e.Wrap("can't remove page", err)
	}

	path := filepath.Join(fs.basePath, p.UserName, fName)
	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file: %s", path)
		return e.Wrap(msg, err)
	}

	return nil
}

func (fs FileStorage) Exists(p *storage.Page) (bool, error) {
	fName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("can't find file", err)
	}

	path := filepath.Join(fs.basePath, p.UserName, fName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file exists: %s", path)
		return false, e.Wrap(msg, err)
	}

	return true, nil
}

func (fs FileStorage) decodePage(fPath string) (*storage.Page, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, e.Wrap("can't open on decode page", err)
	}

	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode gob into page", err)
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
