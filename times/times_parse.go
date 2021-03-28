package times

import "time"

const (
	LayoutDate          = "2006-01-02"
	LayoutDateTime      = "2006-01-02 15:04:05"
	LayoutHourMinute    = "15:04"
	LayoutDateTimeMil    = "2006-01-02 15:04:05.99999"
	LayoutSlashDate     = "2006/01/02"
	LayoutSlashDateTime = "2006/01/02 15:04:05"
)

// Parse 将字符串格式的时间（value）按照指定的格式（layout）转换为 time.Time
func parse(timeStr string, layout string) (time.Time, error) {
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return t, err
	}
	return t, nil
}

// ParseDate 将字符串格式的日期转换为 time.Time
// 如："2006-01-02" -> 2006-01-02 ，时分秒都为 0
func ParseDate(timeStr string) (time.Time, error) {
	return parse(timeStr, LayoutDate)
}

// ParseDate 将字符串格式的日期时间转换为 time.Time
// 如："2006-01-02 15:04:05" -> 2006-01-02 15:04:05
func ParseDateTime(timeStr string) (time.Time, error) {
	return parse(timeStr, LayoutDateTime)
}


