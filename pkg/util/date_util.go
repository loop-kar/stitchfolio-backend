package util

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

var emptyDate time.Time = time.Time{}

const (
	DateTimeLayout_YYYYMMDD = time.DateOnly
	DateTimeLayout_DDMMYY   = "02-01-2006"
)

func FindMonthsBetweenTwoDates(date1, date2 *time.Time) int {
	if date1.After(*date2) {
		return -1
	}

	isDay1PastDay2 := date1.Day() > date2.Day()
	diff := date2.Sub(*date1)
	hours := diff.Hours()
	if isDay1PastDay2 {
		hours -= 24
	}
	months := math.Ceil(hours / 24 / 30)

	return int(months)

}

func TransformTimeToMMYYYY(time *time.Time) string {
	month := ConvertMToMM(int(time.Month()))
	return fmt.Sprintf("%s/%d", month, time.Year())
}

func GenerateDateTimeFromYYYYMM(dateString string, date int) (time.Time, error) {
	// must be of format month/year eg: 2021-10 , 2022-1 , 2022-01

	parts := strings.Split(dateString, constants.DATE_SEPERATOR)
	if len(parts) != 2 {
		return time.Time{}, errs.NewXError(errs.INVALID_REQUEST, "Invalid Date Format.Use format 'YEAR-MONTH", nil)
	}

	// day should be of double digits
	dateStr := ConvertMToMM(date)

	//format "2006-01-02" "YYYY-MM-DD"
	formattedDateString := fmt.Sprintf("%s-%s-%s", parts[0], parts[1], dateStr)

	return time.Parse(time.DateOnly, formattedDateString)

}

func ConvertMToMM(value int) string {
	if value > 9 {
		return fmt.Sprintf("%d", value)
	}

	prefix := "0"
	return fmt.Sprintf("%s%d", prefix, value)
}

func GenerateDateTimeFromString(dateString *string, layout ...string) (*time.Time, error) {
	if IsNilOrEmptyString(dateString) {
		return nil, nil
	}

	var format = time.DateOnly
	if len(layout) == 1 {
		format = layout[0]
	}

	time, err := time.Parse(format, *dateString)
	if err != nil {
		return nil, err
	}

	return &time, nil
}

func GetLocalTime() time.Time {
	indLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return time.Now()
	}

	now := time.Now().In(indLocation)
	return now
}

func DateTimeToStringOrDefault(time *time.Time, layout string) string {

	if time == nil {
		return ""
	}

	return time.Format(layout)
}

func CurrentDateString(layout ...string) string {

	var format = time.DateOnly
	if len(layout) == 1 {
		format = layout[0]
	}

	return time.Now().Format(format)
}

func IsEmptyDate(date *time.Time) bool {
	return emptyDate.Equal(*date)
}

func IsSamedate(date1, date2 *time.Time) bool {
	return date1 != nil && date2 != nil && date1.Format(time.DateOnly) == date2.Format(time.DateOnly)
}
