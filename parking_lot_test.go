package main

import (
	"testing"
)

func TestParkingLot_Init(t *testing.T) {
	type ParkingInitTest struct {
		n    int
		want int
	}
	tests := []ParkingInitTest{
		{0, 0},
		{1, 1},
		{3, 3},
		{10000, 10000},
	}
	for _, tt := range tests {
		t.Run("ParkingInitTest", func(t *testing.T) {
			p := new(ParkingLot)
			if got := p.Init(tt.n); got != tt.want {
				t.Errorf("ParkingLot.Init() = %v, want %v", got, tt.want)
			}
			if got := p.Init(tt.n); got != tt.want {
				t.Errorf("ParkingLot.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingLot_Park(t *testing.T) {
	type Vehicle struct {
		RegnNumber string
		Color      string
	}
	type VehicleParkingTest struct {
		Vehicle
		wantedSlotNum int
	}
	tests := []VehicleParkingTest{
		{Vehicle{RegnNumber: "KA-01MJ-4190", Color: "Red"}, 1},
		{Vehicle{RegnNumber: "BR1D-5621", Color: "Grey"}, 2},
		{Vehicle{RegnNumber: "BRL-106", Color: "Black"}, 3},
	}

	// Test parking without init
	for _, tt := range tests {
		t.Run("VehicleParkingTest", func(t *testing.T) {
			p := new(ParkingLot)
			_, err := p.Park(tt.RegnNumber, tt.Color)
			if (err == nil) || err != UnableToPark {
				t.Errorf("ParkingLot.Park() expected error = %v, but got %v", UnableToPark, err)
				return
			}
		})
	}

	// Test parking after init
	for _, tt := range tests {
		t.Run("VehicleParkingTest", func(t *testing.T) {
			p := new(ParkingLot)
			p.Init(10)
			got, err := p.Park(tt.RegnNumber, tt.Color)
			if (err != nil) {
				t.Errorf("ParkingLot.Park() expected error = %v, but got %v", nil, err)
				return
			}
			if got != tt.wantedSlotNum {
				t.Errorf("ParkingLot.Park() expected slot number to be %v, but got %v", tt.wantedSlotNum, got)
			}
		})
	}

	// Test parking on full parking lot
	for idx, tt := range tests {
		t.Run("VehicleParkingTest", func(t *testing.T) {
			p := new(ParkingLot)
			p.Init(1)
			got, err := p.Park(tt.RegnNumber, tt.Color)
			if idx == 2 {
				if (err != nil) {
					t.Errorf("ParkingLot.Park() expected error = %v, but got %v", ParkingLotFull, err)
					return
				}
			} else {
				if (err != nil) {
					t.Errorf("ParkingLot.Park() expected error = %v, but got %v", nil, err)
					return
				}
				if got != tt.wantedSlotNum {
					t.Errorf("ParkingLot.Park() expected slot number to be %v, but got %v", tt.wantedSlotNum, got)
				}
			}
		})
	}
}