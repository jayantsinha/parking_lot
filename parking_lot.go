package main

import (
	"strings"
)

type Vehicle struct {
	RegnNumber string
	Color      string
}

type Slot struct {
	Num  int
	Vhcl *Vehicle
}

// ParkingLot stores the vehicle and slot info
type ParkingLot struct {
	Slots  []*Slot
	IsFull bool
}

var emptySlots *EmptySlots

// Init creates a parking lot with given number of slots
func (p *ParkingLot) Init(numSlots int) (int, error) {
	emptySlots = NewEmptySlots()
	if numSlots < 1 {
		return -1, ParkingLotNotCreated
	}
	p.Slots = make([]*Slot, numSlots, numSlots)
	for i := 0; i < numSlots; i++ {
		p.Slots[i] = &Slot{
			Num:  i,
			Vhcl: nil,
		}
		emptySlots.Add(i)
	}
	p.IsFull = false

	return len(p.Slots), nil
}

// Park parks a vehicle to the next available slot with the lowest serial number
func (p *ParkingLot) Park(regnNumber, color string) (int, error) {
	// Check if parking lot is initialized
	if p.Slots == nil {
		return -1, UnableToPark
	}

	// Check if parking lot is full
	if p.IsFull {
		return -1, ParkingLotFull
	}

	// Check for valid registration number and color string
	if !isValidRegistrationNumber(regnNumber) || !isValidColor(color) {
		return -1, InvalidInput
	}

	// Check for duplicate registration number
	if p.hasVehicle(regnNumber) {
		return -1, InvalidInput
	}

	slotToFill := emptySlots.GetMin()
	if slotToFill == -1 {
		p.IsFull = true
		return -1, ParkingLotFull
	}

	emptySlots.Remove(slotToFill)
	p.Slots[slotToFill] = &Slot{Vhcl: &Vehicle{
		RegnNumber: strings.TrimSpace(regnNumber),
		Color:      strings.TrimSpace(color)},
		Num: slotToFill,
	}
	return slotToFill + 1, nil
}

// Leave removes the vehicle from the specified slot number
func (p *ParkingLot) Leave(slotNum int) (int, error) {
	// Check if request is for an empty slot
	if emptySlots.Contains(slotNum - 1) {
		return -1, UnableToVacateEmptySlot
	}

	// Check if slot is not in the parking lot
	if len(p.Slots) < slotNum || slotNum < 1 {
		return -1, UnableToVacateOnNonExistentSlot
	}

	p.Slots[slotNum-1].Vhcl = nil
	emptySlots.Add(slotNum - 1)

	return slotNum, nil
}

// Status returns all the vehicles parked in the parking lot with their respective slt number
func (p *ParkingLot) Status() []*Slot {
	return p.Slots
}

// FindSlotNumbersByColor returns all the slot numbers which have a vehicle with the given color
func (p *ParkingLot) FindSlotNumbersByColor(color string) ([]int, error) {
	// For empty parking lot
	if p.Slots == nil {
		return []int{}, ParkingLotEmpty
	}

	slots := make([]int, 0)
	correctedColor := strings.TrimSpace(strings.ToLower(color))
	for _, v := range p.Slots {
		if v.Vhcl != nil && strings.ToLower(v.Vhcl.Color) == correctedColor {
			slots = append(slots, v.Num+1)
		}
	}

	if len(slots) == 0 {
		return []int{}, NotFound
	}

	return slots, nil
}

// FindRegistrationNumbersByColor returns all registration numbers of the vehicles with the given color
func (p *ParkingLot) FindRegistrationNumbersByColor(color string) ([]string, error) {
	// For empty parking lot
	if p.Slots == nil {
		return []string{}, ParkingLotEmpty
	}

	rnums := make([]string, 0)
	correctedColor := strings.TrimSpace(strings.ToLower(color))
	for _, v := range p.Slots {
		if v.Vhcl != nil && strings.ToLower(v.Vhcl.Color) == correctedColor {
			rnums = append(rnums, v.Vhcl.RegnNumber)
		}
	}

	if len(rnums) == 0 {
		return []string{}, NotFound
	}

	return rnums, nil
}

// FindSlotByRegistrationNumber returns the slot number where the vehicle with the given registration number is parked
func (p *ParkingLot) FindSlotByRegistrationNumber(regnNumber string) (int, error) {
	// For empty parking lot
	if p.Slots == nil {
		return -1, ParkingLotEmpty
	}
	slot := 0
	correctedRegn := strings.TrimSpace(strings.ToLower(regnNumber))
	for _, v := range p.Slots {
		if v.Vhcl != nil && strings.ToLower(v.Vhcl.RegnNumber) == correctedRegn {
			slot = v.Num + 1
			break
		}
	}

	// If registration number not found
	if slot == 0 {
		return -1, NotFound
	}

	return slot, nil
}

// isValidRegistrationNumber validates the registration number
// Registration number should have a minimum of 3 characters and any of these: (A-Z, a-z, 0-9 or -)
func isValidRegistrationNumber(regnNum string) bool {
	if len(regnNum) < 3 {
		return false
	}
	fn := func(r rune) bool {
		return !((r >= 'A' && r <= 'z') || (r >= '0' && r <= '9') || r == '\u002D')
	}
	if strings.IndexFunc(regnNum, fn) != -1 {
		return false
	}
	return true
}

// isValidColor checks color string for valid characters and min len of 3
// no special characters or numbers are allowed including space
func isValidColor(color string) bool {
	if len(color) < 3 {
		return false
	}
	fn := func(r rune) bool {
		return !(r > 'A' && r < 'z')
	}
	if strings.IndexFunc(color, fn) != -1 {
		return false
	}
	return true
}

// hasVehicle checks if the vehicle with registration number is already parked or not
func (p *ParkingLot) hasVehicle(regnNum string) bool {
	regnNumForSearch := strings.TrimSpace(strings.ToLower(regnNum))
	for _, v := range p.Slots {
		if v.Vhcl != nil && strings.TrimSpace(strings.ToLower(v.Vhcl.RegnNumber)) == regnNumForSearch {
			return true
		}
	}
	return false
}