package crud

import (
	"github.com/uncle-gua/mgm"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func newBook(name string, pages int) *book {
	return &book{
		Name:  name,
		Pages: pages,
	}
}
