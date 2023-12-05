package code

import "github.com/starter-go/application"

// AutoStarter ...
type AutoStarter struct {

	//starter:component

	// _as func(afs.FS) //starter:as("#")

}

func (inst *AutoStarter) _impl() application.Lifecycle { return inst }

// Life ...
func (inst *AutoStarter) Life() *application.Life {
	return &application.Life{}
}
