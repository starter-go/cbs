package main

import (
	"os"

	"github.com/starter-go/cbs/modules/cbs"
	"github.com/starter-go/starter"
)

func main() {
	mod := cbs.Module()
	i := starter.Init(os.Args)
	i.MainModule(mod)
	i.WithPanic(true).Run()
}
