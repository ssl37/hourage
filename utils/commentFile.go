package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func AppendCommentFile(comment string) {
	today := time.Now()
	fileName := os.Getenv("COMMENTFILE")
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0660)
	ErrorAbort(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("%s %s\n", today.Format(time.DateOnly), comment))
	writer.Flush()
}
