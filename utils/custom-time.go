package utils

import "time"

//format time

func TimeFormatting(input string) string {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	return t.Format(layout)
}

//primitive date time to string
func DateTimeToString(input time.Time) string {
	layout := "2006-01-02 15:04:05"
	return input.Format(layout)
}
