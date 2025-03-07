package utils

import (
	"fmt"
	"time"
	"strings"
)

func parsePeriod(period string) (start time.Time, end time.Time) {
	end = time.Now()
	loc, _ := time.LoadLocation("America/Denver")
	dashat := strings.Index(period, "-")
	if dashat > -1 {
		start, _ = parsePeriod(period[:dashat])
		end, _ = parsePeriod(period[dashat+1:])
	} else {
		var err error
		start, err = time.ParseInLocation("Jan06", period, loc)
		if err != nil {
			appendYear := fmt.Sprintf("%s%d", period, end.Year())
			start, err = time.ParseInLocation("Jan2006", appendYear, loc)
			ErrorAbort(err)
		}
		end = start.AddDate(0,1,0)
		start = start.AddDate(0,0,-1)
	}

	return start, end
}

func GetHrs(period string) ([]Hourage) {

	end := time.Now()
	var start time.Time
	if period == Day {
		start = time.Date(end.Year(), end.Month(), end.Day(), 0,0,0,0,time.Local)
	} else if period == Month {
		start = time.Date(end.Year(), end.Month(),1,0,0,0,0,time.Local)
	} else {
		start, end = parsePeriod(period)
	}

	hrsInEntries := ReadHrsFile(TimeIn, start, end)
	hrsOutEntries := ReadHrsFile(TimeOut, start, end)
	modEntries := ReadHrsFile(ModTime, start, end)

	n := 0
	hrs:= make([]Hourage, len(hrsInEntries))

	for _, in := hrsInEntries[:] {
		startTime := in.GetTS()
		noneFound := true
		for _, mods := range modEntries[:] {
			if mods.GetTS().Format(time.DateOnly) == in.GetTS().Format(time.Date)
		}
		for _, out := hrsOutEntries[:] {
			if out.GetTS().Format(time.DateOnly) == in.GetTS().Format(time.DateOnly) && noneFound {
				diff := out.GetTS().Sub(startTime)
				hourage, err := NewHourage(ModTime, startTime, diff)
				ErrorAbort(err)
				n += copy(hrs[n:], []Hourage{*hourage})
				noneFound = false
			}
		}
		if noneFound {
			hourage, err := NewHourage(TimeIn, startTime, ZeroDuration)
			ErrorAbort(err)
			n += copy(hrs[n:], []Hourage{*hourage})
		}
	}

	return hrs
}
