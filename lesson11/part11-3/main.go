package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err)
	err3 := fmt.Errorf("ошибка3:%w", err2)
	fmt.Printf("Final error is ошибка1: %v", errors.Is(err3, err))
}
