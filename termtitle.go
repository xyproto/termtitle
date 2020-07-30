package termtitle

import (
	"errors"
	"fmt"
	"os"
)

// hasE checks if the given environment variable name is set
func hasE(envVar string) bool {
	_, ok := os.LookupEnv(envVar)
	return ok
}

// TitleString returns a string that can be used for setting the title
// for the currently running terminal emulator, or an error if not supported.
func TitleString(title string) (string, error) {
	var formatString string
	if hasE("KONSOLE_VERSION") { // konsole?
		formatString = "\033]30;%s\007"
	} else if hasE("GNOME_TERMINAL_SERVICE") { // gnome-terminal?
		formatString = "\033]0;%s\a"
	}
	if len(formatString) == 0 {
		return "", errors.New("found no supported terminal emulator")
	}
	return fmt.Sprintf(formatString, title), nil
}

// SetTitle tries to set the title of the currently running terminal emulator.
// An error is returned if no supported terminal emulator is found.
func Set(title string) error {
	s, err := TitleString(title)
	if err != nil {
		return err
	}
	fmt.Print(s)
	return nil
}
