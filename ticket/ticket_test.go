package ticket

import (
	"github.com/stretchr/testify/assert"
	"parking_lot/slot"
	"parking_lot/vehicle"
	"testing"
)

func TestTicketReturnNoExceptionWhenCreateTicket(t *testing.T) {
	firstVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	firstSlot := slot.NewSlot()
	ticket, _ := NewTicket(firstVehicle, firstSlot)
	assert.NotNil(t, ticket)
}

func TestTicketReturnExceptionWhenVehicleIsNull(t *testing.T) {
	_, err := NewTicket(nil, slot.NewSlot())
	assert.Error(t, err)
}

func TestTicketReturnExceptionWhenSlotIsNull(t *testing.T) {
	firstVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	_, err := NewTicket(firstVehicle, nil)
	assert.Error(t, err)
}
