package edtbeat

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/yanorei32/lrchelper/def"
	"github.com/yanorei32/lrchelper/lyric"
)

func Parse(f *bufio.Reader, timing def.Timing) ([]lyric.Line, error) {
	lines := []lyric.Line{}

	for {
		ls, err := f.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		lsa := strings.SplitN(ls, ":", 2)

		if len(lsa) < 2 {
			continue
		}

		lsaba := strings.SplitN(lsa[0], "/", 2)

		if len(lsaba) < 2 {
			return nil, errors.New("Failed to parse beatinfo")
		}

		bar, err := strconv.ParseInt(lsaba[0], 10, 32)
		if err != nil {
			return nil, errors.New("Failed to parse bar")
		}

		beat, err := strconv.ParseInt(lsaba[1], 10, 32)
		if err != nil {
			return nil, errors.New("Failed to parse beat")
		}

		l := lyric.Line{}

		l.Position = time.Minute * time.Duration(
			(int(bar)-1) * timing.Bpb + (int(beat)-1),
		) / time.Duration(timing.Bpm) + timing.GlobalOffset

		l.Text = strings.TrimSpace(lsa[1])

		lines = append(lines, l)
	}

	return lines, nil
}
