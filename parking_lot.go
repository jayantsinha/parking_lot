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

type ParkingLot struct {
	Slots  []*Slot
	IsFull bool
}

var emptySlots *EmptySlots

func (p *ParkingLot) Init(numSlots int) int {
	emptySlots = NewEmptySlots()
	p.Slots = make([]*Slot, numSlots, numSlots)
	for i := 0; i < numSlots; i++ {
		p.Slots[i] = &Slot{
			Num:  i,
			Vhcl: nil,
		}
		emptySlots.Add(i)
	}
	p.IsFull = false

	return len(p.Slots)
}

func (p *ParkingLot) Park(regnNumber, color string) (int, error) {
	if p.Slots == nil {
		return -1, UnableToPark
	}

	if p.IsFull {
		return -1, ParkingLotFull
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

func (p *ParkingLot) Status() []*Slot {
	return p.Slots
}

func (p *ParkingLot) FindSlotNumbersByColor(color string) ([]int, error) {
	slots := make([]int, 0)
	correctedColor := strings.TrimSpace(strings.ToLower(color))
	for _, v := range p.Slots {
		if strings.ToLower(v.Vhcl.Color) == correctedColor {
			slots = append(slots, v.Num+1)
		}
	}

	if len(slots) == 0 {
		return []int{}, NotFound
	}

	return slots, nil
}

func (p *ParkingLot) FindRegistrationNumbersByColor(color string) ([]string, error) {
	rnums := make([]string, 0)
	correctedColor := strings.TrimSpace(strings.ToLower(color))
	for _, v := range p.Slots {
		if strings.ToLower(v.Vhcl.Color) == correctedColor {
			rnums = append(rnums, v.Vhcl.RegnNumber)
		}
	}

	if len(rnums) == 0 {
		return []string{}, NotFound
	}

	return rnums, nil
}

//func (p *ParkingLot) FindSlotByRegistrationNumber(regnNumber string) (int, error) {
//
//}
