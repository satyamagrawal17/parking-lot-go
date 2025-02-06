package slot

import (
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
