package jobcan

import (
	"flag"
	"strconv"
	"time"
)

var (
	now time.Time
)

func init() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now = time.Now().In(loc)
}

type userInput struct {
	Email    string
	Password string
}

func newLoginUser() userInput {
	return userInput{
		Email:    loginEmail(),
		Password: loginPassword(),
	}
}

type attendanceInput struct {
	Year  string
	Month string
}

var (
	atndYear  string
	atndMonth string
)

func init() {
	defAtndYear, defAtndMonth := defautAttendYearMonth()
	flag.StringVar(&atndYear, "atndYear", defAtndYear, "")
	flag.StringVar(&atndMonth, "atndMonth", defAtndMonth, "")
}

func defautAttendYearMonth() (string, string) {
	nowYearInt := now.Year()
	nowMonthInt := int(now.Month())
	var (
		atndYearInt  = nowYearInt
		atndMonthInt = nowMonthInt
	)
	if now.Day() > 15 {
		if nowMonthInt == 12 {
			atndYearInt = nowYearInt + 1
			atndMonthInt = 1
		} else {
			atndMonthInt = nowMonthInt + 1
		}
	}
	monthStr := strconv.Itoa(atndMonthInt)
	if len(monthStr) == 1 {
		monthStr = "0" + monthStr
	}
	return strconv.Itoa(atndYearInt), monthStr
}

func newEditAttendance() attendanceInput {
	return attendanceInput{
		Year:  atndYear,
		Month: atndMonth,
	}
}

var (
	applyMonth string
	applyDay   string
)

func init() {
	defApplyMonth, defApplyDay := defautApplyMonthDay()
	flag.StringVar(&applyMonth, "applyMonth", defApplyMonth, "")
	flag.StringVar(&applyDay, "applyDay", defApplyDay, "")
}

func defautApplyMonthDay() (string, string) {
	return strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day())
}

type applyInput struct {
	Month string
	Day   string
}

func newApply() applyInput {
	return applyInput{
		Month: applyMonth,
		Day:   applyDay,
	}
}
