package date

import "time"

const (
	apiDateLayout = "2022-01-08 15:04:01"
)

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
func GetNow() time.Time {
	return time.Now().UTC()
}
