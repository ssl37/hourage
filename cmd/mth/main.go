package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sslutz/hourage/utils"
	"strings"
	"time"
)

func main() {
	timeframe := utils.Month
	if len(os.Args) > 1 {
		timeframe = os.Args[1]
	}
	entries := utils.GetHrs(timeframe)
	if len(entries) < 1 {
		log.Fatal("No hours found. Possibly the beginning of the month? Or no work done?")
	}

	sum := 0.0
	slices.Reverse(entries)

	for _, entry := range entries {
		if entry.GetPrefix() == utils.ModTime {
			fmt.Printf("%s  %6s\n", utils.FmtDate(entry.GetTS()), utils.FmtDuration(entry.GetDuration()))
			sum += entry.GetDuration().Hours()
		} else {
			diff := time.Since(entry.GetTS())
			fmt.Printf("%s  %6s\n", utils.FmtDate(entry.GetTS()), utils.FmtDuration(diff))
			sum += diff.Hours()
		}
	}

	expected := utils.ExpectedHoursPerMonth()
	deficit := expected - sum

	fmt.Println(strings.Repeat("-", 18))
	fmt.Printf("Total:      %6.2f\n", sum)
	fmt.Printf("Expected:   %6.2f\n", expected)
	fmt.Printf("Deficit:    %6.2f\n", deficit)
}
