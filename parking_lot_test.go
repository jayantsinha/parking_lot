package main

import "testing"

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