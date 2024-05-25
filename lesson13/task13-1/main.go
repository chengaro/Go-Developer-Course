package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	type Contract struct {
		Number   int    `json:"number"`
		Landlord string `json:"landlord"`
		Tenat    string `json:"tenat"`
	}
	var c Contract

	j := strings.NewReader("{\"number\":1,\"landlord\":\"Остап Бендер\",\"tenat\":\"Шура Балаганов\"}")

	err := json.NewDecoder(j).Decode(&c)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", c)
}
