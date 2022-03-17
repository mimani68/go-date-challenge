package date

import (
	"testing"
	"time"
)

func TestValidator(t *testing.T) {
	ans := DateValidator(time.Date(2000, time.Month(1), 01, 0, 0, 1, 0, time.UTC))
	if !ans {
		t.Errorf("Invalid date")
	}
}
