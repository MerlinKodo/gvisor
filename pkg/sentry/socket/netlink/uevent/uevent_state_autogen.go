// automatically generated by stateify.

package uevent

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (p *Protocol) StateTypeName() string {
	return "pkg/sentry/socket/netlink/uevent.Protocol"
}

func (p *Protocol) StateFields() []string {
	return []string{}
}

func (p *Protocol) beforeSave() {}

// +checklocksignore
func (p *Protocol) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
}

func (p *Protocol) afterLoad() {}

// +checklocksignore
func (p *Protocol) StateLoad(stateSourceObject state.Source) {
}

func init() {
	state.Register((*Protocol)(nil))
}
