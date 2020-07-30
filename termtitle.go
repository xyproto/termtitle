package termtitle

import (
	"errors"
	"fmt"
	"os"
)

// SetTitle tries to set the title of the currently running terminal emulator.
// An error is returned if no supported terminal emulator is found.
func SetTitle(title string) error {
	var formatString string
	if os.Getenv("KONSOLE_VERSION") != "" { // konsole?
		formatString = "\033]30;%s\007"
	} else if os.Getenv("GNOME_TERMINAL_SERVICE") != "" { // gnome-terminal?
		formatString = "\033]0;%s\a"
	}
	if len(formatString) == 0 {
		return errors.New("found no supporter terminal emulator")
	}
	fmt.Printf(formatString, title)
	return nil
}
