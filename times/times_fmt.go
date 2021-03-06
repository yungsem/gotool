package times

import "time"

// Format 将 t 按照指定的格式（layout）格式化成字符串
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatDate 将 t 格式化成字符串
// 如：2021-03-26 21:44:12 -> "2021-03-26"
func FormatDate(t time.Time) string {
	return Format(t, LayoutDate)
}

// FormatDateTime 将 t 格式化成字符串
// 如：2021-03-26 21:44:12 -> "2021-03-26 21:44:12"
func FormatDateTime(t time.Time) string {
	return Format(t, LayoutDateTime)
}

// FormatHourMinute 将 t 格式化成字符串
// 如：2021-03-26 21:44:12 -> "21:44"
func FormatHourMinute(t time.Time) string {
	return Format(t, LayoutHourMinute)
}

