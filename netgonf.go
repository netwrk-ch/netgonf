package Netgonf

import (
	"fmt"

	"github.com/Juniper/go-netconf/netconf"
)

type Session interface {
	Exec(...netconf.RPCMethod) (*netconf.RPCReply, error)
	Close() error
}

func methodGet(filter string) netconf.RPCMethod {
	baseRpc := `<get><filter>%s</filter></get>`
	return netconf.RawMethod(fmt.Sprintf(baseRpc, filter))
}

func methodGetRunningConfig(filter string) netconf.RPCMethod {
	baseRpc := `<get-config><source><running/></source><filter>%s</filter></get-config>`
	return netconf.RawMethod(fmt.Sprintf(baseRpc, filter))
}

func execRPC(session Session, rpc netconf.RPCMethod) (*netconf.RPCReply, error) {
	reply, err := session.Exec(rpc)
	if err != nil {
		return nil, err
	}
	return reply, err
}

func UnlockDatastore(session Session) (result string, err error) {
	raw := netconf.RawMethod(`<unlock><target><running/></target></unlock>`)
	reply, err := session.Exec(raw)
	if err != nil {
		return "", err
	}
	return reply.Data, nil
}

func LockDatastore(session Session) (result string, err error) {
	raw := netconf.RawMethod(`<lock><target><running/></target></lock>`)
	reply, err := session.Exec(raw)
	if err != nil {
		return "", err
	}
	return reply.Data, nil
}

func SaveConfig(session Session) (result string, err error) {
	raw := netconf.RawMethod(`<save-config xmlns="http://cisco.com/yang/cisco-ia"/>`)
	reply, err := session.Exec(raw)
	if err != nil {
		return "", err
	}
	return reply.Data, nil
}

func ApplyConfig(session Session, config string) (result string, err error) {
	raw := netconf.MethodEditConfig("running", config)
	reply, err := session.Exec(raw)
	if err != nil {
		return "", err
	}
	return reply.Data, nil
}

func GetRunningConfig(session Session, filter string) (*netconf.RPCReply, error) {
	return execRPC(session, methodGetRunningConfig(filter))
}

func Get(session Session, filter string) (*netconf.RPCReply, error) {
	return execRPC(session, methodGet(filter))
}
