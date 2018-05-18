// +build windows

package lib

const (
	// DefaultHooksDirPath Default directory containing hooks config files
	DefaultHooksDirPath = "C:\\crio\\share\\oci\\hooks.d"
	// OverrideHooksDirPath Directory where admin can override the default configuration
	OverrideHooksDirPath = "C:\\crio\\etc\\oci\\hooks.d"
)
