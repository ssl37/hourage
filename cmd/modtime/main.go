package main

import (
	"fmt"
	"log"
	"os"
	"sslutz/hourage/utils"
	"time"
)

func main() {
	entries := utils.GetHrs(utils.Day)
	if len(entries) < 1 {
		log.Fatal("No beginning hours found. Please run 'gmb' to proceed")
	}

	timeOff, err := time.ParseDuration(os.Args[1] + "h")
	utils.ErrorAbort(err)
	timeOffEntry, _ := utils.NewHourage(utils.ModTime, time.Now(), timeOff)
	utils.AppendHrsFile(timeOffEntry)

	fmt.Printf("Modified by %s\n", utils.FmtHourage(timeOffEntry))

	if len(os.Args) > 2 {
		comment := fmt.Sprintf("%s (%.2f)", os.Args[2], timeOff.Hours())
		fmt.Printf("Adding Comment: %s\n", comment)
		utils.AppendCommentFile(comment)
	}
}
