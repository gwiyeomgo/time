package time

import (
	"errors"
	"time"
)

const (
	DateLayout6  = "060102"
	DateLayout8  = "20060102"
	DateLayout10 = "2006-01-02"
	DateLayout12 = "200601021504"
	DateLayout14 = "20060102150405"
	DateLayout19 = "2006-01-02 15:04:05"
	AsiaSeoul    = "Asia/Seoul"
	Local        = "Local"
	Month        = "01"
	Year         = "2006"
)

/*
mocking 을 위한 패키지 수준 시간 설정 함수 추가
https://labs.yulrizka.com/en/stubbing-time-dot-now-in-golang/
*/
type nowFuncT func() time.Time

var nowFunc nowFuncT

/*
SetNowFunc 함수로 시간을 설정한
패키지 내부 현재 시간
*/
func now() time.Time {
	if nowFunc == nil {
		return time.Now().UTC()
	}
	return nowFunc()
}
func SetNowTime(fn nowFuncT) {
	nowFunc = fn
}

type Time struct {
	location string
	current  time.Time
}

/*
세계적으로 표준화된 시간대 UTC
이 함수가 항상 UTC 기준의 현재 시간을 반환
*/
func (t Time) CurrentTime() time.Time {
	if t.location == "" {
		t.location = Local
	}
	loc, err := time.LoadLocation(t.location)
	if err != nil {
		panic(err.Error())
	}

	return time.Now().UTC().In(loc)
}

var errorEmptyTime = errors.New("빈 시간입니다.")

/*
Go 언어에서 time.Date 함수는 월을 1부터 12까지의 정수로 표현
previousMonth가 11인 경우 (12월은 11로 표현됨), 1을 더해서 12로
*/
func (t Time) PreviousMonth() string {
	if t.current.IsZero() {
		panic(errorEmptyTime.Error())
	}
	currentMonth := t.current.Month()

	previousMonth := currentMonth - 1
	if currentMonth == time.January {
		previousMonth = time.December
	}
	return time.Date(0, previousMonth+1, 0, 0, 0, 0, 0, time.UTC).Format(Month)
}

func (t Time) PreviousYear() string {
	if t.current.IsZero() {
		panic(errorEmptyTime.Error())
	}
	currentYear := t.current.Year()
	previousYear := currentYear - 1
	return time.Date(previousYear+1, 0, 0, 0, 0, 0, 0, time.UTC).Format(Year)
}
