package driver

import (
	"testing"
)

func TestController(t *testing.T) {
	_, err := convertRequestCapacityBytes(-1, 10)
	if err == nil {
		t.Error("should be error")
	}

	_, err = convertRequestCapacityBytes(10, -1)
	if err == nil {
		t.Error("should be error")
	}

	_, err = convertRequestCapacityBytes(20, 10)
	if err == nil {
		t.Error("should be error")
	}

	v, err := convertRequestCapacityBytes(0, 10)
	if err != nil {
		t.Error("should not be error")
	}
	if v != 1 {
		t.Errorf("should be 1: %d", v)
	}

	v, err = convertRequestCapacityBytes(1, 0)
	if err != nil {
		t.Error("should not be error")
	}
	if v != 1 {
		t.Errorf("should be 1: %d", v)
	}

	v, err = convertRequestCapacityBytes(1<<30, 1<<30)
	if err != nil {
		t.Error("should not be error")
	}
	if v != (1 << 30) {
		t.Errorf("should be 1: %d", v)
	}

	v, err = convertRequestCapacityBytes(1<<30+1, 1<<30+1)
	if err != nil {
		t.Error("should not be error")
	}
	if v != (1<<30 + 1) {
		t.Errorf("should be 2: %d", v)
	}
}
