package doge

import (
	"fmt"
	"strings"

	"github.com/davidjoliver86/shiba"
)

func Bark() string {
	return "BARK"
}

func Such(thing string) string {
	return "such " + strings.ToLower(thing)
}

func BigBark() string {
	return shiba.WhenGrownUp(Bark())
}

func BigSuch(thing string) string {
	return shiba.WhenGrownUp(Such(thing))
}

func From02() {
	fmt.Println("I'm from v0.2.0")
}

func From03() {
	fmt.Println("I'm from v0.3.0")
}
