package time

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/shootz-developer/gtool/constant"
)

const (
	HourTime   = "hour"
	MinuteTime = "minute"
	SecondTime = "second"
)

const (
	Year            = "2006"
	Month           = "01"
	Day             = "02"
	Hour            = "15"
	Minute          = "04"
	Second          = "05"
	Complete        = "2006-01-02 15:05:05"
	StringToTimeOne = "2006-01-02 15:04:05"
	StringToTimeTow = "2006-01-02"
)

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15:04",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyy":                "2006",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"mm:ss":               "04:05",
	}
}

// TodayBeginTimestamp 返回今天零点的时间戳
func TodayBeginTimestamp() int64 {
	now := time.Now()
	return now.Unix() - (int64)(now.Hour()*60*60+now.Minute()*60+now.Second())
}

// QueryBeforeTime 获取当前时间之前的某个时间点
func QueryBeforeTime(num int, flag string) string {
	currentTime := time.Now()
	deal := "-"
	var m time.Duration
	if flag == HourTime {
		m, _ = time.ParseDuration(deal + strconv.Itoa(num) + "m")
	} else if flag == MinuteTime {
		m, _ = time.ParseDuration(deal + strconv.Itoa(num) + "h")
	} else if flag == SecondTime {
		m, _ = time.ParseDuration(deal + strconv.Itoa(num) + "s")
	} else {
		return ""
	}
	result := currentTime.Add(m)
	return result.Format(Complete)
}

// QueryAfterTime 获取当前时间之后的某个时间点
func QueryAfterTime(num int, flag string) string {
	currentTime := time.Now()
	var m time.Duration
	if flag == HourTime {
		m, _ = time.ParseDuration(strconv.Itoa(num) + "m")
	} else if flag == MinuteTime {
		m, _ = time.ParseDuration(strconv.Itoa(num) + "h")
	} else if flag == SecondTime {
		m, _ = time.ParseDuration(strconv.Itoa(num) + "s")
	} else {
		return ""
	}
	result := currentTime.Add(m)
	return result.Format(Complete)
}

// QueryNowTime 获取当前时间
func QueryNowTime() string {
	dateTime := time.Now().Format(Complete)
	return dateTime
}

// JudgeTimeOrder 判断一个时间是否在一个时间之后
func JudgeTimeOrder(t string) bool {
	stringTime, _ := time.Parse(Complete, t)
	beforeOrAfter := stringTime.After(time.Now())
	return beforeOrAfter
}

// GetWeekDay 获取周几方法
func GetWeekDay(time time.Time) int {
	return int(time.Weekday())
}

// MinuteAddOrSub 时间分钟加减计算
func MinuteAddOrSub(t time.Time, num int64) time.Time {
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "m")
	return t.Add(m)
}

// HourAddOrSub 时间小时加减计算
func HourAddOrSub(t time.Time, num int64) time.Time {
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "h")
	return t.Add(m)
}

// DayAddOrSub 时间天加减计算
func DayAddOrSub(t time.Time, num int64) time.Time {
	num = num * 24
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "h")
	return t.Add(m)
}

// DateFormat 日期格式化处理
func DateFormat(date string) string {
	newDate := ""
	for i, _ := range date {
		if date[i] == 'T' {
			newDate = fmt.Sprintf("%s ", newDate)
		} else if date[i] == 'Z' {
			continue
		} else {
			newDate = fmt.Sprintf("%s%c", newDate, date[i])
		}
	}

	return newDate
}

// DateUtil 时间格式化处理工具，将开始时间和结束时间统一成一个时间
func DateUtil(start, end string) string {
	st, _ := time.Parse(constant.StandardTime, DateFormat(start))
	ed, _ := time.Parse(constant.StandardTime, DateFormat(end))
	time := fmt.Sprintf("%d", st.Year())
	if int(st.Month()) < 10 {
		time = fmt.Sprintf("%s-0%d", time, int(st.Month()))
	} else {
		time = fmt.Sprintf("%s-%d", time, int(st.Month()))
	}

	if st.Day() < 10 {
		time = fmt.Sprintf("%s-0%d", time, st.Day())
	} else {
		time = fmt.Sprintf("%s-%d", time, st.Day())
	}

	time = fmt.Sprintf("%s(%s)", time, constant.WeekDayMap[st.Weekday().String()])

	if st.Hour() < 10 {
		time = fmt.Sprintf(" %s 0%d", time, st.Hour())
	} else {
		time = fmt.Sprintf(" %s %d", time, st.Hour())
	}

	if st.Minute() < 10 {
		time = fmt.Sprintf("%s:0%d", time, st.Minute())
	} else {
		time = fmt.Sprintf("%s:%d", time, st.Minute())
	}

	if ed.Hour() < 10 {
		time = fmt.Sprintf("%s~0%d", time, ed.Hour())
	} else {
		time = fmt.Sprintf("%s~%d", time, ed.Hour())
	}

	if ed.Minute() < 10 {
		time = fmt.Sprintf("%s:0%d", time, ed.Minute())
	} else {
		time = fmt.Sprintf("%s:%d", time, ed.Minute())
	}

	return time
}

// FormatStartTime 格式化开始时间
func FormatStartTime(startTime string) string {
	time, _ := time.Parse(constant.StandardTime, startTime)
	formatTime := fmt.Sprintf("%d", time.Year())
	if time.Month() < 10 {
		formatTime = fmt.Sprintf("%s-0%d", formatTime, time.Month())
	} else {
		formatTime = fmt.Sprintf("%s-%d", formatTime, time.Month())
	}

	if time.Day() < 10 {
		formatTime = fmt.Sprintf("%s-0%d", formatTime, time.Day())
	} else {
		formatTime = fmt.Sprintf("%s-%d", formatTime, time.Day())
	}

	formatTime = fmt.Sprintf("%s(%s) ", formatTime, constant.WeekDayMap[time.Weekday().String()])

	return formatTime
}

// FormatEndTime 格式化结束时间
func FormatEndTime(startTime, endTime string) string {
	stTime, err := time.Parse(constant.StandardTime, startTime)
	if err != nil {
		log.Printf("Parse startTime err: [%+v],startTime = %s", err, startTime)
	}
	edTime, err := time.Parse(constant.StandardTime, endTime)
	if err != nil {
		log.Printf("Parse endTime err: [%+v],endTime = %s", err, endTime)
	}

	formatTime := ""
	if stTime.Hour() < 10 {
		formatTime = fmt.Sprintf("0%d:", stTime.Hour())
	} else {
		formatTime = fmt.Sprintf("%d:", stTime.Hour())

	}

	if stTime.Minute() < 10 {
		formatTime = fmt.Sprintf("%s0%d~", formatTime, stTime.Minute())
	} else {
		formatTime = fmt.Sprintf("%s%d~", formatTime, stTime.Minute())
	}

	if edTime.Hour() < 10 {
		formatTime = fmt.Sprintf("%s0%d:", formatTime, edTime.Hour())
	} else {
		formatTime = fmt.Sprintf("%s%d:", formatTime, edTime.Hour())

	}

	if edTime.Minute() < 10 {
		formatTime = fmt.Sprintf("%s0%d", formatTime, edTime.Minute())
	} else {
		formatTime = fmt.Sprintf("%s%d", formatTime, edTime.Minute())

	}

	return formatTime
}

// AddMinute 加减分钟数
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// GetZeroHourTimestamp 获得开始时间时间戳
func GetZeroHourTimestamp() int64 {
	ts := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", ts)
	return t.UTC().Unix() - 8*3600
}

// GetNightTimestamp 获得晚上时间时间戳
func GetNightTimestamp() int64 {
	return GetZeroHourTimestamp() + 86400 - 1
}

// FormatTimeToStr 将时间转换为字符串
func FormatTimeToStr(t time.Time, format string) string {
	return t.Format(timeFormat[format])
}

// FormatStrToTime 将字符串转换为时间
func FormatStrToTime(str, format string) (time.Time, error) {
	v, ok := timeFormat[format]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not found", format)
	}

	return time.Parse(v, str)
}

// BeginOfMinute 一天开始的那分钟
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute  一天结束的那分钟
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour 一天开始的那小时
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour 一天结束的那小时
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay 一天开始
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay 一天结束
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek 一周开始，从周日算
func BeginOfWeek(t time.Time) time.Time {
	y, m, d := t.AddDate(0, 0, 0-int(BeginOfDay(t).Weekday())).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfWeek 一周结束，从周六算
func EndOfWeek(t time.Time) time.Time {
	y, m, d := BeginOfWeek(t).AddDate(0, 0, 7).Add(-time.Nanosecond).Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfMonth 一个月的开始
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 一个月的结束
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear 一年的开始
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 一年的结束
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// FomratTime 格式化时间
func FomratTime(unformatTime string) string {
	fomatTime := ""
	for i := 0; i < len(unformatTime); i++ {
		if unformatTime[i] == 'T' {
			fomatTime = fmt.Sprintf("%s ", fomatTime)
		} else if unformatTime[i] == 'Z' {
			continue
		} else {
			fomatTime = fmt.Sprintf("%s%c", fomatTime, unformatTime[i])
		}
	}
	return fomatTime
}
