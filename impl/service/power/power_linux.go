// +build linux

package power

import (
	"syscall"
	"time"
)

func shutdown(wait time.Duration) chan error {
	return execPower(syscall.LINUX_REBOOT_CMD_POWER_OFF, wait)
}

func reboot(wait time.Duration) chan error {
	return execPower(syscall.LINUX_REBOOT_CMD_RESTART, wait)
}

func execPower(cmd int, wait time.Duration) chan error {
	// this channel will only be used when error happened
	ch := make(chan error)
	go func() {
		time.Sleep(wait)
		ch <- syscall.Reboot(cmd)
	}()

	return ch
}
