package main

import (
	c "belajar-package/currency"
	"fmt"
)

func main() {
	var testChann = make(chan string)

	fmt.Printf("testChann: %v\n", testChann)

	var countryName = c.Country{"Indonesia"}
	var currencyName = c.Currency{countryName, "IDR"}
	var money = c.Money{currencyName, 500000}
	var price = c.Price{money, "Lenovo"}

	fmt.Println(money.Amount)
	fmt.Println(price)
	fmt.Println("test")
}
