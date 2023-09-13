package Netgonf

import (
	"testing"

	"github.com/Juniper/go-netconf/netconf"
)

type TestSession struct {
}

func (t *TestSession) Exec(...netconf.RPCMethod) (*netconf.RPCReply, error) {
	return nil, nil
}

func (t *TestSession) Close() error {
	return nil
}

func TestMethodGet(t *testing.T) {
	t.Run("TestMethodGet", func(t *testing.T) {
		want := netconf.RawMethod("<get><filter>test</filter></get>")
		got := methodGet("test")
		if got != want {
			t.Errorf("methodGet() = %q, want %q", got, want)
		}
	})
}

func TestMethodGetRunningConfig(t *testing.T) {
	t.Run("TestMethodGetRunningConfig", func(t *testing.T) {
		want := netconf.RawMethod("<get-config><source><running/></source><filter>test</filter></get-config>")
		got := methodGetRunningConfig("test")
		if got != want {
			t.Errorf("methodGet() = %q, want %q", got, want)
		}
	})
}
