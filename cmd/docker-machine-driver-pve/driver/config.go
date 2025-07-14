package driver

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/mcnflag"
)

// Available flags.
const (
	flagURL              = "pve-url"
	flagInsecureTLS      = "pve-insecure-tls"
	flagTokenID          = "pve-token-id" //nolint:gosec // False-positive
	flagTokenSecret      = "pve-token-secret"
	flagResourcePool     = "pve-resource-pool"
	flagTemplateID       = "pve-template"
	flagISODevice        = "pve-iso-device"
	flagNetworkInterface = "pve-network-interface"
	flagSSHUser          = "pve-ssh-user"
	flagSSHPort          = "pve-ssh-port"
	flagProcessorSockets = "pve-processor-sockets"
	flagProcessorCores   = "pve-processor-cores"
	flagMemory           = "pve-memory"
	flagMemoryBalloon    = "pve-memory-balloon"
)

// Default values for flags.
const (
	defaultSSHUser = "service"
	defaultSSHPort = 22
)

// Driver's configuration.
type config struct {
	// Proxmox VE URL (e.g. 'https://<PROXMOX VE ADDRESS>:8006').
	URL string

	// Disables Proxmox VE TLS certificate verification.
	InsecureTLS bool

	// Proxmox VE API Token ID (including username and realm, e.g. 'root@pam!rancher').
	TokenID string

	// Proxmox VE API Token secret.
	TokenSecret string

	// Proxmox VE Resource Pool name.
	ResourcePoolName string

	// ID of the Proxmox VE template.
	TemplateID int

	// Bus/Device of the CD/DVD Drive to mount cloud-init ISO to (e.g. 'scsi1').
	ISODeviceName string

	// Bus/Device of the network interface to read machine's IP address from (e.g. 'net0').
	NetworkInterfaceName string

	// Number of processor sockets. If set to 0 (default), this configuration is skipped.
	ProcessorSockets int

	// Number of processor cores. If set to 0 (default), this configuration is skipped.
	ProcessorCores int

	// Amount of memory in MiB. If set to 0 (default), this configuration is skipped.
	Memory int

	// Minimum amount of memory in MiB. If set to 0 (default), defaults to "pve-memory" or skips if it's also 0.
	MemoryBalloon int
}

// GetCreateFlags implements drivers.Driver.
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			Name:   flagURL,
			EnvVar: flagEnvVarFromFlagName(flagURL),
			Usage:  "Proxmox VE URL (e.g. 'https://<PROXMOX VE ADDRESS>:8006')",
		},
		mcnflag.BoolFlag{
			Name:   flagInsecureTLS,
			EnvVar: flagEnvVarFromFlagName(flagInsecureTLS),
			Usage:  "Disables Proxmox VE TLS certificate verification",
		},
		mcnflag.StringFlag{
			Name:   flagTokenID,
			EnvVar: flagEnvVarFromFlagName(flagTokenID),
			Usage:  "Proxmox VE API Token ID (including username and realm, e.g. 'root@pam!rancher')",
		},
		mcnflag.StringFlag{
			Name:   flagTokenSecret,
			EnvVar: flagEnvVarFromFlagName(flagTokenSecret),
			Usage:  "Proxmox VE API Token secret",
		},
		mcnflag.StringFlag{
			Name:   flagResourcePool,
			EnvVar: flagEnvVarFromFlagName(flagResourcePool),
			Usage:  "Proxmox VE Resource Pool name",
		},
		mcnflag.IntFlag{
			Name:   flagTemplateID,
			EnvVar: flagEnvVarFromFlagName(flagTemplateID),
			Usage:  "ID of the Proxmox VE template",
		},
		mcnflag.StringFlag{
			Name:   flagISODevice,
			EnvVar: flagEnvVarFromFlagName(flagISODevice),
			Usage:  "Bus/Device of the CD/DVD Drive to mount cloud-init ISO to (e.g. 'scsi1')",
		},
		mcnflag.StringFlag{
			Name:   flagNetworkInterface,
			EnvVar: flagEnvVarFromFlagName(flagNetworkInterface),
			Usage:  "Bus/Device of the network interface to read machine's IP address from (e.g. 'net0')",
		},
		mcnflag.StringFlag{
			Name:   flagSSHUser,
			EnvVar: flagEnvVarFromFlagName(flagSSHUser),
			Usage:  fmt.Sprintf("Username for the SSH user that will be created via cloud-init, defaults to '%s'", defaultSSHUser),
		},
		mcnflag.IntFlag{
			Name:   flagSSHPort,
			EnvVar: flagEnvVarFromFlagName(flagSSHPort),
			Usage:  fmt.Sprintf("Port to use when connecting to the machine via SSH, defaults to '%d'", defaultSSHPort),
		},
		mcnflag.IntFlag{
			Name:   flagProcessorSockets,
			EnvVar: flagEnvVarFromFlagName(flagProcessorSockets),
			Usage:  "Number of processor sockets. If set to 0 (default), this configuration is skipped.",
		},
		mcnflag.IntFlag{
			Name:   flagProcessorCores,
			EnvVar: flagEnvVarFromFlagName(flagProcessorCores),
			Usage:  "Number of processor cores. If set to 0 (default), this configuration is skipped.",
		},
		mcnflag.IntFlag{
			Name:   flagMemory,
			EnvVar: flagEnvVarFromFlagName(flagMemory),
			Usage:  "Amount of memory in MiB. If set to 0 (default), this configuration is skipped.",
		},
		mcnflag.IntFlag{
			Name:   flagMemoryBalloon,
			EnvVar: flagEnvVarFromFlagName(flagMemoryBalloon),
			Usage:  fmt.Sprintf("Minimum amount of memory in MiB. If set to 0 (default), defaults to value of '--%s' or skips if it's also set to 0.", flagMemory),
		},
	}
}

// SetConfigFromFlags implements drivers.Driver.
//
//nolint:cyclop
func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	d.URL = opts.String(flagURL)
	if d.URL == "" {
		return fmt.Errorf("flag '--%s' is required", flagURL)
	}

	if _, err := url.Parse(d.URL); err != nil {
		return fmt.Errorf("failed to parse Proxmox VE URL (flag '--%s'): %w", flagURL, err)
	}

	d.InsecureTLS = opts.Bool(flagInsecureTLS)

	d.TokenID = opts.String(flagTokenID)
	if d.TokenID == "" {
		return fmt.Errorf("flag '--%s' is required", flagTokenID)
	}

	d.TokenSecret = opts.String(flagTokenSecret)
	if d.TokenSecret == "" {
		return fmt.Errorf("flag '--%s' is required", flagTokenSecret)
	}

	d.ResourcePoolName = opts.String(flagResourcePool)
	if d.ResourcePoolName == "" {
		return fmt.Errorf("flag '--%s' is required", flagResourcePool)
	}

	d.TemplateID = opts.Int(flagTemplateID)
	if d.TemplateID <= 0 {
		return fmt.Errorf("flag '--%s' is required and must be >= 0", flagTemplateID)
	}

	d.ISODeviceName = strings.ToLower(opts.String(flagISODevice))
	if d.ISODeviceName == "" {
		return fmt.Errorf("flag '--%s' is required", flagISODevice)
	}

	d.NetworkInterfaceName = opts.String(flagNetworkInterface)
	if d.NetworkInterfaceName == "" {
		return fmt.Errorf("flag '--%s' is required", flagNetworkInterface)
	}

	d.SSHUser = opts.String(flagSSHUser)
	if d.SSHUser == "" {
		d.SSHUser = defaultSSHUser
	}

	d.SSHPort = opts.Int(flagSSHPort)
	if d.SSHPort == 0 {
		d.SSHPort = defaultSSHPort
	} else if d.SSHPort < 0 {
		return fmt.Errorf("flag '--%s' must be > 0", flagSSHPort)
	}

	d.ProcessorSockets = opts.Int(flagProcessorSockets)
	if d.ProcessorSockets != 0 && d.ProcessorSockets < 1 {
		return fmt.Errorf("flag '--%s' must be >= 1; set to 0 to disable", flagProcessorSockets)
	}

	d.ProcessorCores = opts.Int(flagProcessorCores)
	if d.ProcessorCores != 0 && d.ProcessorCores < 1 {
		return fmt.Errorf("flag '--%s' must be >= 1; set to 0 to disable", flagProcessorCores)
	}

	d.Memory = opts.Int(flagMemory)
	if d.Memory != 0 && d.Memory < 1 {
		return fmt.Errorf("flag '--%s' must be >= 1; set to 0 to disable", flagMemory)
	}

	d.MemoryBalloon = opts.Int(flagMemoryBalloon)
	if d.MemoryBalloon == 0 {
		d.MemoryBalloon = d.Memory
	} else if d.MemoryBalloon < 1 {
		return fmt.Errorf("flag '--%s' must be >= 1; set to 0 to disable", flagMemoryBalloon)
	}

	if d.MemoryBalloon > d.Memory {
		return fmt.Errorf("flag '--%s' must be <= than flag '--%s'", flagMemoryBalloon, flagMemory)
	}

	return nil
}

// Creates flag's EnvVar from it's name.
func flagEnvVarFromFlagName(name string) string {
	return strings.ToUpper(
		strings.ReplaceAll(
			name,
			"-",
			"_",
		),
	)
}
