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
