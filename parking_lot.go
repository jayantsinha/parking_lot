package main

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
	p.Slots = append(p.Slots, &Slot{Vhcl: &Vehicle{RegnNumber:regnNumber, Color:color}, Num: slotToFill})
	return slotToFill+1, nil
}

//func (p *ParkingLot) Leave(slotNum int) (int, error) {
//
//}

//func (p *ParkingLot) Status() []*Slot {
//
//}
//
//func (p *ParkingLot) FindSlotNumbersByColor(color string) ([]int, error) {
//
//}
//
//func (p *ParkingLot) FindRegistrationNumbersByColor(color string) ([]string, error) {
//
//}
//
//func (p *ParkingLot) FindSlotByRegistrationNumber(regnNumber string) (int, error) {
//
//}