package times

import "time"

const (
	Spring = 1 + iota
	Summer
	Autumn
	Winter
)

type TimeInfo struct {
	Year      int
	Season    int
	Month     int
	YearWeek  int
	MonthWeek int
	Weekday   int
}

// Info 返回 t 对应的详细信息：年份，季度，月份，年周，月周，周几
func Info(t time.Time) *TimeInfo {
	info := &TimeInfo{
		Year:    t.Year(),
		Month:   int(t.Month()),
		Weekday: int(t.Weekday()),
	}
	// Season
	info.Season = Season(t)
	// YearWeek
	info.YearWeek = YearWeek(t)
	// MonthWeek
	info.MonthWeek = MonthWeek(t)
	return info
}

// Season 返回 t 对应的季度
func Season(t time.Time) int {
	month := int(t.Month())
	if month <= int(time.March) {
		return Spring
	} else if month <= int(time.June) {
		return Summer
	} else if month <= int(time.September) {
		return Autumn
	} else {
		return Winter
	}
}

// YearWeek 返回 t 对应的年周
func YearWeek(t time.Time) int {
	_, w := t.ISOWeek()
	return w
}

// MonthWeek 返回 t 对应的月周
func MonthWeek(t time.Time) int {
	// t 所在月份的第一天
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	// 第一天对应的年周
	_, firstDayYearWeek := firstDay.ISOWeek()
	return YearWeek(t) - firstDayYearWeek + 1
}

// 根据年份和年周获取该年周的开始时间
func FirstDayOfYearWeek(year int, yearWeek int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.Local)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < yearWeek { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

// WeekRange 表示一周的时间范围
type WeekRange struct {
	start time.Time
	end   time.Time
}

// RangeOfYearWeek 根据年份和年周获取年周的范围
func RangeOfYearWeek(year int, yearWeek int) *WeekRange {
	start := FirstDayOfYearWeek(year, yearWeek)
	end := start.AddDate(0, 0, 6)
	return &WeekRange{
		start: start,
		end:   end,
	}
}

// ZeroTime 将 t 的时间置为 0 ，日期保持不变
func ZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.YearDay(), 0, 0, 0, 0, t.Location())
}

// AddDays 将 t 增加 days 天
// days 可以为负数，表示减少 days 天
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// Gap 计算 end 和 start 之间的时间间隔，即 end - start
func Gap(start time.Time, end time.Time) time.Duration {
	return end.Sub(start)
}
