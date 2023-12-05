package code

import "github.com/starter-go/application"

// BackendStarter ...
type BackendStarter struct {

	//starter:component

	// _as func(afs.FS) //starter:as("#")

}

func (inst *BackendStarter) _impl() application.Lifecycle { return inst }

// Life ...
func (inst *BackendStarter) Life() *application.Life {
	return &application.Life{}
}
