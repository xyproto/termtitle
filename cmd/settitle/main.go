package main

import (
	"os"
	"strings"
	"time"

	"github.com/xyproto/termtitle"
)

func main() {
	title := time.Now().String()
	if len(os.Args) > 1 {
		title = strings.Join(os.Args[1:], "")
	}
	termtitle.Set(title)
}
