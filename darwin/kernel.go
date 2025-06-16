package darwin

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	kernelBootTimeMIB  = "kern.boottime"
	kernelHostNameMIB  = "kern.hostname"
	kernelOSReleaseMIB = "kern.osrelease"
)

func KernelBootTime() (time.Duration, error) {
	bootTime, err := unix.SysctlRaw(kernelBootTimeMIB)
	var timeval syscall.Timeval

	if err != nil {
		return 0, fmt.Errorf("failed to get kernel clockrate: %w", err)
	}

	if len(bootTime) != int(unsafe.Sizeof(timeval)) {
		return 0, fmt.Errorf("unexpected output of sysctl kern.boottime: %v (len: %d)", bootTime, len(bootTime))
	}

	timeval = *(*syscall.Timeval)(unsafe.Pointer(&bootTime[0]))
	sec, nsec := timeval.Unix()

	return time.Since(time.Unix(sec, nsec)), nil
}

func KernelHostName() (string, error) {
	hostName, err := syscall.Sysctl(kernelHostNameMIB)

	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %w", err)
	}

	return hostName, nil
}

// Returns the Darwin Kernel Version of the current operating system. Utilizes
// the `kern.osrelease` MIB name.
//
// This is the equivalent of `uname -r`.
func KernelVersion() (string, error) {
	version, err := syscall.Sysctl(kernelOSReleaseMIB)

	if err != nil {
		return "", fmt.Errorf("failed to get kernel version: %w", err)
	}

	return version, nil
}
