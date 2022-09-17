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

// dateFormat 日期格式化处理
func dateFormat(date string) string {
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

// DateUtilSingle 时间格式化处理工具，将开始时间和结束时间统一成一个时间
func DateUtilSingle(start, end string) string {
	st, _ := time.Parse(constant.StandardTime, dateFormat(start))
	ed, _ := time.Parse(constant.StandardTime, dateFormat(end))
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
		formatTime = fmt.Sprintf("%s%d-", formatTime, time.Month())
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
