package main

import "testing"

func TestCreateParkingLot(t *testing.T) {
	tests := []struct {
		command    string
		wantRet bool
	}{
		{"create_parking_lot 0", false},
		{"create_parking_lot -1", false},
		{"create_parking_lot 1", true},
		{"create_parking_lot 10", false},
		{"create_parking_lot 10000", false},
	}
	for _, tt := range tests {
		t.Run("CreateParkingLotHandlerTest", func(t *testing.T) {
			args := separateFlags(tt.command)
			if gotRet := CreateParkingLot(args); gotRet != tt.wantRet {
				t.Errorf("CreateParkingLot() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestParkVehicle(t *testing.T) {
	isParkingLotCreated = false
	args := separateFlags("create_parking_lot 6")
	CreateParkingLot(args)
	tests := []struct {
		command    string
		wantRet bool
	}{
		{"park KA-01-HH-1234 White", true},
		{"park KA-01-HH-9999 White", true},
		{"park KA-01-HH-9999 White", false},
		{"park KA-01-BB-0001 Black", true},
		{"park KA-01-HH-7777 Red", true},
		{"park KA-01-HH-2701 Blue", true},
		{"park KA-01-HH-3141 Black", true},
		{"park KA-01-HH-999 White", false},
		{"park", false},
		{"park white", false},
		{"park b^& Red", false},
	}
	for _, tt := range tests {
		t.Run("ParkVehicleHandlerTest", func(t *testing.T) {
			args = separateFlags(tt.command)
			if gotRet := ParkVehicle(args); gotRet != tt.wantRet {
				t.Errorf("ParkVehicle() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestLeaveSlot(t *testing.T) {
	isParkingLotCreated = false
	args := separateFlags("create_parking_lot 1")
	CreateParkingLot(args)
	args = separateFlags("park KA-01-HH-1234 White")
	ParkVehicle(args)
	args = separateFlags("leave 1")
	ret := LeaveSlot(args)
	if !ret {
		t.Errorf("LeaveSlot() got: %v, want: %v", ret, true)
	}
	args = separateFlags("leave 4")
	ret = LeaveSlot(args)
	if ret {
		t.Errorf("LeaveSlot() got: %v, want: %v", ret, false)
	}
}

func TestGetStatus(t *testing.T) {
	args := separateFlags("create_parking_lot 1")
	CreateParkingLot(args)
	args = separateFlags("park KA-01-HH-1234 White")
	ParkVehicle(args)
	ret := GetStatus([]string{"status"})
	if !ret {
		t.Errorf("GetStatus() got: %v, want: %v", ret, true)
	}
}