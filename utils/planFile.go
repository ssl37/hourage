package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func WritePlanFile(planTime time.Time) {
	fileName := os.Getenv("PLANFILE")
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0660)
	ErrorAbort(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(planTime.Format(time.DateTime))
	writer.Flush()
}

func ReadPlanFile() (TimeOut time.Time) {
	loc, _ := time.LoadLocation("America/Denver")
	fileName := os.Getenv("PLANFILE")
	file, errOpen := os.Open(fileName)
	ErrorAbort(errOpen)

	buffer := make([]byte, 19)
	_, errRead := file.Read(buffer)
	ErrorAbort(errRead)
	defer file.Close()

	out := string(buffer[:])
	timeOut, errParse := time.ParseInLocation(time.DateTime, out, loc)
	ErrorAbort(errParse)
	fmt.Printf("Planned Departure: %s\n", out)
	return timeOut
}
