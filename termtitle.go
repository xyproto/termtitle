package termtitle

import (
	"errors"
	"fmt"
	"os"
)

// TitleString returns a string that can be used for setting the title
// for the currently running terminal emulator, or an error if not supported.
func TitleString(title string) (string, error) {
	var formatString string
	if len(os.Getenv("KONSOLE_VERSION")) == 0 { // konsole?
		formatString = "\033]30;%s\007"
	} else if len(os.Getenv("GNOME_TERMINAL_SERVICE")) == 0 { // gnome-terminal?
		formatString = "\033]0;%s\a"
	}
	if len(formatString) == 0 {
		return "", errors.New("found no supporter terminal emulator")
	}
	return fmt.Sprintf(formatString, title), nil
}

// SetTitle tries to set the title of the currently running terminal emulator.
// An error is returned if no supported terminal emulator is found.
func SetTitle(title string) error {
	s, err := TitleString(title)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
