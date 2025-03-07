package utils

import (
	"log"
)

func ErrorAbort(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
