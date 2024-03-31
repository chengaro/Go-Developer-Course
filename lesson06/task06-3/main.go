package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}

func (c contract) String() string {
	return fmt.Sprintf("Договор № %s от %s", c.Number, c.Date)
}

func main() {
	c := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}
	fmt.Printf("%s", c)
}
