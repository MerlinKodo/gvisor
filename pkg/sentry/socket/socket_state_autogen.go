// automatically generated by stateify.

package socket

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (i *IPControlMessages) StateTypeName() string {
	return "pkg/sentry/socket.IPControlMessages"
}

func (i *IPControlMessages) StateFields() []string {
	return []string{
		"HasTimestamp",
		"Timestamp",
		"HasInq",
		"Inq",
		"HasTOS",
		"TOS",
		"HasTTL",
		"TTL",
		"HasHopLimit",
		"HopLimit",
		"HasTClass",
		"TClass",
		"HasIPPacketInfo",
		"PacketInfo",
		"HasIPv6PacketInfo",
		"IPv6PacketInfo",
		"OriginalDstAddress",
		"SockErr",
	}
}

func (i *IPControlMessages) beforeSave() {}

// +checklocksignore
func (i *IPControlMessages) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	var TimestampValue int64
	TimestampValue = i.saveTimestamp()
	stateSinkObject.SaveValue(1, TimestampValue)
	stateSinkObject.Save(0, &i.HasTimestamp)
	stateSinkObject.Save(2, &i.HasInq)
	stateSinkObject.Save(3, &i.Inq)
	stateSinkObject.Save(4, &i.HasTOS)
	stateSinkObject.Save(5, &i.TOS)
	stateSinkObject.Save(6, &i.HasTTL)
	stateSinkObject.Save(7, &i.TTL)
	stateSinkObject.Save(8, &i.HasHopLimit)
	stateSinkObject.Save(9, &i.HopLimit)
	stateSinkObject.Save(10, &i.HasTClass)
	stateSinkObject.Save(11, &i.TClass)
	stateSinkObject.Save(12, &i.HasIPPacketInfo)
	stateSinkObject.Save(13, &i.PacketInfo)
	stateSinkObject.Save(14, &i.HasIPv6PacketInfo)
	stateSinkObject.Save(15, &i.IPv6PacketInfo)
	stateSinkObject.Save(16, &i.OriginalDstAddress)
	stateSinkObject.Save(17, &i.SockErr)
}

func (i *IPControlMessages) afterLoad() {}

// +checklocksignore
func (i *IPControlMessages) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.HasTimestamp)
	stateSourceObject.Load(2, &i.HasInq)
	stateSourceObject.Load(3, &i.Inq)
	stateSourceObject.Load(4, &i.HasTOS)
	stateSourceObject.Load(5, &i.TOS)
	stateSourceObject.Load(6, &i.HasTTL)
	stateSourceObject.Load(7, &i.TTL)
	stateSourceObject.Load(8, &i.HasHopLimit)
	stateSourceObject.Load(9, &i.HopLimit)
	stateSourceObject.Load(10, &i.HasTClass)
	stateSourceObject.Load(11, &i.TClass)
	stateSourceObject.Load(12, &i.HasIPPacketInfo)
	stateSourceObject.Load(13, &i.PacketInfo)
	stateSourceObject.Load(14, &i.HasIPv6PacketInfo)
	stateSourceObject.Load(15, &i.IPv6PacketInfo)
	stateSourceObject.Load(16, &i.OriginalDstAddress)
	stateSourceObject.Load(17, &i.SockErr)
	stateSourceObject.LoadValue(1, new(int64), func(y any) { i.loadTimestamp(y.(int64)) })
}

func (to *SendReceiveTimeout) StateTypeName() string {
	return "pkg/sentry/socket.SendReceiveTimeout"
}

func (to *SendReceiveTimeout) StateFields() []string {
	return []string{
		"send",
		"recv",
	}
}

func (to *SendReceiveTimeout) beforeSave() {}

// +checklocksignore
func (to *SendReceiveTimeout) StateSave(stateSinkObject state.Sink) {
	to.beforeSave()
	stateSinkObject.Save(0, &to.send)
	stateSinkObject.Save(1, &to.recv)
}

func (to *SendReceiveTimeout) afterLoad() {}

// +checklocksignore
func (to *SendReceiveTimeout) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &to.send)
	stateSourceObject.Load(1, &to.recv)
}

func init() {
	state.Register((*IPControlMessages)(nil))
	state.Register((*SendReceiveTimeout)(nil))
}
