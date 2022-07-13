package main

import (
	"fmt"
	"log"
)

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	var test *int32
	log.Println(test)
	log.Println(&data.height)
	test = &data.age
	log.Println(*test)

	fmt.Printf("%b \n", 1025)      //biner
	fmt.Printf("%c \n", 0141)      //unicode
	fmt.Printf("%d \n", 11002)     //decimal
	fmt.Printf("%o \n", data.age)  //octal
	fmt.Printf("%x \n", data.name) //hexa
	fmt.Printf("%e \n", data.height)
	fmt.Printf("%.10f \n", data.height)
	fmt.Printf("%p \n", &data.name)
	fmt.Printf("%T \n", data.height)
	fmt.Printf("%v \n", data)
	fmt.Printf("%+v \n", data)
	fmt.Printf("%#v \n", data)
}
