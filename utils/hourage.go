package utils

import (
	"fmt"
	"time"
)

var ZeroDuration, _ = time.ParseDuration("0h")

type Hourage struct {
	prefix   string
	ts       time.Time
	duration time.Duration
}

func NewHourage(prefix string, ts time.Time, duration time.Duration) (*Hourage, error) {
	if prefix != ModTime && prefix != TimeIn && prefix != TimeOut {
		return nil, fmt.Errorf("Prefix %s is not allowed", prefix)
	}
	h := Hourage{prefix: prefix}
	h.ts = ts
	h.duration = 0
	if prefix == ModTime {
		h.duration = duration
		h.ts = ts.Add(120 * time.Minute)
	}
	return &h, nil
}

func FmtHourage(h *Hourage) string {
	if h.prefix == ModTime {
		return fmt.Sprintf("%s: %s %+07.2fh", h.prefix, h.ts.Format(time.DateOnly), h.duration.Hours())
	}
	return fmt.Sprintf("%s: %s", h.prefix, h.ts.Format(time.DateTime))
}

func (h *Hourage) GetDuration() time.Duration {
	if h.prefix != ModTime {
		return ZeroDuration
	}
	return h.duration
}

func (h *Hourage) GetTS() time.Time {
	return h.ts
}

func (h *Hourage) GetPrefix() string {
	return h.prefix
}
