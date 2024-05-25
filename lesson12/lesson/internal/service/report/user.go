package report

import (
	"lesson/internal/entity"
	"log"
	"time"
)

type Report struct {
	DateFrom time.Time
	DateTo   time.Time
	Name     string
}

func (r Report) Do(f Formatter, d []entity.User) error {
	rep, err := f.Show(d)
	if err != nil {
		return err
	}
	log.Printf("создание отчета: %s", rep)
	return nil
}
