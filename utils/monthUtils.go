package utils

import (
	"time"
)

func ExpectedHourPerMonth() float64 {
	Today := time.Now()
	from := time.Date(Today.Year(), Today.Month(), 1, 0, 0, 0, 0, time.Local)
	to := time.Date(Today.Year(), Today.Month()+1, 1, 0, 0, 0, 0, time.Local)
	totalDays := float32(to.Sub(from) / (24.0 * time.Hour))
	weekDays := float32(from.Weekday()) - float32(to.Weekday())
	businessDays := int((totalDays*5 - weekDays*2) / 7)
	if to.Weekday() == time.Sunday {
		businessDays++
	}
	if from.Weekday() == time.Sunday {
		businessDays--
	}

	return float64(8.0 * float64(businessDays))
}
