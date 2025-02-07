package ticket

import (
	"errors"
	"parking_lot/slot"
	"parking_lot/vehicle"
)

type Ticket struct {
	vehicle *vehicle.Vehicle
	slot    *slot.Slot
}

func NewTicket(vehicle *vehicle.Vehicle, slot *slot.Slot) (*Ticket, error) {
	if vehicle == nil {
		return nil, errors.New("invalid vehicle")
	}
	if slot == nil {
		return nil, errors.New("invalid slot")
	}
	return &Ticket{
		vehicle: vehicle,
		slot:    slot,
	}, nil
}
