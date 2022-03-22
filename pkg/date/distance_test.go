package date

import (
	"testing"

	"dev.io/v1/pkg/time"
)

func TestDistance(t *testing.T) {
	start := time.Date(1900, 1, 01, 0, 0, 1)
	end := time.Date(2010, 1, 01, 0, 0, 1)
	ans, _ := Distance(end, start)
	if ans <= 0 {
		t.Errorf("Distance finding has problem")
	}
}
