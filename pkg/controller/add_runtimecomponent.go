package controller

import (
	"github.com/application-stacks/runtime-component-operator/pkg/controller/runtimecomponent"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, runtimecomponent.Add)
}
