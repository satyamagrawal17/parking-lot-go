package slot

import (
	"github.com/stretchr/testify/assert"
	"parking_lot/vehicle"
	"testing"
)

func TestSlotReturnNoExceptionWhenCreateSlot(t *testing.T) {
	slot := NewSlot()
	assert.NotNil(t, slot)
}

func TestParkReturnExceptionWhenSlotIsOccupied(t *testing.T) {
	firstSlot := NewSlot()
	firstVehicle, _ := vehicle.NewVehicle("KA-01-HH-1234", vehicle.GREEN)
	firstErr := firstSlot.Park(firstVehicle)
	assert.NoError(t, firstErr)

	secondErr := firstSlot.Park(firstVehicle)
	assert.Error(t, secondErr)
}
