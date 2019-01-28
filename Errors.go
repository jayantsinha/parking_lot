package main

import "errors"

var (
	ParkingLotNotCreated            = errors.New("Parking lot not created")
	UnableToPark                    = errors.New("Unable to park. Please create a parking lot first.")
	ParkingLotFull                  = errors.New("Sorry, parking lot is full")
	UnableToVacateEmptySlot         = errors.New("Slot already empty; unable to vacate.")
	UnableToVacateOnNonExistentSlot = errors.New("Slot not found; unable to vacate.")
	NotFound                        = errors.New("Not Found")
	InvalidInput                    = errors.New("Invalid Input")
	ParkingLotEmpty                 = errors.New("Parking lot empty!")
)
