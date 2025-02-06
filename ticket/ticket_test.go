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

func TestTicketReturnExceptionWhenVehicleIsNull(t *testing.T) {
	_, err := NewTicket(nil, slot.NewSlot())
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestTicketReturnExceptionWhenSlotIsNull(t *testing.T) {
	firstVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	_, err := NewTicket(firstVehicle, nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
