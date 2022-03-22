package date

import (
	"testing"

	"dev.io/v1/pkg/time"
)

func TestValidator(t *testing.T) {
	ans := DateValidator(time.Date(2000, 1, 01, 0, 0, 0))
	if !ans {
		t.Errorf("Invalid date")
	}
}
