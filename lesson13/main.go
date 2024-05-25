package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Id    int `json:"id"`
	Extra *E  `json:"extra,omitempty"`
}

type E struct {
	Message string `json:"message,omitempty"`
}

func (t T) String() string {
	if t.Extra == nil {
		return fmt.Sprintf("Id: %d", t.Id)
	}
	return fmt.Sprintf(`{"id": %d, "extra": {"message": "%s"}}"`, t.Id, t.Extra.Message)
}

func main() {
	c := T{
		Id: 1,
	}
	fmt.Printf("структура %+v", c)
	fmt.Println()
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", string(res))
	fmt.Println()

	var a T
	json.Unmarshal(res, &a)
	fmt.Printf("структура %+v\n", a)
	a.Extra = &E{Message: "msg"}
	fmt.Printf("%s", a)
}
