package moment

import (
	"fmt"
	"time"
)

// =====================================================================================================================
// 天

// StartOfDay 获取当天的开始时间
func StartOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Truncate 方法会将时间戳截断到指定的精度，这里我们将时间戳截断到天的开始
	// now := time.Now()
	// return now.Truncate(24 * time.Hour)
}

// StartOfDayStr 获取当天的开始时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func StartOfDayStr(f ...string) string {
	if len(f) > 0 {
		return StartOfDay().Format(f[0])
	}
	return StartOfDay().Format(time.DateTime)
}

// EndOfDay 获取当天的结束时间
func EndOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
}

// EndOfDayStr 获取当天的结束时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func EndOfDayStr(f ...string) string {
	if len(f) > 0 {
		return EndOfDay().Format(f[0])
	}
	return EndOfDay().Format(time.DateTime)
}

// =====================================================================================================================
// 周

// StartOfWeek 获取本周的开始时间
func StartOfWeek() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, offset)
	return weekStart
}

// StartOfWeekStr 获取本周的开始时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func StartOfWeekStr(f ...string) string {
	if len(f) > 0 {
		return StartOfWeek().Format(f[0])
	}
	return StartOfWeek().Format(time.DateTime)
}

// EndOfWeek 获取本周的结束时间
func EndOfWeek() time.Time {
	now := time.Now()
	offset := int(time.Sunday - now.Weekday())
	if offset < 0 {
		offset = 6
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()).AddDate(0, 0, offset)
	return weekEnd
}

// EndOfWeekStr 获取本周的结束时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func EndOfWeekStr(f ...string) string {
	if len(f) > 0 {
		return EndOfWeek().Format(f[0])
	}
	return EndOfWeek().Format(time.DateTime)
}

// IsWeekend 判断某一天是否为周末
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday 判断某一天是否为工作日
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// =====================================================================================================================
// 月

// StartOfMonth 获取本月的开始时间
func StartOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// StartOfMonthStr 获取本月的开始时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func StartOfMonthStr(f ...string) string {
	if len(f) > 0 {
		return StartOfMonth().Format(f[0])
	}
	return StartOfMonth().Format(time.DateTime)
}

// EndOfMonth 获取本月的结束时间
func EndOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 999999999, now.Location())
}

// EndOfMonthStr 获取本月的结束时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func EndOfMonthStr(f ...string) string {
	if len(f) > 0 {
		return EndOfMonth().Format(f[0])
	}
	return EndOfMonth().Format(time.DateTime)
}

// IsFirstDayOfMonth 判断是否是每月的第一天
func IsFirstDayOfMonth(t time.Time) bool {
	return t.Day() == 1
}

// IsLastDayOfMonth 判断是否是每月的最后一天
func IsLastDayOfMonth(t time.Time) bool {
	return t.Day() == t.AddDate(0, 1, 0).AddDate(0, 0, -1).Day()
}

// =====================================================================================================================
// 季

// StartOfQuarter 获取季度的开始日期
func StartOfQuarter(t time.Time) time.Time {
	quarter := GetQuarter(t)
	startMonth := (quarter-1)*3 + 1
	return time.Date(t.Year(), time.Month(startMonth), 1, 0, 0, 0, 0, t.Location())
}

// StartOfQuarterStr 获取季度的开始日期
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func StartOfQuarterStr(t time.Time, f ...string) string {
	if len(f) > 0 {
		return StartOfQuarter(t).Format(f[0])
	}
	return StartOfQuarter(t).Format(time.DateTime)
}

// EndOfQuarter 获取季度的结束日期
func EndOfQuarter(t time.Time) time.Time {
	quarter := GetQuarter(t)
	endMonth := quarter * 3
	nextYear := t.Year()
	if endMonth == 12 {
		endMonth = 1
		nextYear++
	}
	endDate := time.Date(nextYear, time.Month(endMonth), 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)
	return time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, t.Location())
}

// EndOfQuarterStr 获取季度的结束日期
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func EndOfQuarterStr(t time.Time, f ...string) string {
	if len(f) > 0 {
		return EndOfQuarter(t).Format(f[0])
	}
	return EndOfQuarter(t).Format(time.DateTime)
}

// GetQuarter 获取指定日期所在的季度
func GetQuarter(t time.Time) int {
	month := int(t.Month())
	return (month-1)/3 + 1
}

// IsFirstDayOfQuarter 判断是否是季度的第一天
func IsFirstDayOfQuarter(t time.Time) bool {
	return t.Day() == 1 && t.Month() == time.Month((GetQuarter(t)-1)*3+1)
}

// IsLastDayOfQuarter 判断是否是季度的最后一天
func IsLastDayOfQuarter(t time.Time) bool {
	return t.Day() == 30 && t.Month() == time.Month(GetQuarter(t)*3)
}

// =====================================================================================================================
// 年

// StartOfYear 获取今年的开始时间
func StartOfYear() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
}

// StartOfYearStr 获取今年的开始时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func StartOfYearStr(f ...string) string {
	if len(f) > 0 {
		return StartOfYear().Format(f[0])
	}
	return StartOfYear().Format(time.DateTime)
}

// EndOfYear 获取今年的结束时间
func EndOfYear() time.Time {
	now := time.Now()
	return time.Date(now.Year()+1, 1, 0, 23, 59, 59, 999999999, now.Location())
}

// EndOfYearStr 获取今年的结束时间
// f 为可选参数，用于指定时间格式，默认为 time.DateTime
func EndOfYearStr(f ...string) string {
	if len(f) > 0 {
		return EndOfYear().Format(f[0])
	}
	return EndOfYear().Format(time.DateTime)
}

// IsFirstDayOfYear 判断是否是今年的第一天
func IsFirstDayOfYear(t time.Time) bool {
	return t.Day() == 1 && t.Month() == 1
}

// IsLastDayOfYear 判断是否是今年的最后一天
func IsLastDayOfYear(t time.Time) bool {
	return t.Day() == 31 && t.Month() == 12
}

// IsLeapYear 判断是否为闰年
// 闰年的判断规则为：
// 1. 能被4整除但不能被100整除的年份是闰年。
// 2. 能被400整除的年份也是闰年。
// 3. 其他年份都不是闰年。
// 例如：2000年是闰年，1900年不是闰年，2004年是闰年，2005年不是闰年。
// 闰年的作用是为了调整公历的日期，使得每年的天数保持不变。
// 闰年的平均年份长度为365.2425天，比平年多0.0078天。
// 因此，闰年的影响是在每年的2月29日增加一天，使得2月的天数变为29天。
// 平年的2月只有28天。
// 闰年的影响是为了调整公历的日期，使得每年的天数保持不变。
// 因此，闰年的影响是为了调整公历的日期，使得每年的天数保持不变。
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// =====================================================================================================================
// Other

// IsSameDay 判断两个时间是否为同一天
func IsSameDay(first, second time.Time) bool {
	return first.YearDay() == second.YearDay() && first.Year() == second.Year()
}

// DiffInDays 计算两个时间之间的天数差
func DiffInDays(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// ConvertTimeZone 时区转换
func ConvertTimeZone(t time.Time, to *time.Location) time.Time {
	return t.In(to)
}

// CurrentTimeInLocation 获取特定时区的当前时间
func CurrentTimeInLocation(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

// TimeAgo 将时间转换为多久之前的形式
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff < time.Minute {
		return "刚刚"
	} else if diff < time.Hour {
		return fmt.Sprintf("%d 分钟前", int(diff.Minutes()))
	} else if diff < 24*time.Hour {
		return fmt.Sprintf("%d 小时前", int(diff.Hours()))
	} else if diff < 30*24*time.Hour {
		return fmt.Sprintf("%d 天前", int(diff.Hours()/24))
	} else if diff < 365*24*time.Hour {
		return fmt.Sprintf("%d 个月前", int(diff.Hours()/24/30))
	}
	return fmt.Sprintf("%d 年前", int(diff.Hours()/24/365))
}

// GetDatesInRange 获取指定时间段内的日期列表
func GetDatesInRange(start, end time.Time) []time.Time {
	var dates []time.Time
	current := start
	for !current.After(end) {
		dates = append(dates, current)
		current = current.AddDate(0, 0, 1)
	}
	return dates
}

// DayOfYear 计算指定日期是该年的第几天
func DayOfYear(t time.Time) int {
	return t.YearDay()
}
