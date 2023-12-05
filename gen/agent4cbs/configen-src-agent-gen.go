package agent4cbs
import (
    p9ddeceb52 "github.com/starter-go/cbs/src/agent/golang/code"
     "github.com/starter-go/application"
)

// type p9ddeceb52.ExampleController in package:github.com/starter-go/cbs/src/agent/golang/code
//
// id:com-9ddeceb527910e7a-code-ExampleController
// class:
// alias:
// scope:singleton
//
type p9ddeceb527_code_ExampleController struct {
}

func (inst* p9ddeceb527_code_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9ddeceb527910e7a-code-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9ddeceb527_code_ExampleController) new() any {
    return &p9ddeceb52.ExampleController{}
}

func (inst* p9ddeceb527_code_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9ddeceb52.ExampleController)
	nop(ie, com)

	


    return nil
}


