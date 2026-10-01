package printcheck

import (
	"github.com/fatih/color"
)

func PrintCheck(name string, passed bool) {
	if passed {
		color.Green("OK: %s", name)
	} else {
		color.Red("FAIL: %s", name)
	}
}
