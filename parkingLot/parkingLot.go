package parkingLot

import (
	"errors"
	"parking_lot/slot"
)

type ParkingLot struct {
	slots              []slot.Slot
	ticketToSlotMapper map[Ticket]slot.Slot
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, errors.New("Invalid capacity")
	}
	slots := make([]slot.Slot, capacity)
	ticketToSlotMapper := make(map[Ticket]slot.Slot)
	return &ParkingLot{
		slots:              slots,
		ticketToSlotMapper: ticketToSlotMapper,
	}, nil
}
