package main

import "errors"

type MotorCycle struct{
	colour string
	doorCount int
	driverName string
}

func NewMotorCycle() *MotorCycle{
	return &MotorCycle{
		colour : "red",
		doorCount: 0,
	}
}


func (m *MotorCycle)DoorCount() int {
	return m.doorCount
}


func (m *MotorCycle)GetColour() string {
	return m.colour
}


func (m *MotorCycle)AddDriver(name string) error {
	
	if name == ""{
		return errors.New("You did not include a name.")
	}

	m.driverName = name
	return nil
}


func (m *MotorCycle)GetDriver()  string {
	
	return m.driverName
}
