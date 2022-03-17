package date

import "time"

func Distance(secondDateTime time.Time, firstDateTime time.Time) (int, error) {
	result := int(secondDateTime.Unix()) - int(firstDateTime.Unix())
	result = result / (24 * 3600)
	return result - 1, nil
}
