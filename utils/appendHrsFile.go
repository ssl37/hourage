package utils

import (
	"os"
	"bufio"
)

func AppendHrsFile(hr *Hourage) {
	fileName := os.Getenv("HRSFILE")
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0660)
	ErrorAbort(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(FmtHourage(hr)+"\n")
	writer.Flush()
}