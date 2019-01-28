package main

import (
	"fmt"
	"strconv"
	"strings"
)

var p *ParkingLot

var (
	Version              = "1.4.1"
	InvalidNumOfArgument = "Invalid number of arguments"
	InvalidValuePassed   = "Invalid value passed"
)

var isParkingLotCreated bool

func CallRespFunc(args []string, f func(args []string) bool) bool {
	return f(args)
}

// CreateParkingLot is the handler for create_parking_lot command
func CreateParkingLot(args []string) (ret bool) {
	if len(args) != 2 {
		fmt.Println("Invalid syntax. Use 'create_parking_lot <num_slots>'")
		return
	}
	p = new(ParkingLot)
	numSlots, err := strToInt(args[1])
	if err != nil {
		fmt.Println(InvalidValuePassed)
		return
	}
	n, err := p.Init(numSlots)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret = true
	fmt.Println("Created a parking lot with", n, "slots")
	isParkingLotCreated = true
	return
}

// ParkVehicle is the handler for park command
func ParkVehicle(args []string) (ret bool) {
	if !isParkingLotCreated {
		fmt.Println("Create a parking lot first using 'create_parking_lot <num_slots>'")
		return
	}
	// check for valid number of arguments
	if len(args) != 3 {
		fmt.Println(InvalidNumOfArgument)
		return
	}
	n, err := p.Park(args[1], args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	ret = true
	fmt.Println("Allocated slot number:", n)
	return
}

// LeaveSlot is the handler for leave command
func LeaveSlot(args []string) (ret bool) {
	// check for valid number of arguments
	if len(args) != 2 {
		fmt.Println(InvalidNumOfArgument)
		return
	}
	slot, err := strToInt(args[1])
	if err != nil {
		fmt.Println(InvalidValuePassed)
		return
	}
	n, err := p.Leave(slot)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret = true
	fmt.Println("Slot number", n, "is free")
	return
}

// GetStatus is the handler for status command
func GetStatus(args []string) (ret bool) {
	// check for valid number of arguments
	if len(args) != 1 {
		fmt.Println(InvalidNumOfArgument)
		return
	}

	heading := []string{"Slot No.", "Registration No", "Color"}
	slots := p.Status()
	if len(slots) == 0 {
		fmt.Print(heading[0], "\t", heading[1], "\t", heading[2], "\n")
		return
	}

	fmt.Print(heading[0], "\t", heading[1], "\t", heading[2], "\n")
	for _, s := range slots {
		if s.Vhcl != nil {
			fmt.Print(s.Num+1, "\t")
			fmt.Print(s.Vhcl.RegnNumber, "\t")
			fmt.Println(s.Vhcl.Color)
		}
	}
	ret = true
	return
}

// GetRegNumsByColor is the handler for registration_numbers_for_cars_with_colour command
func GetRegNumsByColor(args []string) (ret bool) {
	// check for valid number of arguments
	if len(args) != 2 {
		fmt.Println(InvalidNumOfArgument)
		return
	}

	rnums, err := p.FindRegistrationNumbersByColor(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Join(rnums, ", "))
	ret = true
	return
}

// GetSlotNumByRegNum is the handler for slot_number_for_registration_number command
func GetSlotNumByRegNum(args []string) (ret bool) {
	// check for valid number of arguments
	if len(args) != 2 {
		fmt.Println(InvalidNumOfArgument)
		return
	}

	slot, err := p.FindSlotByRegistrationNumber(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(slot)
	ret = true
	return
}

// GetSlotsByColor is the handler for slot_numbers_for_cars_with_colour command
func GetSlotsByColor(args []string) (ret bool) {
	// check for valid number of arguments
	if len(args) != 2 {
		fmt.Println(InvalidNumOfArgument)
		return
	}

	slots, err := p.FindSlotNumbersByColor(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(slots), " ", ", ", -1), "[]"))
	ret = true
	return
}

// EndSession is the handler for exit command
func EndSession(args []string) (ret bool) {
	fmt.Println(8)
	return
}

func ShowVersion(args []string) (ret bool) {
	fmt.Println("Parking Lot version", Version)
	return
}

func strToInt(val string) (int, error) {
	return strconv.Atoi(val)
}
