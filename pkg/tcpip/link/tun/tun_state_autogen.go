// automatically generated by stateify.

package tun

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (d *Device) StateTypeName() string {
	return "pkg/tcpip/link/tun.Device"
}

func (d *Device) StateFields() []string {
	return []string{
		"Queue",
		"endpoint",
		"notifyHandle",
		"flags",
	}
}

// +checklocksignore
func (d *Device) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Queue)
	stateSinkObject.Save(1, &d.endpoint)
	stateSinkObject.Save(2, &d.notifyHandle)
	stateSinkObject.Save(3, &d.flags)
}

func (d *Device) afterLoad() {}

// +checklocksignore
func (d *Device) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Queue)
	stateSourceObject.Load(1, &d.endpoint)
	stateSourceObject.Load(2, &d.notifyHandle)
	stateSourceObject.Load(3, &d.flags)
}

func (r *tunEndpointRefs) StateTypeName() string {
	return "pkg/tcpip/link/tun.tunEndpointRefs"
}

func (r *tunEndpointRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *tunEndpointRefs) beforeSave() {}

// +checklocksignore
func (r *tunEndpointRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *tunEndpointRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*Device)(nil))
	state.Register((*tunEndpointRefs)(nil))
}
