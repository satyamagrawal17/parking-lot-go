package parkingLot

import (
	"errors"
	"parking_lot/slot"
	"parking_lot/ticket"
	"parking_lot/vehicle"
)

type ParkingLot struct {
	slots              []*slot.Slot
	ticketToSlotMapper map[*ticket.Ticket]*slot.Slot
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, errors.New("invalid capacity")
	}
	slots := make([]*slot.Slot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = slot.NewSlot()
	}
	ticketToSlotMapper := make(map[*ticket.Ticket]*slot.Slot)
	return &ParkingLot{
		slots:              slots,
		ticketToSlotMapper: ticketToSlotMapper,
	}, nil
}

func (p *ParkingLot) findNearestSlot() *slot.Slot {
	for _, s := range p.slots {
		if s != nil {
			return s
		}
	}
	return nil
}

func (p *ParkingLot) Park(v *vehicle.Vehicle) (*ticket.Ticket, error) {
	if v == nil {
		return nil, errors.New("invalid vehicle")
	}

	nearestSlot := p.findNearestSlot()
	if nearestSlot != nil {
		slotErr := nearestSlot.Park(v)
		if slotErr != nil {
			return nil, slotErr
		}
		myTicket, ticketErr := ticket.NewTicket(v, nearestSlot)
		if ticketErr != nil {
			return nil, ticketErr
		}
		p.ticketToSlotMapper[myTicket] = nearestSlot
		return myTicket, nil

	}
	return nil, errors.New("parking lot is full")
}

func (p *ParkingLot) UnPark(t *ticket.Ticket) error {
	if t == nil {
		return errors.New("invalid ticket")
	}
	validErr := t.ValidateAndUnPark()
	if validErr != nil {
		return errors.New("invalid ticket")
	}
	parkedSlot, ok := p.ticketToSlotMapper[t]
	if !ok {
		return errors.New("invalid ticket")
	}
	unParkErr := parkedSlot.UnPark()
	if unParkErr != nil {
		return unParkErr
	}
	delete(p.ticketToSlotMapper, t)
	return nil
}
