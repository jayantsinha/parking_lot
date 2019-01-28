package main

import "errors"

var (
	ParkingLotNotCreated            = errors.New("Parking lot not created")
	UnableToPark                    = errors.New("Unable to park. Please create a parking lot first.")
	ParkingLotFull                  = errors.New("Parking lot is full. Unable to park.")
	UnableToVacateEmptySlot         = errors.New("Slot already empty; unable to vacate.")
	UnableToVacateOnNonExistentSlot = errors.New("Slot not found; unable to vacate.")
)
