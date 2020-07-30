# termtitle

Change the title if the currently running terminal emulator supports it.

Currently only supports `konsole` and `gnome-terminal`. The `MustSet` function will try the same terminal codes as for `gnome-terminal`, though, which may also be supported by other terminal emulators.

Example use:

~~~go
package main

import (
    "github.com/xyproto/termtitle"
)

func main() {
    termtitle.Set("TESTING 1 2 3")
}
~~~

## Terminal codes

For `konsole` a working string seems to be:

    \033]0;TITLE\a

While for `gnome-terminal`, this one works:

    \033]30;TITLE\007

`TITLE` is the title to be set.

## General info

* Version: 1.1.0
* License: MIT
