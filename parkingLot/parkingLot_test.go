package parkingLot

import "testing"

func TestParkingLotReturnNoExceptionWhenCreateParkingLotWithFiveSlots(t *testing.T) {
	parkingLot, _ := NewParkingLot(5)
	if parkingLot == nil {
		t.Errorf("Expected parkingLot instance, got nil")
	}
}

func TestParkingLotReturnExceptionWhenCreateParkingLotWithZeroSlots(t *testing.T) {
	_, err := NewParkingLot(0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestParkingLotReturnExceptionWhenCreateParkingLotWithMinus5Slots(t *testing.T) {
	_, err := NewParkingLot(-5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
