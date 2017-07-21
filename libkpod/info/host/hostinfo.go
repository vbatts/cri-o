package host

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/docker/docker/pkg/system"
	"github.com/kubernetes-incubator/cri-o/libkpod/info"
	"golang.org/x/net/context"
)

func init() {
	info.RegisterInfoGiver(hostInfo)
}

func hostInfo(ctx context.Context) (string, map[string]interface{}, error) {
	// lets say OS, arch, number of cpus, amount of memory, maybe os distribution/version, hostname, kernel version, uptime
	i := map[string]interface{}{}
	i["os"] = runtime.GOOS
	i["arch"] = runtime.GOARCH
	i["cpus"] = runtime.NumCPU()
	mi, err := system.ReadMemInfo()
	if err != nil {
		i["meminfo"] = info.Err(err)
	} else {
		// TODO this might be a place for github.com/dustin/go-humanize
		i["MemTotal"] = mi.MemTotal
		i["MemFree"] = mi.MemFree
		i["SwapTotal"] = mi.SwapTotal
		i["SwapFree"] = mi.SwapFree
	}
	if kv, err := readKernelVersion(); err != nil {
		i["kernel"] = info.Err(err)
	} else {
		i["kernel"] = kv
	}

	if up, err := readUptime(); err != nil {
		i["uptime"] = info.Err(err)
	} else {
		i["uptime"] = up
	}
	if host, err := os.Hostname(); err != nil {
		i["hostname"] = info.Err(err)
	} else {
		i["hostname"] = host
	}
	return "host", i, nil
}

func readKernelVersion() (string, error) {
	buf, err := ioutil.ReadFile("/proc/version")
	if err != nil {
		return "", err
	}
	f := bytes.Fields(buf)
	if len(f) < 2 {
		return string(bytes.TrimSpace(buf)), nil
	}
	return string(f[2]), nil
}

func readUptime() (string, error) {
	buf, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return "", err
	}
	f := bytes.Fields(buf)
	if len(f) < 1 {
		return "", fmt.Errorf("invalid uptime")
	}
	return string(f[0]), nil
}
