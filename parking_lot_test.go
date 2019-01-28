package main

import (
	"reflect"
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
				t.Errorf("ParkingLot.Park() expected error = %v, but got: %v", UnableToPark, err)
				return
			}
		})
	}

	// Test parking after init
	p := new(ParkingLot)
	p.Init(10)
	for _, tt := range tests {
		t.Run("VehicleParkingTest", func(t *testing.T) {
			got, err := p.Park(tt.RegnNumber, tt.Color)
			if err != nil {
				t.Errorf("ParkingLot.Park() expected error = %v, but got: %v", nil, err)
				return
			}
			if got != tt.wantedSlotNum {
				t.Errorf("ParkingLot.Park() expected slot number to be %v, but got %v", tt.wantedSlotNum, got)
			}
		})
	}

	// Test parking on full parking lot
	p = new(ParkingLot)
	p.Init(1)
	for idx, tt := range tests {
		t.Run("VehicleParkingTest", func(t *testing.T) {
			got, err := p.Park(tt.RegnNumber, tt.Color)
			if idx >= 1 {
				if err == nil {
					t.Errorf("ParkingLot.Park() expected error = %v, but got: %v", ParkingLotFull, err)
					return
				}
			} else {
				if err != nil {
					t.Errorf("ParkingLot.Park() expected error = %v, but got: %v", nil, err)
					return
				}
				if got != tt.wantedSlotNum {
					t.Errorf("ParkingLot.Park() expected slot number to be %v, but got: %v", tt.wantedSlotNum, got)
				}
			}
		})
	}
}

func TestParkingLot_Leave(t *testing.T) {
	slotTests := []struct {
		slotNum     int
		vacatedSlot int
	}{
		{1, 1},
		{2, 2},
		{2, -1},
		{3, -1},
		{-3, -1},
	}
	p = new(ParkingLot)
	p.Init(2)
	_, _ = p.Park("2122", "Red")
	_, _ = p.Park("1111", "Purple")
	for idx, tt := range slotTests {
		t.Run("ParkingLotLeaveTest", func(t *testing.T) {
			got, err := p.Leave(tt.slotNum)
			if idx == 2 {
				if err == nil {
					t.Errorf("ParkingLot.Leave() expected error = %v but got %v", UnableToVacateEmptySlot, err)
					return
				}
			} else if idx >= 3 {
				if err == nil {
					t.Errorf("ParkingLot.Leave() expected error = %v but got %v , got ret %d", UnableToVacateOnNonExistentSlot, err, got)
					return
				}
			} else {
				if err != nil {
					t.Errorf("ParkingLot.Leave() error = %v for %d", err, tt.slotNum)
					return
				}
				if got != tt.vacatedSlot {
					t.Errorf("ParkingLot.Leave() got %v, want %v", got, tt.vacatedSlot)
				}
			}
		})
	}
}

func TestParkingLot_Status(t *testing.T) {
	p = new(ParkingLot)
	// Check on empty parking lot
	got := len(p.Status())
	if got != 0 {
		t.Errorf("ParkingLot.Status(): got length of Status() = %d; expected %d", got, 0)
	}
	p.Init(2)
	_, _ = p.Park("2122", "Red")
	_, _ = p.Park("1111", "Purple")
	got = len(p.Status())
	if got != 2 {
		t.Errorf("ParkingLot.Status(): got length of Status() = %d; expected %d", got, 2)
	}
}

func TestParkingLot_FindSlotNumbersByColor(t *testing.T) {
	p = new(ParkingLot)
	p.Init(6)
	_, _ = p.Park("1111", "Red")
	_, _ = p.Park("2222", "Purple")
	_, _ = p.Park("3333", "White")
	_, _ = p.Park("4444", "White")
	_, _ = p.Park("5555", "Blue")
	_, _ = p.Park("6666", "White")
	tests := []struct {
		color   string
		want    []int
		wantErr bool
	}{
		{"Red", []int{1}, false},
		{"Purple", []int{2}, false},
		{"White", []int{3, 4, 6}, false},
		{"Blue ", []int{5}, false},
		{"BLUE", []int{5}, false},
		{"Yellow", []int{}, true},
	}
	for idx, tt := range tests {
		t.Run("ParkingLotSlotNumberByColorTest", func(t *testing.T) {
			got, err := p.FindSlotNumbersByColor(tt.color)
			if idx >= 5 {
				if err == nil {
					t.Errorf("ParkingLot.FindSlotNumbersByColor() got error = %v, but wanted error = %v", err, NotFound)
					return
				}
			} else {
				if (err != nil) != tt.wantErr {
					t.Errorf("ParkingLot.FindSlotNumbersByColor() got error = %v, but wanted error = %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ParkingLot.FindSlotNumbersByColor() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestParkingLot_FindRegistrationNumbersByColor(t *testing.T) {
	p = new(ParkingLot)
	p.Init(6)
	_, _ = p.Park("1111", "Red")
	_, _ = p.Park("2222", "Purple")
	_, _ = p.Park("3333", "White")
	_, _ = p.Park("4444", "White")
	_, _ = p.Park("5555", "Blue")
	_, _ = p.Park("6666", "White")
	tests := []struct {
		color   string
		want    []string
		wantErr bool
	}{
		{"Red", []string{"1111"}, false},
		{"Purple", []string{"2222"}, false},
		{"White", []string{"3333", "4444", "6666"}, false},
		{"Blue ", []string{"5555"}, false},
		{"BLUE", []string{"5555"}, false},
		{"Yellow", []string{}, true},
	}
	for idx, tt := range tests {
		t.Run("ParkingLotFindRegistrationNumbersByColorTest", func(t *testing.T) {
			got, err := p.FindRegistrationNumbersByColor(tt.color)
			if idx >= 5 {
				if err == nil {
					t.Errorf("ParkingLot.FindRegistrationNumbersByColor() got error = %v, but wanted error = %v", err, NotFound)
					return
				}
			} else {
				if (err != nil) != tt.wantErr {
					t.Errorf("ParkingLot.FindRegistrationNumbersByColor() got error = %v, but wanted error = %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ParkingLot.FindRegistrationNumbersByColor() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestParkingLot_FindSlotByRegistrationNumber(t *testing.T) {
	p = new(ParkingLot)
	p.Init(6)
	_, _ = p.Park("KA-01MJ-4190", "Red")
	_, _ = p.Park("BRL-106", "Black")
	_, _ = p.Park("DL-3CA-7766", "White")
	_, _ = p.Park("BR-1d-5621 ", "Grey")
	tests := []struct {
		regnNumber string
		wantSlot   int
		wantErr    bool
	}{
		{"KA-01MJ-4190", 1, false},
		{" Ka-01mj-4190 ", 1, false},
		{"brl-106", 2, false},
		{"DL3CA7766", -1, true},
		{"", -1, true},
	}
	for _, tt := range tests {
		t.Run("ParkingLotFindSlotByRegistrationNumberTest", func(t *testing.T) {
			got, err := p.FindSlotByRegistrationNumber(tt.regnNumber)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParkingLot.FindSlotByRegistrationNumber(): got err: %v but want: %v", NotFound, err)
					return
				}
			} else {
				if err != nil {
					t.Errorf("ParkingLot.FindSlotByRegistrationNumber() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.wantSlot {
					t.Errorf("ParkingLot.FindSlotByRegistrationNumber() got slot = %v, want %v", got, tt.wantSlot)
				}
			}
		})
	}
}
