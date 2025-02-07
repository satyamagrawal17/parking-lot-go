package vehicle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVehicleReturnNoExceptionWhenCreateVehicle(t *testing.T) {
	vehicle, _ := NewVehicle("KA-01-HH-1234", RED)

	//if vehicle == nil {
	//	t.Errorf("Expected vehicle instance, got nil")
	//}
	assert.NotNil(t, vehicle)
}
func TestVehicleReturnExceptionWhenRegistrationNumberIsEmpty(t *testing.T) {
	_, err := NewVehicle("", RED)
	assert.Error(t, err)

}

func TestVehicleReturnExceptionWhenColorIsInvalid(t *testing.T) {
	_, err := NewVehicle("KA-01-HH-1234", 0)
	assert.Error(t, err)
}
