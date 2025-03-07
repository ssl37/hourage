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
		log.Fatal("No beginning hours found. Please run 'gmb' to proceed")
	}

	if entries[0].GetPrefix() == utils.ModTime {
		fmt.Printf("You've already said goodbye: %s\n", utils.FmtDuration(entries[0].GetDuration()))
	}

	diff := time.Since(entries[0].GetTS())
	fmt.Println(utils.FmtDuration(diff))
}
