package configure

import "testing"

const testCmd = "ps -A | grep ps && ls . && go version | awk '{print $3;}'"

func TestRun(t *testing.T) {
	result, err := run("zsh", testCmd)
	if err != nil {
		t.Errorf("run zsh test failed: %v", err)
	}
	println(result)

	result, err = run("bash", testCmd)
	if err != nil {
		t.Errorf("run bash test failed: %v", err)
	}
	println(result)

	result, err = run("sh", testCmd)
	if err != nil {
		t.Errorf("run sh test failed: %v", err)
	}
	println(result)
}
