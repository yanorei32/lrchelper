package edtlegacy

import (
	"fmt"
	"time"

	"github.com/yanorei32/lrchelper/def"
	"github.com/yanorei32/lrchelper/lrc"
)

func GenerateBase(length time.Duration, timing def.Timing) (s string) {
	tick := 0

	for {
		pos := time.Minute * time.Duration(tick) / time.Duration(timing.Bpm)
		pos += timing.GlobalOffset

		if length < pos {
			return
		}

		s += fmt.Sprintf(
			"[%s] %3d: %2d/%2d \n",
			lrc.FormatPosition(pos),
			tick / timing.Bpb + 1,
			tick % timing.Bpb + 1,
			timing.Bpb,
		)

		tick += 1
	}
}


