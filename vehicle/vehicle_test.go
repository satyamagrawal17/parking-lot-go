package vehicle

import (
	"testing"
)

func TestVehicleReturnNoExceptionWhenCreateVehicle(t *testing.T) {
	var vehicle *Vehicle = NewVehicle("KA-01-HH-1234", RED)

	if vehicle == nil {
		t.Errorf("Expected vehicle instance, got nil")
	}

}
