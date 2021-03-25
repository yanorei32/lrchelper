package edtlegacy

import (
	"fmt"
	"time"

	"github.com/yanorei32/lrchelper/lrc"
)

func GenerateBase(length time.Duration, bpm int, bpb int, offset time.Duration) (s string) {
	tick := 0

	for {
		pos := time.Minute * time.Duration(tick) / time.Duration(bpm) + offset

		if length < pos {
			return
		}

		s += fmt.Sprintf("%s %3d: %2d/%2d \n", lrc.FormatPosition(pos), tick / bpb + 1, tick % bpb + 1, bpb)

		tick += 1
	}
}


