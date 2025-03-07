package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func getHrsEntry(file *os.File) (hourage *Hourage) {
	loc, _ := time.LoadLocation("America/Denver")
	modification, _ := time.ParseDuration("0h")
	buffer := make([]byte, 23)

	_, err := file.Read(buffer)
	ErrorAbort(err)

	entry := string(buffer[:])
	prefix := entry[:2]
	if prefix == ModTime {
		timeEntry, err := time.ParseInLocation(time.DateOnly, entry[4:14], loc)
		ErrorAbort(err)
		modification, _ := time.ParseDuration(entry[15:])
		hourage, err = NewHourage(prefix, timeEntry, modification)
		ErrorAbort(err)
		return hourage
	}
	timeEntry, err := time.ParseInLocation(time.DateTime, entry[4:], loc)
	ErrorAbort(err)
	hourage, err = NewHourage(prefix, timeEntry, modification)
	ErrorAbort(err)
	return hourage
}

func ReadHrs(prefix string, length time.Duration) (entries []Hourage) {
	now := time.Now()
	start := now.Add(-1 * length)
	return ReadHrsFile(prefix, start, now)
}

func ReadHrsFile(prefix string, start time.Time, end time.Time) (entries []Hourage) {
	fileName := os.Getenv("HRSFILE")
	file, errOpen := os.Open(fileName)
	ErrorAbort(errOpen)
	defer file.Close()

	daysToGet := int((end.Sub(start)).Hours()/24) + 1
	if daysToGet <= 0 {
		errMsg := fmt.Sprintf("No days in range %s to %s", start.Format(time.DateOnly), end.Format(time.DateOnly))
		ErrorAbort(errors.New(errMsg))
	}
	entries = make([]Hourage, 4*daysToGet)
	n := 0

	file.Seek(-24, os.SEEK_END)
	hourage := getHrsEntry(file)
	for hourage.ts.After(end) {
		file.Seek(-47, os.SEEK_CUR)
		hourage = getHrsEntry(file)
	}
	for hourage.ts.After(start) && hourage.ts.Before(end) {
		if hourage.prefix == prefix || prefix == ALL {
			n += copy(entries[n:], []Hourage{*hourage})
		}
		file.Seek(-47, os.SEEK_CUR)
		hourage = getHrsEntry(file)
	}

	return entries[:n]
}
