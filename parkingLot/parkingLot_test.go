package parkingLot

import (
	"github.com/stretchr/testify/assert"
	"parking_lot/vehicle"
	"testing"
)

func TestParkingLotReturnNoExceptionWhenCreateParkingLotWithFiveSlots(t *testing.T) {
	parkingLot, _ := NewParkingLot(5)
	assert.NotNil(t, parkingLot)
}

func TestParkingLotReturnExceptionWhenCreateParkingLotWithZeroSlots(t *testing.T) {
	_, err := NewParkingLot(0)
	assert.Error(t, err)
}

func TestParkingLotReturnExceptionWhenCreateParkingLotWithMinus5Slots(t *testing.T) {
	_, err := NewParkingLot(-5)
	assert.Error(t, err)
}

// Test for park method
func TestParkReturnNoExceptionWhenParkedVehicle(t *testing.T) {
	myParkingLot, _ := NewParkingLot(5)
	assert.NotNil(t, myParkingLot)
	myVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	assert.NotNil(t, myVehicle)
	myTicket, _ := myParkingLot.Park(myVehicle)
	assert.NotNil(t, myTicket)
}

func TestParkMethodReturnExceptionWhenSlotsAreFull(t *testing.T) {
	myParkingLot, _ := NewParkingLot(1)
	assert.NotNil(t, myParkingLot)
	myVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	assert.NotNil(t, myVehicle)
	myFirstTicket, _ := myParkingLot.Park(myVehicle)
	assert.NotNil(t, myFirstTicket)
	myVehicle, _ = vehicle.NewVehicle("KA-01-HH-9999", vehicle.RED)
	assert.NotNil(t, myVehicle)
	_, errSecondTicket := myParkingLot.Park(myVehicle)
	assert.Error(t, errSecondTicket)
}

func TestParkMethodReturnExceptionWhenVehicleIsNull(t *testing.T) {
	myParkingLot, _ := NewParkingLot(5)
	assert.NotNil(t, myParkingLot)
	_, err := myParkingLot.Park(nil)
	assert.Error(t, err)
}

// Test for unPark method
func TestUnParkMethodReturnNoExceptionWhenVehicleIsUnParked(t *testing.T) {
	myParkingLot, _ := NewParkingLot(5)
	assert.NotNil(t, myParkingLot)
	myVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.RED)
	assert.NotNil(t, myVehicle)
	myTicket, _ := myParkingLot.Park(myVehicle)
	assert.NotNil(t, myTicket)
	err := myParkingLot.UnPark(myTicket)
	assert.Nil(t, err)
}
