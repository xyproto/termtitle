package main

import (
	"fmt"
	"os"

	"github.com/xyproto/termtitle"
)

func main() {
	title := "TESTING 1 2 3"
	s, err := termtitle.TitleString(title)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Printing: %q\n", s)
	termtitle.Set(title)
}
