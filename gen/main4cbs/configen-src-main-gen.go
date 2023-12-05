package main4cbs
import (
    pa83488aff "github.com/starter-go/cbs/src/main/golang/code"
     "github.com/starter-go/application"
)

// type pa83488aff.AutoStarter in package:github.com/starter-go/cbs/src/main/golang/code
//
// id:com-a83488aff9ce1c4a-code-AutoStarter
// class:
// alias:
// scope:singleton
//
type pa83488aff9_code_AutoStarter struct {
}

func (inst* pa83488aff9_code_AutoStarter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a83488aff9ce1c4a-code-AutoStarter"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa83488aff9_code_AutoStarter) new() any {
    return &pa83488aff.AutoStarter{}
}

func (inst* pa83488aff9_code_AutoStarter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa83488aff.AutoStarter)
	nop(ie, com)

	


    return nil
}



// type pa83488aff.BackendStarter in package:github.com/starter-go/cbs/src/main/golang/code
//
// id:com-a83488aff9ce1c4a-code-BackendStarter
// class:
// alias:
// scope:singleton
//
type pa83488aff9_code_BackendStarter struct {
}

func (inst* pa83488aff9_code_BackendStarter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a83488aff9ce1c4a-code-BackendStarter"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa83488aff9_code_BackendStarter) new() any {
    return &pa83488aff.BackendStarter{}
}

func (inst* pa83488aff9_code_BackendStarter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa83488aff.BackendStarter)
	nop(ie, com)

	


    return nil
}



// type pa83488aff.FrontendStarter in package:github.com/starter-go/cbs/src/main/golang/code
//
// id:com-a83488aff9ce1c4a-code-FrontendStarter
// class:
// alias:
// scope:singleton
//
type pa83488aff9_code_FrontendStarter struct {
}

func (inst* pa83488aff9_code_FrontendStarter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a83488aff9ce1c4a-code-FrontendStarter"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa83488aff9_code_FrontendStarter) new() any {
    return &pa83488aff.FrontendStarter{}
}

func (inst* pa83488aff9_code_FrontendStarter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa83488aff.FrontendStarter)
	nop(ie, com)

	


    return nil
}


