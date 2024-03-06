package time

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// time.Now() 함수는 실제로 호출될 때마다 시간이 업데이트되는 동적인 값이며,
// 코드에서 time.Now() 함수를 여러 번 호출하면 각 호출마다 약간씩 다른 결과를 얻을 수 있어서 초를 제외한 값으로 비교
func TestCurrentTime_현재시간_구함_성공(t *testing.T) {
	//given
	current := time.Now().UTC().In(time.Local)
	//when
	month := Time{}.CurrentTime()
	//then
	assert.Equal(t, month.Truncate(time.Second), current.Truncate(time.Second))
}
func TestCurrentTime_한국의_현재시간_구함_성공(t *testing.T) {
	location := AsiaSeoul
	month := Time{location: location}.CurrentTime()
	assert.Equal(t, month.Location().String(), location)
}
func TestPreviousMonth_지난달_숫자값을_문자로_반환_성공(t *testing.T) {
	//given
	current := time.Date(2023, 7, 31, 12, 00, 00, 00, time.UTC)
	//when
	month := Time{current: current}.PreviousMonth()
	//then
	assert.Equal(t, month, "06")
}

func TestPreviousMonth_지난달_숫자값을_문자로_반환시_빈시간일때_실패(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != errorEmptyTime.Error() {
				t.Errorf("예상된 패닉과 다른 메시지가 반환되었습니다. 예상된 메시지: %q, 실제 메시지: %q", errorEmptyTime.Error(), r)
			}
		} else {
			t.Error("panic이 발생하지 않았습니다.")
		}
	}()

	month := Time{}.PreviousMonth()
	fmt.Println(month)
}
func TestTime_PreviousYear_지난해_숫자값을_문자로_반환_성공(t *testing.T) {
	//given
	current := time.Date(2023, 7, 31, 12, 00, 00, 00, time.UTC)
	//when
	year := Time{current: current}.PreviousYear()
	//then
	assert.Equal(t, year, "2022")
}
func TestPreviousMonth_지난해_숫자값을_문자로_반환시_빈시간일때_실패(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != errorEmptyTime.Error() {
				t.Errorf("예상된 패닉과 다른 메시지가 반환되었습니다. 예상된 메시지: %q, 실제 메시지: %q", errorEmptyTime.Error(), r)
			}
		} else {
			t.Error("panic이 발생하지 않았습니다.")
		}
	}()

	year := Time{}.PreviousYear()
	fmt.Println(year)
}

/* 패키지 수준 시간 설정 함수 추가 */
// package main 에서 init 함수에 위치시켜 최초 time.now() 값을 초기화 합니다.
// 본 코드에서 시간을 할당하는 부분에 time.now() 코드를 쓰고
// 유닛 테스트시 SetNowTime 으로 원하는 시간으로 설정
// 유닛 테스트 종료시 SetNowTime 으로 다시 원래 시간으로 복귀시킵니다.
func Test_Mocking_시간을_설정하는_기능_테스트(t *testing.T) {

	utcTime := now()
	fmt.Println("기본 utc 시간:", utcTime)

	SetNowTime(func() time.Time {
		return time.Date(2023, time.November, 4, 11, 0, 0, 0, time.UTC)
	})
	defer func() {
		SetNowTime(func() time.Time {
			return time.Date(2023, time.November, 5, 15, 0, 0, 0, time.UTC)
		})
		restoreTime := now()
		fmt.Println("복구 utc 시간:", restoreTime)
	}()
	changeTime := now()
	fmt.Println("변경 utc 시간:", changeTime)
}

func TestCheckPastDate(t *testing.T) {
	result, _ := checkPastDate("20240208")
	result1, _ := checkPastDate("20240209")
	result2, _ := checkPastDate("20240207")

	assert.Equal(t, false, result)
	assert.Equal(t, false, result1)
	assert.Equal(t, true, result2)
}

// "20060102"
// "2006-01-02"
func Test_convertDateFormat(t *testing.T) {
	result, _ := convertDateFormat("20240208", DateLayout8, DateLayout10)
	assert.Equal(t, "2024-02-08", result)

	result2, err := convertDateFormat("20240208", DateLayout10, DateLayout8)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, "", result2)

	result3, err := convertDateFormat("2024-02-08", DateLayout10, DateLayout19)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, "2024-02-08 00:00:00", result3)
}
