package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err)
	err3 := fmt.Errorf("ошибка3:%w", err2)
	log.Fatalln(err3)
}
