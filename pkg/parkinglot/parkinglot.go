package parkinglot

import (
	"parking/pkg/slot"
)

type Parkinglot interface {
	GetAllSlots() []slot.Slot
	GetASlot(slot.Vehicle) []slot.Slot
	BookAslot(slot.Location) bool
	FreeASlot(slot.Location) bool
}

var CinemaHall RegularParkingLot

type RegularParkingLot struct {
}

func (RegularParkingLot) GetAllSlots() []slot.Slot {
	//get * from parkinglot
	return nil
}

func (RegularParkingLot) GetASlot(slot.Vehicle) []slot.Slot {
	//get type=vehicle from parkinglot
	return nil
}

func (RegularParkingLot) BookAslot(slot.Location) bool {
	return false
}

func (RegularParkingLot) FreeASlot(slot.Location) bool {
	return false
}
