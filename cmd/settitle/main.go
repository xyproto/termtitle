package main

import (
	"fmt"

	"github.com/xyproto/termtitle"
)

func main() {
	title := "TESTING 1 2 3"
	s, err := termtitle.TitleString(title)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", s)
	termtitle.SetTitle(title)
}
