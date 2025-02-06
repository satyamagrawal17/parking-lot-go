package parkingLot

import "testing"

func TestParkingLotReturnNoExceptionWhenCreateParkingLotWithFiveSlots(t *testing.T) {
	parkingLot, _ := NewParkingLot(5)
	if parkingLot == nil {
		t.Errorf("Expected parkingLot instance, got nil")
	}
}
