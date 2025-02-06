package ticket

import (
	"parking_lot/slot"
	"parking_lot/vehicle"
	"testing"
)

func TestTicketReturnNoExceptionWhenCreateTicket(t *testing.T) {
	firstVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	firstSlot := slot.NewSlot()
	ticket, _ := NewTicket(firstVehicle, firstSlot)
	if ticket == nil {
		t.Errorf("Expected ticket instance, got nil")
	}
}
