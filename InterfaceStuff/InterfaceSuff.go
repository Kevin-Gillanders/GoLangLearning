package main

import (
	"fmt"
	"reflect"
)

type IVehicle interface{
	DoorCount() int
	GetColour() string
	GetDriver() string
	AddDriver(name string) error
}


func main() {
	vehicleSlice := [] IVehicle { NewMotorCycle(), NewCar()}
	for _, vehicle := range vehicleSlice{
		fmt.Printf("The type of this object is %v\n", reflect.TypeOf(vehicle))
		fmt.Printf("DoorCount : %v\n" , vehicle.DoorCount())
		fmt.Printf("The colour : %v\n", vehicle.GetColour())
		fmt.Printf("The Driver : %v\n", vehicle.GetDriver())
		vehicle.AddDriver("Marcus Parks")
		fmt.Printf("The Driver : %v\n", vehicle.GetDriver())
	}
}

	//{ NewMotorCycle(), NewCar()}

// 	for vehicle in range vehicleSlice {

// 	}

// 	// motorCycle := NewMotorCycle()


// }

