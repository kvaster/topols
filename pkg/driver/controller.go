package driver

import (
	internalDriver "github.com/kvaster/topols/internal/driver"
)

// NewControllerServer is an externally consumable wrapper.
// It allows starting a new controller server even without access to the package internals.
var NewControllerServer = internalDriver.NewControllerServer
