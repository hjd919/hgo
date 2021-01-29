package watermark

import (
	"github.com/issue9/watermark"
)

func New() {
	w, err := watermark.New("./path/to/watermark/file", 2, watermark.Center)
	if err != nil {
		panic(err)
	}

	err = w.MarkFile("./path/to/file")
}
