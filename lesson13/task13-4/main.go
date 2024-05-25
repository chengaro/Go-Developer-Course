package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

func main() {
	type contract struct {
		Number   int    `xml:"number"`
		Landlord string `xml:"landlord"`
		Tenat    string `xml:"tenat"`
	}

	type contracts struct {
		List []contract `xml:"contract"`
	}

	c := contracts{}
	c.List = append(c.List, contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	})
	c.List = append(c.List, contract{
		Number:   2,
		Landlord: "Шура Балаганов",
		Tenat:    "Остап Бендер",
	})

	x, err := xml.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", string(x))
}
