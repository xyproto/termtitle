# termtitle

Change the title if the currently running terminal emulator supports it.

## Currently supported terminal emulators

* `konsole`
* `alacritty`
* `gnome-terminal`

## Example use

~~~
package main

import (
    "github.com/xyproto/termtitle"
)

func main() {
    termtitle.Set("TESTING 1 2 3")
}
~~~

## General info

* Version: 1.2.0
* License: MIT
