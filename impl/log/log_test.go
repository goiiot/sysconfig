package log

import "testing"

func TestDebugLogger(t *testing.T) {
	Init(Debug)
	D("debug")
	I("info")
	W("warning")
	E("error")
	if err := F("fatal", true); err == nil {
		t.Fail()
	}
}

func TestInfoLogger(t *testing.T) {
	Init(Info)
	D("debug")
	I("info")
	W("warning")
	E("error")
	if err := F("fatal", true); err == nil {
		t.Fail()
	}
}

func TestWarningLogger(t *testing.T) {
	Init(Warning)
	D("debug")
	I("info")
	W("warning")
	E("error")
	if err := F("fatal", true); err == nil {
		t.Fail()
	}
}

func TestErrorLogger(t *testing.T) {
	Init(Error)
	D("debug")
	I("info")
	W("warning")
	E("error")
	if err := F("fatal", true); err == nil {
		t.Fail()
	}
}

func TestFatalLogger(t *testing.T) {
	Init(Fatal)
	D("debug")
	I("info")
	W("warning")
	E("error")
	if err := F("fatal", true); err == nil {
		t.Fail()
	}
}

func TestSilentLogger(t *testing.T) {
	Init(Silent)
	D("debug")
	I("info")
	W("warning")
	E("error")
	if err := F("fatal", true); err != nil {
		t.Fail()
	}
}
