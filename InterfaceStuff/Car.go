package main

import "errors"

type Car struct{
	colour string
	doorCount int
	driverName string
}

func NewCar() *Car{
	return &Car{
		colour : "Blue",
		doorCount: 5,
	}
}


func (m *Car)DoorCount() int {
	return m.doorCount
}


func (m *Car)GetColour() string {
	return m.colour
}


func (m *Car)AddDriver(name string) error {
	
	if name == ""{
		return errors.New("You did not include a name.")
	}

	m.driverName = name
	return nil
}


func (m *Car)GetDriver() string {
	
	return m.driverName
}
