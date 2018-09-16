package configure

import (
	"os/exec"
	"strings"
)

func run(shell string, cmd ...string) ([]byte, error) {
	arg1 := func() string {
		switch strings.TrimSpace(strings.ToLower(shell)) {
		case "bash", "sh", "zsh":
			return "-c"
		default:
			return "-c"
		}
	}()

	result, err := exec.Command(shell, arg1, strings.Join(cmd, " ")).Output()
	return result, err
}
