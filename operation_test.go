package aifinitsdk

import (
	"fmt"
	"testing"
)

func TestDoorOpenCloseStatus_String(t *testing.T) {
	tests := []struct {
		name     string
		status   DoorOpenCloseStatus
		expected string
	}{
		{
			name:     "Success - Door Opened",
			status:   DoorOpenCloseStatusOpened,
			expected: "Door opened successfully",
		},
		{
			name:     "Success - Door Closed",
			status:   DoorOpenCloseStatusClosed,
			expected: "Door closed successfully",
		},
		{
			name:     "Shopping Not Finished",
			status:   DoorOpenCloseStatusShoppingNotFinished,
			expected: "Failed to open - previous shopping not finished",
		},
		{
			name:     "Power Off",
			status:   DoorOpenCloseStatusPowerOff,
			expected: "Failed to open - device power off, running on UPS",
		},
		{
			name:     "Unknown Status",
			status:   DoorOpenCloseStatus(9999),
			expected: "Unknown status code: 9999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.status.String()
			if got != tt.expected {
				t.Errorf("DoorOpenCloseStatus.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func ExampleDoorOpenCloseStatus_String_basic() {
	status := DoorOpenCloseStatusOpened
	fmt.Println(status)
	// Output: Door opened successfully
}

func ExampleDoorOpenCloseStatus_String_error() {
	err := fmt.Errorf("door operation failed: %s", DoorOpenCloseStatusPowerOff)
	fmt.Println(err)
	// Output: door operation failed: Failed to open - device power off, running on UPS
}

func ExampleDoorOpenCloseStatus_String_format() {
	fmt.Printf("Current door status: %s\n", DoorOpenCloseStatusShoppingNotFinished)
	// Output: Current door status: Failed to open - previous shopping not finished
}
