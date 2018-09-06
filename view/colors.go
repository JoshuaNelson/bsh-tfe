package draw

import (
	"github.com/nsf/termbox-go"
)

func Color256(c int) termbox.Attribute {
	if c < 0 {
		c = 0
	} else if c > 255 {
		c = 255
	}

	return termbox.Attribute(c+1)
}
