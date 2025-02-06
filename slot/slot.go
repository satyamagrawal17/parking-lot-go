package slot

import (
	"errors"
	"parking_lot/vehicle"
)

type Slot struct {
	vehicle *vehicle.Vehicle
}

func NewSlot() *Slot {
	return &Slot{
		vehicle: nil,
	}
}

func (s *Slot) Park(v *vehicle.Vehicle) error {
	if s.vehicle != nil {
		return errors.New("Slot is already occupied")
	}
	s.vehicle = v
	return nil
}
