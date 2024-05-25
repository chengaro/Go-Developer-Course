package format

import (
	"lesson/internal/entity"
	"log"
)

type Xml struct {
	version string
	root    string
}

func (x Xml) Show(d []entity.User) (string, error) {
	log.Println("Данные в формате xml")
	return "Данные в формате xml", nil
}
