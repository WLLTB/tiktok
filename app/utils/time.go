package utils

import "time"

func FormatTime(time time.Time, timeFormat string) string {
	return time.Format(timeFormat)
}

func ParseTime(value string, timeFormat string) (time.Time, error) {
	return time.Parse(timeFormat, value)
}
