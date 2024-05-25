package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Contract struct {
		Number   int    `json:"number"`
		Landlord string `json:"landlord"`
		Tenat    string `json:"tenat"`
	}
	c := Contract{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}

	j, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", j)

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(b.String())

}
