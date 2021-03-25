package lyric

import (
	"fmt"
	"time"

	"github.com/yanorei32/lrchelper/lrc"
)

func (l Lyric) Format(offset time.Duration) (s string) {
	if l.Artist != "" {
		s += fmt.Sprintf("[ar:%s]\n", l.Artist)
	}

	if l.Album != "" {
		s += fmt.Sprintf("[al:%s]\n", l.Album)
	}

	if l.Title != "" {
		s += fmt.Sprintf("[ti:%s]\n", l.Title)
	}

	if l.Length != "" {
		s += fmt.Sprintf("[length:%s]\n", l.Length)
	}

	if l.Author != "" {
		s += fmt.Sprintf("[au:%s]\n", l.Author)
	}

	if l.Author != "" {
		s += fmt.Sprintf("[au:%s]\n", l.Author)
	}

	if l.By != "" {
		s += fmt.Sprintf("[by:%s]\n", l.By)
	}

	if l.Version != "" {
		s += fmt.Sprintf("[ve:%s]\n", l.Version)
	}

	s += "[offset: 0]\n\n"

	for _, line := range l.Lines {
		s += fmt.Sprintf(
			"[%s]%s\n",
			lrc.FormatPosition(line.Position+offset),
			line.Text,
		)
	}

	return
}
