package darwin

import (
	"fmt"
	"syscall"
)

const kernelOSReleaseMIB = "kern.osrelease"

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
