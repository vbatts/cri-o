// +build windows

package lib

import "github.com/kubernetes-incubator/cri-o/oci"

// Defaults for linux/unix if none are specified
var (
	conmonPath            = "C:\\crio\\bin\\conmon"
	seccompProfilePath    = "C:\\crio\\etc\\seccomp.json"
	cniConfigDir          = "C:\\cni\\etc\\net.d\\"
	cniBinDir             = "C:\\cni\\bin\\"
	lockPath              = "C:\\crio\\run\\crio.lock"
	containerExitsDir     = oci.ContainerExitsDir
	defaultOCIRuntimePath = "C:\\crio\\bin\\runchcs"
	defaultLogDir         = "C:\\crio\\log\\pods\\"
)
