package lrc

import (
	"time"
	"fmt"
)

func FormatPosition(d time.Duration) string {
	d = d.Round(time.Millisecond * 10)
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	cs := d / (time.Millisecond * 10)

	return fmt.Sprintf("%02d:%02d.%02d", m, s, cs)
}

func ParsePosition(s string) (time.Duration, error) {
	return time.ParseDuration(
		fmt.Sprintf("%sm%ss", s[0:2], s[3:]),
	)
}

