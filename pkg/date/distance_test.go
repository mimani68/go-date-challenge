package date

import (
	"testing"
	"time"
)

func TestDistance(t *testing.T) {
	start := time.Date(1900, time.Month(1), 01, 0, 0, 1, 0, time.UTC)
	end := time.Date(2010, time.Month(1), 01, 0, 0, 1, 0, time.UTC)
	ans, _ := Distance(end, start)
	if ans <= 0 {
		t.Errorf("Distance finding has problem")
	}
}
