package main

type ParkingLotDesigner interface {
	Init() (int, error)
	Park(string, string) (int, error)
	Leave(int) (int, error)
	Status() []*struct{}
	FindSlotNumbersByColor(string) ([]int, error)
	FindRegistrationNumbersByColor(string) ([]string, error)
	FindSlotByRegistrationNumber(string) (int, error)
}