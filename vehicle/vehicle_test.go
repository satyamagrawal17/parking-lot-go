package vehicle

import (
	"testing"
)

func TestVehicleReturnNoExceptionWhenCreateVehicle(t *testing.T) {
	vehicle, _ := NewVehicle("KA-01-HH-1234", RED)

	if vehicle == nil {
		t.Errorf("Expected vehicle instance, got nil")
	}
}
func TestVehicleReturnExceptionWhenRegistrationNumberIsEmpty(t *testing.T) {
	_, err := NewVehicle("", RED)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestVehicleReturnExceptionWhenColorIsInvalid(t *testing.T) {
	_, err := NewVehicle("KA-01-HH-1234", 0)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
