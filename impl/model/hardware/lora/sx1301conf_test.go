package lora

import "testing"

func TestGetPktFwdConfig(t *testing.T) {
	if c := GetDefaultPktFwdConfig(); c == defaultPktFwdConfig {
		t.Fail()
	}
}
