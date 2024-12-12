package main

import (
	"venera/internal/core"
)

func main() {
	var v float32 = 1.03 // version
	var s bool = false   // stable
	core.Start(v, s)
}
