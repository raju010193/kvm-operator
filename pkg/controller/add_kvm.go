package controller

import (
	"github.com/raju140/kvm-operator/pkg/controller/kvm"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kvm.Add)
}
