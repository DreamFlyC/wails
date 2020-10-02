package goruntime

import (
	"fmt"

	"github.com/wailsapp/wails/v2/internal/crypto"
	"github.com/wailsapp/wails/v2/internal/servicebus"
)

// System defines all System related operations
type System interface {
	IsDarkMode() bool
}

// system exposes the System interface
type system struct {
	bus *servicebus.ServiceBus
}

// newSystem creates a new System struct
func newSystem(bus *servicebus.ServiceBus) System {
	return &system{
		bus: bus,
	}
}

// On pass through
func (r *system) IsDarkMode() bool {

	// Create unique system callback
	uniqueCallback := crypto.RandomID()

	// Subscribe to the respose channel
	responseTopic := "systemresponse:" + uniqueCallback
	systemResponseChannel, err := r.bus.Subscribe(responseTopic)
	if err != nil {
		fmt.Printf("ERROR: Cannot subscribe to bus topic: %+v\n", err.Error())
	}

	message := "system:isdarkmode:" + uniqueCallback
	r.bus.Publish(message, nil)

	// Wait for result
	var result *servicebus.Message = <-systemResponseChannel

	// Delete subscription to response topic
	r.bus.UnSubscribe(responseTopic)

	return result.Data().(bool)
}
