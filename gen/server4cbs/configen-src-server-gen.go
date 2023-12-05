package server4cbs
import (
    pe3f72c153 "github.com/starter-go/cbs/src/server/golang/code"
     "github.com/starter-go/application"
)

// type pe3f72c153.ExampleController in package:github.com/starter-go/cbs/src/server/golang/code
//
// id:com-e3f72c1531da7ab2-code-ExampleController
// class:
// alias:
// scope:singleton
//
type pe3f72c1531_code_ExampleController struct {
}

func (inst* pe3f72c1531_code_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e3f72c1531da7ab2-code-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe3f72c1531_code_ExampleController) new() any {
    return &pe3f72c153.ExampleController{}
}

func (inst* pe3f72c1531_code_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe3f72c153.ExampleController)
	nop(ie, com)

	


    return nil
}


