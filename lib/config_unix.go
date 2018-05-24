// +build !windows

package lib

import "github.com/kubernetes-incubator/cri-o/oci"

// Defaults for linux/unix if none are specified
var (
	conmonPath            = "/usr/local/libexec/crio/conmon"
	seccompProfilePath    = "/etc/crio/seccomp.json"
	cniConfigDir          = "/etc/cni/net.d/"
	cniBinDir             = "/opt/cni/bin/"
	lockPath              = "/run/crio.lock"
	containerExitsDir     = oci.ContainerExitsDir
	defaultOCIRuntimePath = "/usr/bin/runc"
	defaultLogDir         = "/var/log/crio/pods"
)
