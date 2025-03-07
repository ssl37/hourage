package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sslutz/hourage/utils"
	"strings"
	"time"
)

func main() {
	sal := []string{"Beautiful", "Ancient One", "Hot Stuff",
		"My Dude", "Get Rekt", "Smarty Pants", "Go Code",
		"Code Cowboy", "Raging Bull", "Dances with Wolves",
		"Duke of Destruction", "Ravenkeeper", "Dread Pirate Roberts",
		"Number One", "Captain " + os.Getenv("USER")}

	login := os.Args[1]
	start := strings.LastIndex(login, ":0    ") + 10
	end := strings.Index(login, "still logged") - 2
	if end == -3 {
		log.Fatal("Could not find your current log in. Exiting...")
	}
	date := strings.TrimSpace(login[start:end])
	parsedDate, err := time.Parse(time.ANSIC, date)
	utils.ErrorAbort(err)

	entries := utils.GetHrs(utils.Day)
	if len(entries) > 0 {
		log.Fatal("You already logged in...")
	}

	timeInEntry, _ := utils.NewHourage(utils.TimeIn, parsedDate, utils.ZeroDuration)
	utils.AppendHrsFile(timeInEntry)

	fmt.Printf("(%s)  Good Morning, %s!\n", (parsedDate).Format(time.TimeOnly), sal[rand.Intn(len(sal))])
	fmt.Printf("You will hit 8 hours at %s\n", (parsedDate.Add(time.Hour * 8)).Format(time.TimeOnly))

}
