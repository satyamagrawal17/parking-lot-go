package parkingLot

import (
	"errors"
	"parking_lot/slot"
	"parking_lot/ticket"
)

type ParkingLot struct {
	slots              []*slot.Slot
	ticketToSlotMapper map[*ticket.Ticket]*slot.Slot
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, errors.New("Invalid capacity")
	}
	slots := make([]*slot.Slot, capacity)
	ticketToSlotMapper := make(map[*ticket.Ticket]*slot.Slot)
	return &ParkingLot{
		slots:              slots,
		ticketToSlotMapper: ticketToSlotMapper,
	}, nil
}
