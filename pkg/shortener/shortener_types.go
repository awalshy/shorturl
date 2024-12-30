package shortener

import (
	"github.com/awalshy/shorturl/pkg/storage"
)

type Shortener struct {
	Storage *storage.Storage
}

var (
	shortener *Shortener
)

func NewShortener() *Shortener {
	if shortener == nil {
		shortener = &Shortener{
			Storage: storage.GetStorage(),
		}
	}
	return shortener
}