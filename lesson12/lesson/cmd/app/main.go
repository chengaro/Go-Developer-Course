package main

import (
	"lesson/internal/entity"
	"lesson/internal/service/format"
	"lesson/internal/service/report"
)

func main() {
	f := format.Xml{}
	r := report.Report{}
	data := make([]entity.User, 0, 1)
	r.Do(f, data)
}
