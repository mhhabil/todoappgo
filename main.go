package main

import (
	"todoappgo/models"
	"fmt"
)

type Product struct {
	Name  string
	Price int
	Desc  string
}

type Merchant struct {
	Name     string
	Products []Product
}

func main() {
	fmt.Println("Hello, World!")

	var username string
	var password string
	day := "siang" // siang | malam
	if day == "siang" {
		username = "fikri"
		password = "fikri123"
	} else {
		username = "agus"
		password = "agus123"
	}

	product1 := Product{
		Name:  "Product 1",
		Price: 1000,
		Desc:  "Product 1 description",
	}
	fmt.Printf("Product 1 : %+v\n", product1)

	product2 := Product{
		Name:  "Product 2",
		Price: 2000,
		Desc:  "Product 2 description",
	}
	fmt.Printf("Product 2 : %+v\n", product2)

	merchant := Merchant{
		Name:     "Merchant 1",
		Products: []Product{product1, product2},
	}
	fmt.Printf("Merchant : %+v\n", merchant)

	User := model.User{}
	status, err := User.CreateUser(username, password)
	if err != nil {
		fmt.Println(err) // ada error
	} else {
		fmt.Println(status) // success
	}
}