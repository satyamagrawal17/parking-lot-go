package slot

import (
	"parking_lot/vehicle"
	"testing"
)

func TestSlotReturnNoExceptionWhenCreateSlot(t *testing.T) {
	slot := NewSlot()
	if slot == nil {
		t.Errorf("Expected slot instance, got nil")
	}
}

func TestParkReturnExceptionWhenSlotIsOccupied(t *testing.T) {
	firstSlot := NewSlot()
	firstVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.GREEN)
	firstSlot.Park(firstVehicle)
	error := firstSlot.Park(firstVehicle)
	if error == nil {
		t.Errorf("Expected Error, got nil")
	}
}
