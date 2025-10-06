package main

import "fmt"

func main() {

	var honda = car {
		name: "Honda",
		model: "2019",
	}

	fmt.Println(honda.getCarModel())
	fmt.Println(honda.getCarName())

	var fiat = car {
		name: "Fiat",
		model: "2020",
	}

	fmt.Println(fiat.getCarModel())
	fmt.Println(fiat.getCarName())
}


type Car interface {
	getCarName() string
	getCarModel() string
}

type car struct {
	name string
	model string
}


func (c car) getCarName() string {
	return fmt.Sprintf("this car name is %v", c.name)
}

func (c car) getCarModel() string {
	return fmt.Sprintf("this car model is %v", c.model)
}