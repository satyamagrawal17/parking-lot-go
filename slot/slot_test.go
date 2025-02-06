package slot

import "testing"

func TestSlotReturnNoExceptionWhenCreateSlot(t *testing.T) {
	slot := NewSlot()
	if slot == nil {
		t.Errorf("Expected slot instance, got nil")
	}
}
