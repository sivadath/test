package slot

import "math/rand"

type Vehicle string

const (
	TwoWheeler  Vehicle = "2-wheeler"
	Fourwheeler Vehicle = "4-wheeler"
)

type Location struct {
	x int64
	y int64
}

type Slot struct {
	Location
	availability bool
	slotType     Vehicle
}

func New() *Slot {
	return &Slot{Location: Location{x: rand.Int63(), y: rand.Int63()}}
}

func NewSlotOfVehicle(v Vehicle) *Slot {
	return &Slot{Location: Location{x: rand.Int63(), y: rand.Int63()}, slotType: v}
}

func (s *Slot) UpdateKind(v Vehicle) {
	s.slotType = v
}

func (s *Slot) GetAvailability() bool {
	return s.availability
}

func (s *Slot) GetType() Vehicle {
	return s.slotType
}

func (s *Slot) GetLocation() Location {
	return s.Location
}
