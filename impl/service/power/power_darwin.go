// +build darwin

package power

import (
	"time"
)

func execPower(wait time.Duration) chan error {
	ch := make(chan error)
	return ch
}

func shutdown(wait time.Duration) chan error {
	// TODO
	return execPower(wait)
}

func reboot(wait time.Duration) chan error {
	// TODO
	return execPower(wait)
}
