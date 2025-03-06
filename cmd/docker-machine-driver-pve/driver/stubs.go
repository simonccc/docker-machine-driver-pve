package driver

import (
	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/mcnflag"
	"github.com/rancher/machine/libmachine/state"
)

// Create implements drivers.Driver.
func (d *Driver) Create() error {
	panic("unimplemented")
}

// GetCreateFlags implements drivers.Driver.
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	panic("unimplemented")
}

// GetIP implements drivers.Driver.
func (d *Driver) GetIP() (string, error) {
	panic("unimplemented")
}

// GetMachineName implements drivers.Driver.
func (d *Driver) GetMachineName() string {
	panic("unimplemented")
}

// GetSSHHostname implements drivers.Driver.
func (d *Driver) GetSSHHostname() (string, error) {
	panic("unimplemented")
}

// GetSSHKeyPath implements drivers.Driver.
func (d *Driver) GetSSHKeyPath() string {
	panic("unimplemented")
}

// GetSSHPort implements drivers.Driver.
func (d *Driver) GetSSHPort() (int, error) {
	panic("unimplemented")
}

// GetSSHUsername implements drivers.Driver.
func (d *Driver) GetSSHUsername() string {
	panic("unimplemented")
}

// GetState implements drivers.Driver.
func (d *Driver) GetState() (state.State, error) {
	panic("unimplemented")
}

// GetURL implements drivers.Driver.
func (d *Driver) GetURL() (string, error) {
	panic("unimplemented")
}

// Kill implements drivers.Driver.
func (d *Driver) Kill() error {
	panic("unimplemented")
}

// PreCreateCheck implements drivers.Driver.
func (d *Driver) PreCreateCheck() error {
	panic("unimplemented")
}

// Remove implements drivers.Driver.
func (d *Driver) Remove() error {
	panic("unimplemented")
}

// Restart implements drivers.Driver.
func (d *Driver) Restart() error {
	panic("unimplemented")
}

// SetConfigFromFlags implements drivers.Driver.
func (d *Driver) SetConfigFromFlags(_ drivers.DriverOptions) error {
	panic("unimplemented")
}

// Start implements drivers.Driver.
func (d *Driver) Start() error {
	panic("unimplemented")
}

// Stop implements drivers.Driver.
func (d *Driver) Stop() error {
	panic("unimplemented")
}
