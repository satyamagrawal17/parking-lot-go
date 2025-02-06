package parkingLot

import "testing"

func TestParkingLotReturnNoExceptionWhenCreateParkingLotWithFiveSlots(t *testing.T) {
	var parkingLot ParkingLot = newParkingLot(5)
}
