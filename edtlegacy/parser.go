package edtlegacy

import (
	"bufio"
	"io"
	"strings"

	"github.com/yanorei32/lrchelper/lrc"
	"github.com/yanorei32/lrchelper/lyric"
)

func Parse(f *bufio.Reader) ([]lyric.Line, error) {
	lines := []lyric.Line{}

	for {
		ls, err := f.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(ls) <= 23 {
			continue
		}

		lp, err := lrc.ParsePosition(ls[1:9])
		if err != nil {
			return nil, err
		}

		l := lyric.Line{}
		l.Position = lp
		l.Text = strings.TrimSpace(ls[21 : len(ls)-1])

		lines = append(lines, l)
	}

	return lines, nil
}
