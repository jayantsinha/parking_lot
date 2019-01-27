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

func (p *ParkingLot) Init(numSlots int) int {
	p.Slots = make([]*Slot, numSlots)
	for i := 0; i < numSlots; i++ {
		p.Slots[i] = &Slot{
			Num:  i,
			Vhcl: nil,
		}
	}
	p.IsFull = false
	return len(p.Slots)
}

//func (p *ParkingLot) Park(regnNumber, color string) (int, error) {
//
//}

//func (p *ParkingLot) Leave(slotNum int) (int, error) {
//
//}
//
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