package timer

import "time"

func GetNow() time.Time {
	return time.Now()
}

func GetCalcTime(current time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return current.Add(duration), nil
}
