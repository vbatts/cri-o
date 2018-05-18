// +build !windows

package lib

const (
	// DefaultHooksDirPath Default directory containing hooks config files
	DefaultHooksDirPath = "/usr/share/containers/oci/hooks.d"
	// OverrideHooksDirPath Directory where admin can override the default configuration
	OverrideHooksDirPath = "/etc/containers/oci/hooks.d"
)
