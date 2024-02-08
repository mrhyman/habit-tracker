package storage

import (
	"crypto/sha256"
	"errors"
	"etc/lib/e"
	"fmt"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(username string) (*Page, error)
	Remove(p *Page) error
	Exists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("there is no saved pages")

type Page struct {
	URL      string
	UserName string
	//Created time.Time
}

func (p *Page) Hash() (string, error) {
	h := sha256.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("Can't hash url", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("Can't hash user", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
