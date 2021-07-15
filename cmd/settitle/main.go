package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/termtitle"
)

func main() {
	title := "TESTING 1 2 3"
	if len(os.Args) > 1 {
		title = strings.Join(os.Args[1:], "")
	}
	s, err := termtitle.TitleString(title)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Printing: %q\n", s)
	termtitle.Set(title)
}
