package main

type ParkingLotDesigner interface {
	Init(int) (int, error)
	Park(string, string) (int, error)
	Leave(int) (int, error)
	Status() []*Slot
	FindSlotNumbersByColor(string) ([]int, error)
	FindRegistrationNumbersByColor(string) ([]string, error)
	FindSlotByRegistrationNumber(string) (int, error)
}