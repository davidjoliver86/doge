package doge

import "strings"

func Bark() string {
	return "BARK"
}

func Such(thing string) string {
	return "such " + strings.ToLower(thing)
}
