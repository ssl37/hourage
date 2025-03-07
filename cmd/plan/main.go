package main

import (
	"fmt"
	"log"
	"os"
	"sslutz/hourage/utils"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("American/Denver")
	anHour, _ := time.ParseDuration("1h")
	dur := utils.ZeroDuration

	entries := utils.GetHrs(utils.Day)
	if len(entries) < 1 {
		log.Fatal("No beginning hours found. Please run 'gmb' to proceed")
	}

	if entries[0].GetPrefix() == utils.ModTime {
		fmt.Printf("You've already said goodbye: %s\n", utils.FmtDuration(entries[0].GetDuration()))
	}

	diff := time.Since(entries[0].GetTS())

	if len(os.Args) > 1 {
		plan := os.Args[1]
		if string(plan[0]) == "+" {
			dur, _ = time.ParseDuration(string(plan[1:]) + "h")
		} else {
			plan = fmt.Sprintf("%s %s:00", time.Now().Format(time.DateOnly), plan[0:5])
			t, eeee := time.ParseInLocation(time.DateTime, plan, loc)
			if eeee != nil {
				fmt.Println(eeee)
				log.Fatal("Error while parsing your end time. Aborting")
			}
			dur = time.Until(t) + diff
		}
	} else {
		timeout := utils.ReadPlanFile()
		dur = timeout.Sub(entries[0].GetTS())
		if dur < 0 {
			log.Fatal("Last planned departure was in the past. Aborting")
		}
	}

	timeleft := dur - diff
	timeToLeave := time.Now().Add(timeleft)
	utils.WritePlanFile(timeToLeave)

	for timeleft > anHour {
		fmt.Printf("%s hours remaining\n", utils.FmtDuration(timeleft))
		time.Sleep(anHour)
		timeleft -= anHour
	}

	fmt.Printf("%s hours remaining\n", utils.FmtDuration(timeleft))
	tenMinutesToGo := timeleft - 10*time.Minute
	if tenMinutesToGo < 0 {
		fmt.Printf("%.0f minutes remaining\n", timeleft.Minutes())
		time.Sleep(timeleft)
	} else {
		time.Sleep(tenMinutesToGo)
		fmt.Printf("10 minutes remaining...\n")
		time.Sleep(10 * time.Minute)
	}

	fmt.Printf("Head home already...\n")
}
