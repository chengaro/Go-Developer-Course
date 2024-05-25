package report

import (
	"lesson/internal/entity"
)

type Formatter interface {
	Show(d []entity.User) (string, error)
}
