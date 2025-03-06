package driver

import (
	"github.com/rancher/machine/libmachine/drivers"
)

var _ drivers.Driver = (*Driver)(nil)

// Driver is the implementation of drivers.Driver interface.
type Driver struct {
	*drivers.BaseDriver
}

// Creates a new driver.
func NewDriver(machineName, storePath string) *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: machineName,
			StorePath:   storePath,
		},
	}
}

// DriverName implements drivers.Driver.
func (d *Driver) DriverName() string {
	return "pve"
}
