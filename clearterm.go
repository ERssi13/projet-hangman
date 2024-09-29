package projethangman

import (
	"fmt"
	"runtime"
)

func ClearTerminal() {
	if runtime.GOOS == "windows" {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
}