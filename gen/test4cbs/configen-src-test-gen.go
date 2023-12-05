package test4cbs
import (
    p7415730ba "github.com/starter-go/cbs/src/test/golang/code"
     "github.com/starter-go/application"
)

// type p7415730ba.ExampleController in package:github.com/starter-go/cbs/src/test/golang/code
//
// id:com-7415730baaa1941a-code-ExampleController
// class:
// alias:
// scope:singleton
//
type p7415730baa_code_ExampleController struct {
}

func (inst* p7415730baa_code_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7415730baaa1941a-code-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7415730baa_code_ExampleController) new() any {
    return &p7415730ba.ExampleController{}
}

func (inst* p7415730baa_code_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7415730ba.ExampleController)
	nop(ie, com)

	


    return nil
}


