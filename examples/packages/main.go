package main

import (
	"github.com/steevehook/mdc19/examples/packages/pkg1"
	// bad practice, avoid naming the package differently than the dir
	pkg_another "github.com/steevehook/mdc19/examples/packages/pkg2"
)

func main() {
	pkg1.F1()
	pkg_another.F2()
}
