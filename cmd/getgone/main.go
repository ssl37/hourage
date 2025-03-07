package main

import (
	"fmt"
	"log"
	"sslutz/hourage/utils"
	"time"
)

func main() {
	entries := utils.GetHrs(utils.Day)
	if len(entries) < 1 {
		log.Fatal("No hours returned, aborting... You'll need to stay")
	}

	if entries[0].GetPrefix() == utils.ModTime {
		fmt.Printf("You've already said goodbye: %s\n", utils.FmtDuration(entries[0].GetDuration()))
		return
	}

	diff := time.Since(entries[0].GetTS())
	fmt.Println(utils.FmtDuration(diff))
	timeOutEntry, _ := utils.NewHourage(utils.TimeOut, time.Now().Add(time.Second*10), utils.ZeroDuration)
	utils.AppendHrsFile(timeOutEntry)

	fmt.Println("So Long Suckers!")
}
