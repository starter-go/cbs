package cbs

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cbs"
	"github.com/starter-go/cbs/gen/agent4cbs"
	"github.com/starter-go/cbs/gen/main4cbs"
	"github.com/starter-go/cbs/gen/server4cbs"
	"github.com/starter-go/cbs/gen/test4cbs"
)

// Module ...
func Module() application.Module {
	mb := cbs.NewMainModule()
	mb.Components(main4cbs.ExportComponents)
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	parent := Module()

	mb := cbs.NewMainModule()
	mb.Components(test4cbs.ExportComponents)
	mb.Depend(parent)
	return mb.Create()
}

// ModuleForServer ...
func ModuleForServer() application.Module {

	// parent := Module()

	mb := cbs.NewServerModule()
	mb.Components(server4cbs.ExportComponents)
	// mb.Depend(parent)
	return mb.Create()
}

// ModuleForAgent ...
func ModuleForAgent() application.Module {

	//parent := Module()

	mb := cbs.NewAgentModule()
	mb.Components(agent4cbs.ExportComponents)
	// mb.Depend(parent)
	return mb.Create()
}
