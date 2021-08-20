package main

import "net"

// InitNetworkConfig is a config section for network specific options
type InitNetworkConfig struct {
	Interfaces []net.HardwareAddr `yaml:",omitempty"` // list of active interfaces to use

	Dhcp bool `yaml:",omitempty"`

	IP         string `yaml:",omitempty"`            // e.g. 10.0.2.15/24
	Gateway    string `yaml:",omitempty"`            // e.g. 10.0.2.255
	DNSServers string `yaml:"dns_servers,omitempty"` // comma-separated list of ips, e.g. 10.0.1.1,8.8.8.8
}

// VirtualConsole is a config section for console specific options
type VirtualConsole struct {
	KeymapFile      string `yaml:",omitempty"`
	Utf             bool   `yaml:",omitempty"`
	FontFile        string `yaml:",omitempty"`
	FontMapFile     string `yaml:",omitempty"`
	FontUnicodeFile string `yaml:",omitempty"`
}

// InitConfig represents top-level section for config passed from generator to init
type InitConfig struct {
	Network                *InitNetworkConfig  `yaml:",omitempty"`
	ModuleDependencies     map[string][]string `yaml:",omitempty"`
	ModulePostDependencies map[string][]string `yaml:",omitempty"`
	ModulesForceLoad       []string            `yaml:",omitempty"`
	ModprobeOptions        map[string]string   `yaml:",omitempty"`
	BuiltinModules         set                 `yaml:",omitempty"`
	Kernel                 string              `yaml:",omitempty"` // kernel version this image was built for
	MountTimeout           int                 `yaml:",omitempty"` // mount timeout in seconds
	VirtualConsole         *VirtualConsole     `yaml:",omitempty"`
	EnableLVM              bool                `yaml:",omitempty"`
	EnableMdraid           bool                `yaml:",omitempty"`
}

const initConfigPath = "/etc/booster.init.yaml"
