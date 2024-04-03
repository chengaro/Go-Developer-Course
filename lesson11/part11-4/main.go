package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	code    int
	message string
}

func (e myFirstError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

func main() {
	err := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err)
	err3 := fmt.Errorf("ошибка3:%w", err2)

	if !errors.As(err3, new(myFirstError)) {
		fmt.Println("Не было myFirstError")
	} else {
		fmt.Println("О, наткнулись на myFirstError")
	}
}
