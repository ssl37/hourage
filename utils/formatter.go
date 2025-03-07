package utils

import (
	"fmt"
	"time"
)

func FmtDuration(d time.Duration) string {
	return fmt.Sprintf("%.2f", d.Hours())
}

func FmtDate(h time.Time) string {
	return h.Format(time.DateOnly)
}
