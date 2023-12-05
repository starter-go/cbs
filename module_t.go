package cbs

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/cbs"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

////////////////////////////////////////////////////////////////////////////////

const (
	theAgentModuleResPath  = "src/agent/resources"
	theMainModuleResPath   = "src/main/resources"
	theServerModuleResPath = "src/server/resources"
	theTestModuleResPath   = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/server/resources"
var theServerModuleResFS embed.FS

//go:embed "src/agent/resources"
var theAgentModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule ...
func NewMainModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(starter.Module())
	return mb
}

// NewTestModule ...
func NewTestModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Depend(starter.Module())
	return mb
}

// NewServerModule ...
func NewServerModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theServerModuleResFS, theServerModuleResPath)

	mb.Depend(starter.Module())
	mb.Depend(libgin.Module())
	return mb
}

// NewAgentModule ...
func NewAgentModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#agent")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theAgentModuleResFS, theAgentModuleResPath)

	mb.Depend(starter.Module())
	return mb
}
