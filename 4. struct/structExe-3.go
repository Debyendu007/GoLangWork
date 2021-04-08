package main

import "fmt"

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	vehicle   Vehicle
	fourWheel bool
}

type Sedan struct {
	vehicle Vehicle
	luxury  bool
}

func main() {
	truck1 := Truck{
		vehicle: Vehicle{
			doors: 2,
			color: "Black",
		},
		fourWheel: false,
	}

	sedan1 := Sedan{
		vehicle: Vehicle{
			doors: 5,
			color: "White",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)
	fmt.Println(truck1.vehicle)
	fmt.Println(truck1.fourWheel)
	fmt.Println(sedan1)
	fmt.Println(sedan1.vehicle)
	fmt.Println(sedan1.luxury)
}
