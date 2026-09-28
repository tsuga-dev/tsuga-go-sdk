package tsuga

import (
	"testing"
	"time"
)

func TestBoolValue(t *testing.T) {
	if BoolValue(nil) != false {
		t.Errorf("Expected false for nil")
	}
	v := true
	if BoolValue(&v) != true {
		t.Errorf("Expected true")
	}
}

func TestIntValue(t *testing.T) {
	if IntValue(nil) != 0 {
		t.Errorf("Expected 0 for nil")
	}
	v := 42
	if IntValue(&v) != 42 {
		t.Errorf("Expected 42")
	}
}

func TestStringValue(t *testing.T) {
	if StringValue(nil) != "" {
		t.Errorf("Expected empty string for nil")
	}
	v := "hello"
	if StringValue(&v) != "hello" {
		t.Errorf("Expected 'hello'")
	}
}

func TestTimeValue(t *testing.T) {
	if !TimeValue(nil).IsZero() {
		t.Errorf("Expected zero time for nil")
	}
	now := time.Now()
	if !TimeValue(&now).Equal(now) {
		t.Errorf("Expected equal time")
	}
}
