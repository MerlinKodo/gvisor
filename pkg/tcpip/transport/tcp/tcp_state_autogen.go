// automatically generated by stateify.

package tcp

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (a *acceptQueue) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.acceptQueue"
}

func (a *acceptQueue) StateFields() []string {
	return []string{
		"endpoints",
		"pendingEndpoints",
		"capacity",
	}
}

func (a *acceptQueue) beforeSave() {}

// +checklocksignore
func (a *acceptQueue) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	var endpointsValue []*endpoint
	endpointsValue = a.saveEndpoints()
	stateSinkObject.SaveValue(0, endpointsValue)
	stateSinkObject.Save(1, &a.pendingEndpoints)
	stateSinkObject.Save(2, &a.capacity)
}

func (a *acceptQueue) afterLoad() {}

// +checklocksignore
func (a *acceptQueue) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(1, &a.pendingEndpoints)
	stateSourceObject.Load(2, &a.capacity)
	stateSourceObject.LoadValue(0, new([]*endpoint), func(y any) { a.loadEndpoints(y.([]*endpoint)) })
}

func (h *handshake) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.handshake"
}

func (h *handshake) StateFields() []string {
	return []string{
		"ep",
		"listenEP",
		"state",
		"active",
		"flags",
		"ackNum",
		"iss",
		"rcvWnd",
		"sndWnd",
		"mss",
		"sndWndScale",
		"rcvWndScale",
		"startTime",
		"deferAccept",
		"acked",
		"sendSYNOpts",
		"sampleRTTWithTSOnly",
	}
}

func (h *handshake) beforeSave() {}

// +checklocksignore
func (h *handshake) StateSave(stateSinkObject state.Sink) {
	h.beforeSave()
	stateSinkObject.Save(0, &h.ep)
	stateSinkObject.Save(1, &h.listenEP)
	stateSinkObject.Save(2, &h.state)
	stateSinkObject.Save(3, &h.active)
	stateSinkObject.Save(4, &h.flags)
	stateSinkObject.Save(5, &h.ackNum)
	stateSinkObject.Save(6, &h.iss)
	stateSinkObject.Save(7, &h.rcvWnd)
	stateSinkObject.Save(8, &h.sndWnd)
	stateSinkObject.Save(9, &h.mss)
	stateSinkObject.Save(10, &h.sndWndScale)
	stateSinkObject.Save(11, &h.rcvWndScale)
	stateSinkObject.Save(12, &h.startTime)
	stateSinkObject.Save(13, &h.deferAccept)
	stateSinkObject.Save(14, &h.acked)
	stateSinkObject.Save(15, &h.sendSYNOpts)
	stateSinkObject.Save(16, &h.sampleRTTWithTSOnly)
}

func (h *handshake) afterLoad() {}

// +checklocksignore
func (h *handshake) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &h.ep)
	stateSourceObject.Load(1, &h.listenEP)
	stateSourceObject.Load(2, &h.state)
	stateSourceObject.Load(3, &h.active)
	stateSourceObject.Load(4, &h.flags)
	stateSourceObject.Load(5, &h.ackNum)
	stateSourceObject.Load(6, &h.iss)
	stateSourceObject.Load(7, &h.rcvWnd)
	stateSourceObject.Load(8, &h.sndWnd)
	stateSourceObject.Load(9, &h.mss)
	stateSourceObject.Load(10, &h.sndWndScale)
	stateSourceObject.Load(11, &h.rcvWndScale)
	stateSourceObject.Load(12, &h.startTime)
	stateSourceObject.Load(13, &h.deferAccept)
	stateSourceObject.Load(14, &h.acked)
	stateSourceObject.Load(15, &h.sendSYNOpts)
	stateSourceObject.Load(16, &h.sampleRTTWithTSOnly)
}

func (c *cubicState) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.cubicState"
}

func (c *cubicState) StateFields() []string {
	return []string{
		"TCPCubicState",
		"numCongestionEvents",
		"s",
	}
}

func (c *cubicState) beforeSave() {}

// +checklocksignore
func (c *cubicState) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.TCPCubicState)
	stateSinkObject.Save(1, &c.numCongestionEvents)
	stateSinkObject.Save(2, &c.s)
}

func (c *cubicState) afterLoad() {}

// +checklocksignore
func (c *cubicState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.TCPCubicState)
	stateSourceObject.Load(1, &c.numCongestionEvents)
	stateSourceObject.Load(2, &c.s)
}

func (s *SACKInfo) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.SACKInfo"
}

func (s *SACKInfo) StateFields() []string {
	return []string{
		"Blocks",
		"NumBlocks",
	}
}

func (s *SACKInfo) beforeSave() {}

// +checklocksignore
func (s *SACKInfo) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Blocks)
	stateSinkObject.Save(1, &s.NumBlocks)
}

func (s *SACKInfo) afterLoad() {}

// +checklocksignore
func (s *SACKInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Blocks)
	stateSourceObject.Load(1, &s.NumBlocks)
}

func (r *ReceiveErrors) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.ReceiveErrors"
}

func (r *ReceiveErrors) StateFields() []string {
	return []string{
		"ReceiveErrors",
		"SegmentQueueDropped",
		"ChecksumErrors",
		"ListenOverflowSynDrop",
		"ListenOverflowAckDrop",
		"ZeroRcvWindowState",
		"WantZeroRcvWindow",
	}
}

func (r *ReceiveErrors) beforeSave() {}

// +checklocksignore
func (r *ReceiveErrors) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.ReceiveErrors)
	stateSinkObject.Save(1, &r.SegmentQueueDropped)
	stateSinkObject.Save(2, &r.ChecksumErrors)
	stateSinkObject.Save(3, &r.ListenOverflowSynDrop)
	stateSinkObject.Save(4, &r.ListenOverflowAckDrop)
	stateSinkObject.Save(5, &r.ZeroRcvWindowState)
	stateSinkObject.Save(6, &r.WantZeroRcvWindow)
}

func (r *ReceiveErrors) afterLoad() {}

// +checklocksignore
func (r *ReceiveErrors) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.ReceiveErrors)
	stateSourceObject.Load(1, &r.SegmentQueueDropped)
	stateSourceObject.Load(2, &r.ChecksumErrors)
	stateSourceObject.Load(3, &r.ListenOverflowSynDrop)
	stateSourceObject.Load(4, &r.ListenOverflowAckDrop)
	stateSourceObject.Load(5, &r.ZeroRcvWindowState)
	stateSourceObject.Load(6, &r.WantZeroRcvWindow)
}

func (s *SendErrors) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.SendErrors"
}

func (s *SendErrors) StateFields() []string {
	return []string{
		"SendErrors",
		"SegmentSendToNetworkFailed",
		"SynSendToNetworkFailed",
		"Retransmits",
		"FastRetransmit",
		"Timeouts",
	}
}

func (s *SendErrors) beforeSave() {}

// +checklocksignore
func (s *SendErrors) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.SendErrors)
	stateSinkObject.Save(1, &s.SegmentSendToNetworkFailed)
	stateSinkObject.Save(2, &s.SynSendToNetworkFailed)
	stateSinkObject.Save(3, &s.Retransmits)
	stateSinkObject.Save(4, &s.FastRetransmit)
	stateSinkObject.Save(5, &s.Timeouts)
}

func (s *SendErrors) afterLoad() {}

// +checklocksignore
func (s *SendErrors) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.SendErrors)
	stateSourceObject.Load(1, &s.SegmentSendToNetworkFailed)
	stateSourceObject.Load(2, &s.SynSendToNetworkFailed)
	stateSourceObject.Load(3, &s.Retransmits)
	stateSourceObject.Load(4, &s.FastRetransmit)
	stateSourceObject.Load(5, &s.Timeouts)
}

func (s *Stats) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.Stats"
}

func (s *Stats) StateFields() []string {
	return []string{
		"SegmentsReceived",
		"SegmentsSent",
		"FailedConnectionAttempts",
		"ReceiveErrors",
		"ReadErrors",
		"SendErrors",
		"WriteErrors",
	}
}

func (s *Stats) beforeSave() {}

// +checklocksignore
func (s *Stats) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.SegmentsReceived)
	stateSinkObject.Save(1, &s.SegmentsSent)
	stateSinkObject.Save(2, &s.FailedConnectionAttempts)
	stateSinkObject.Save(3, &s.ReceiveErrors)
	stateSinkObject.Save(4, &s.ReadErrors)
	stateSinkObject.Save(5, &s.SendErrors)
	stateSinkObject.Save(6, &s.WriteErrors)
}

func (s *Stats) afterLoad() {}

// +checklocksignore
func (s *Stats) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.SegmentsReceived)
	stateSourceObject.Load(1, &s.SegmentsSent)
	stateSourceObject.Load(2, &s.FailedConnectionAttempts)
	stateSourceObject.Load(3, &s.ReceiveErrors)
	stateSourceObject.Load(4, &s.ReadErrors)
	stateSourceObject.Load(5, &s.SendErrors)
	stateSourceObject.Load(6, &s.WriteErrors)
}

func (sq *sndQueueInfo) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.sndQueueInfo"
}

func (sq *sndQueueInfo) StateFields() []string {
	return []string{
		"TCPSndBufState",
	}
}

func (sq *sndQueueInfo) beforeSave() {}

// +checklocksignore
func (sq *sndQueueInfo) StateSave(stateSinkObject state.Sink) {
	sq.beforeSave()
	stateSinkObject.Save(0, &sq.TCPSndBufState)
}

func (sq *sndQueueInfo) afterLoad() {}

// +checklocksignore
func (sq *sndQueueInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &sq.TCPSndBufState)
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"TCPEndpointStateInner",
		"TransportEndpointInfo",
		"DefaultSocketOptionsHandler",
		"waiterQueue",
		"uniqueID",
		"hardError",
		"lastError",
		"TCPRcvBufState",
		"rcvMemUsed",
		"ownedByUser",
		"rcvQueue",
		"state",
		"boundNICID",
		"ipv4TTL",
		"ipv6HopLimit",
		"isConnectNotified",
		"h",
		"portFlags",
		"boundBindToDevice",
		"boundPortFlags",
		"boundDest",
		"effectiveNetProtos",
		"recentTSTime",
		"shutdownFlags",
		"tcpRecovery",
		"sack",
		"delay",
		"scoreboard",
		"segmentQueue",
		"userMSS",
		"maxSynRetries",
		"windowClamp",
		"sndQueueInfo",
		"cc",
		"keepalive",
		"userTimeout",
		"deferAccept",
		"acceptQueue",
		"rcv",
		"snd",
		"connectingAddress",
		"amss",
		"sendTOS",
		"gso",
		"stats",
		"tcpLingerTimeout",
		"closed",
		"txHash",
		"owner",
		"ops",
		"lastOutOfWindowAckTime",
	}
}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	var stateValue EndpointState
	stateValue = e.saveState()
	stateSinkObject.SaveValue(11, stateValue)
	stateSinkObject.Save(0, &e.TCPEndpointStateInner)
	stateSinkObject.Save(1, &e.TransportEndpointInfo)
	stateSinkObject.Save(2, &e.DefaultSocketOptionsHandler)
	stateSinkObject.Save(3, &e.waiterQueue)
	stateSinkObject.Save(4, &e.uniqueID)
	stateSinkObject.Save(5, &e.hardError)
	stateSinkObject.Save(6, &e.lastError)
	stateSinkObject.Save(7, &e.TCPRcvBufState)
	stateSinkObject.Save(8, &e.rcvMemUsed)
	stateSinkObject.Save(9, &e.ownedByUser)
	stateSinkObject.Save(10, &e.rcvQueue)
	stateSinkObject.Save(12, &e.boundNICID)
	stateSinkObject.Save(13, &e.ipv4TTL)
	stateSinkObject.Save(14, &e.ipv6HopLimit)
	stateSinkObject.Save(15, &e.isConnectNotified)
	stateSinkObject.Save(16, &e.h)
	stateSinkObject.Save(17, &e.portFlags)
	stateSinkObject.Save(18, &e.boundBindToDevice)
	stateSinkObject.Save(19, &e.boundPortFlags)
	stateSinkObject.Save(20, &e.boundDest)
	stateSinkObject.Save(21, &e.effectiveNetProtos)
	stateSinkObject.Save(22, &e.recentTSTime)
	stateSinkObject.Save(23, &e.shutdownFlags)
	stateSinkObject.Save(24, &e.tcpRecovery)
	stateSinkObject.Save(25, &e.sack)
	stateSinkObject.Save(26, &e.delay)
	stateSinkObject.Save(27, &e.scoreboard)
	stateSinkObject.Save(28, &e.segmentQueue)
	stateSinkObject.Save(29, &e.userMSS)
	stateSinkObject.Save(30, &e.maxSynRetries)
	stateSinkObject.Save(31, &e.windowClamp)
	stateSinkObject.Save(32, &e.sndQueueInfo)
	stateSinkObject.Save(33, &e.cc)
	stateSinkObject.Save(34, &e.keepalive)
	stateSinkObject.Save(35, &e.userTimeout)
	stateSinkObject.Save(36, &e.deferAccept)
	stateSinkObject.Save(37, &e.acceptQueue)
	stateSinkObject.Save(38, &e.rcv)
	stateSinkObject.Save(39, &e.snd)
	stateSinkObject.Save(40, &e.connectingAddress)
	stateSinkObject.Save(41, &e.amss)
	stateSinkObject.Save(42, &e.sendTOS)
	stateSinkObject.Save(43, &e.gso)
	stateSinkObject.Save(44, &e.stats)
	stateSinkObject.Save(45, &e.tcpLingerTimeout)
	stateSinkObject.Save(46, &e.closed)
	stateSinkObject.Save(47, &e.txHash)
	stateSinkObject.Save(48, &e.owner)
	stateSinkObject.Save(49, &e.ops)
	stateSinkObject.Save(50, &e.lastOutOfWindowAckTime)
}

// +checklocksignore
func (e *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.TCPEndpointStateInner)
	stateSourceObject.Load(1, &e.TransportEndpointInfo)
	stateSourceObject.Load(2, &e.DefaultSocketOptionsHandler)
	stateSourceObject.LoadWait(3, &e.waiterQueue)
	stateSourceObject.Load(4, &e.uniqueID)
	stateSourceObject.Load(5, &e.hardError)
	stateSourceObject.Load(6, &e.lastError)
	stateSourceObject.Load(7, &e.TCPRcvBufState)
	stateSourceObject.Load(8, &e.rcvMemUsed)
	stateSourceObject.Load(9, &e.ownedByUser)
	stateSourceObject.LoadWait(10, &e.rcvQueue)
	stateSourceObject.Load(12, &e.boundNICID)
	stateSourceObject.Load(13, &e.ipv4TTL)
	stateSourceObject.Load(14, &e.ipv6HopLimit)
	stateSourceObject.Load(15, &e.isConnectNotified)
	stateSourceObject.Load(16, &e.h)
	stateSourceObject.Load(17, &e.portFlags)
	stateSourceObject.Load(18, &e.boundBindToDevice)
	stateSourceObject.Load(19, &e.boundPortFlags)
	stateSourceObject.Load(20, &e.boundDest)
	stateSourceObject.Load(21, &e.effectiveNetProtos)
	stateSourceObject.Load(22, &e.recentTSTime)
	stateSourceObject.Load(23, &e.shutdownFlags)
	stateSourceObject.Load(24, &e.tcpRecovery)
	stateSourceObject.Load(25, &e.sack)
	stateSourceObject.Load(26, &e.delay)
	stateSourceObject.Load(27, &e.scoreboard)
	stateSourceObject.LoadWait(28, &e.segmentQueue)
	stateSourceObject.Load(29, &e.userMSS)
	stateSourceObject.Load(30, &e.maxSynRetries)
	stateSourceObject.Load(31, &e.windowClamp)
	stateSourceObject.Load(32, &e.sndQueueInfo)
	stateSourceObject.Load(33, &e.cc)
	stateSourceObject.Load(34, &e.keepalive)
	stateSourceObject.Load(35, &e.userTimeout)
	stateSourceObject.Load(36, &e.deferAccept)
	stateSourceObject.Load(37, &e.acceptQueue)
	stateSourceObject.LoadWait(38, &e.rcv)
	stateSourceObject.LoadWait(39, &e.snd)
	stateSourceObject.Load(40, &e.connectingAddress)
	stateSourceObject.Load(41, &e.amss)
	stateSourceObject.Load(42, &e.sendTOS)
	stateSourceObject.Load(43, &e.gso)
	stateSourceObject.Load(44, &e.stats)
	stateSourceObject.Load(45, &e.tcpLingerTimeout)
	stateSourceObject.Load(46, &e.closed)
	stateSourceObject.Load(47, &e.txHash)
	stateSourceObject.Load(48, &e.owner)
	stateSourceObject.Load(49, &e.ops)
	stateSourceObject.Load(50, &e.lastOutOfWindowAckTime)
	stateSourceObject.LoadValue(11, new(EndpointState), func(y any) { e.loadState(y.(EndpointState)) })
	stateSourceObject.AfterLoad(e.afterLoad)
}

func (k *keepalive) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.keepalive"
}

func (k *keepalive) StateFields() []string {
	return []string{
		"idle",
		"interval",
		"count",
		"unacked",
	}
}

func (k *keepalive) beforeSave() {}

// +checklocksignore
func (k *keepalive) StateSave(stateSinkObject state.Sink) {
	k.beforeSave()
	stateSinkObject.Save(0, &k.idle)
	stateSinkObject.Save(1, &k.interval)
	stateSinkObject.Save(2, &k.count)
	stateSinkObject.Save(3, &k.unacked)
}

func (k *keepalive) afterLoad() {}

// +checklocksignore
func (k *keepalive) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &k.idle)
	stateSourceObject.Load(1, &k.interval)
	stateSourceObject.Load(2, &k.count)
	stateSourceObject.Load(3, &k.unacked)
}

func (rc *rackControl) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.rackControl"
}

func (rc *rackControl) StateFields() []string {
	return []string{
		"TCPRACKState",
		"exitedRecovery",
		"minRTT",
		"tlpRxtOut",
		"tlpHighRxt",
		"snd",
	}
}

func (rc *rackControl) beforeSave() {}

// +checklocksignore
func (rc *rackControl) StateSave(stateSinkObject state.Sink) {
	rc.beforeSave()
	stateSinkObject.Save(0, &rc.TCPRACKState)
	stateSinkObject.Save(1, &rc.exitedRecovery)
	stateSinkObject.Save(2, &rc.minRTT)
	stateSinkObject.Save(3, &rc.tlpRxtOut)
	stateSinkObject.Save(4, &rc.tlpHighRxt)
	stateSinkObject.Save(5, &rc.snd)
}

func (rc *rackControl) afterLoad() {}

// +checklocksignore
func (rc *rackControl) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rc.TCPRACKState)
	stateSourceObject.Load(1, &rc.exitedRecovery)
	stateSourceObject.Load(2, &rc.minRTT)
	stateSourceObject.Load(3, &rc.tlpRxtOut)
	stateSourceObject.Load(4, &rc.tlpHighRxt)
	stateSourceObject.Load(5, &rc.snd)
}

func (r *receiver) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.receiver"
}

func (r *receiver) StateFields() []string {
	return []string{
		"TCPReceiverState",
		"ep",
		"rcvWnd",
		"rcvWUP",
		"prevBufUsed",
		"closed",
		"pendingRcvdSegments",
		"lastRcvdAckTime",
	}
}

func (r *receiver) beforeSave() {}

// +checklocksignore
func (r *receiver) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.TCPReceiverState)
	stateSinkObject.Save(1, &r.ep)
	stateSinkObject.Save(2, &r.rcvWnd)
	stateSinkObject.Save(3, &r.rcvWUP)
	stateSinkObject.Save(4, &r.prevBufUsed)
	stateSinkObject.Save(5, &r.closed)
	stateSinkObject.Save(6, &r.pendingRcvdSegments)
	stateSinkObject.Save(7, &r.lastRcvdAckTime)
}

func (r *receiver) afterLoad() {}

// +checklocksignore
func (r *receiver) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.TCPReceiverState)
	stateSourceObject.Load(1, &r.ep)
	stateSourceObject.Load(2, &r.rcvWnd)
	stateSourceObject.Load(3, &r.rcvWUP)
	stateSourceObject.Load(4, &r.prevBufUsed)
	stateSourceObject.Load(5, &r.closed)
	stateSourceObject.Load(6, &r.pendingRcvdSegments)
	stateSourceObject.Load(7, &r.lastRcvdAckTime)
}

func (r *renoState) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.renoState"
}

func (r *renoState) StateFields() []string {
	return []string{
		"s",
	}
}

func (r *renoState) beforeSave() {}

// +checklocksignore
func (r *renoState) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.s)
}

func (r *renoState) afterLoad() {}

// +checklocksignore
func (r *renoState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.s)
}

func (rr *renoRecovery) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.renoRecovery"
}

func (rr *renoRecovery) StateFields() []string {
	return []string{
		"s",
	}
}

func (rr *renoRecovery) beforeSave() {}

// +checklocksignore
func (rr *renoRecovery) StateSave(stateSinkObject state.Sink) {
	rr.beforeSave()
	stateSinkObject.Save(0, &rr.s)
}

func (rr *renoRecovery) afterLoad() {}

// +checklocksignore
func (rr *renoRecovery) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rr.s)
}

func (sr *sackRecovery) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.sackRecovery"
}

func (sr *sackRecovery) StateFields() []string {
	return []string{
		"s",
	}
}

func (sr *sackRecovery) beforeSave() {}

// +checklocksignore
func (sr *sackRecovery) StateSave(stateSinkObject state.Sink) {
	sr.beforeSave()
	stateSinkObject.Save(0, &sr.s)
}

func (sr *sackRecovery) afterLoad() {}

// +checklocksignore
func (sr *sackRecovery) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &sr.s)
}

func (s *SACKScoreboard) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.SACKScoreboard"
}

func (s *SACKScoreboard) StateFields() []string {
	return []string{
		"smss",
		"maxSACKED",
	}
}

func (s *SACKScoreboard) beforeSave() {}

// +checklocksignore
func (s *SACKScoreboard) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.smss)
	stateSinkObject.Save(1, &s.maxSACKED)
}

func (s *SACKScoreboard) afterLoad() {}

// +checklocksignore
func (s *SACKScoreboard) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.smss)
	stateSourceObject.Load(1, &s.maxSACKED)
}

func (s *segment) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segment"
}

func (s *segment) StateFields() []string {
	return []string{
		"segmentEntry",
		"segmentRefs",
		"ep",
		"qFlags",
		"pkt",
		"sequenceNumber",
		"ackNumber",
		"flags",
		"window",
		"csum",
		"csumValid",
		"parsedOptions",
		"options",
		"hasNewSACKInfo",
		"rcvdTime",
		"xmitTime",
		"xmitCount",
		"acked",
		"dataMemSize",
		"lost",
	}
}

func (s *segment) beforeSave() {}

// +checklocksignore
func (s *segment) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var optionsValue []byte
	optionsValue = s.saveOptions()
	stateSinkObject.SaveValue(12, optionsValue)
	stateSinkObject.Save(0, &s.segmentEntry)
	stateSinkObject.Save(1, &s.segmentRefs)
	stateSinkObject.Save(2, &s.ep)
	stateSinkObject.Save(3, &s.qFlags)
	stateSinkObject.Save(4, &s.pkt)
	stateSinkObject.Save(5, &s.sequenceNumber)
	stateSinkObject.Save(6, &s.ackNumber)
	stateSinkObject.Save(7, &s.flags)
	stateSinkObject.Save(8, &s.window)
	stateSinkObject.Save(9, &s.csum)
	stateSinkObject.Save(10, &s.csumValid)
	stateSinkObject.Save(11, &s.parsedOptions)
	stateSinkObject.Save(13, &s.hasNewSACKInfo)
	stateSinkObject.Save(14, &s.rcvdTime)
	stateSinkObject.Save(15, &s.xmitTime)
	stateSinkObject.Save(16, &s.xmitCount)
	stateSinkObject.Save(17, &s.acked)
	stateSinkObject.Save(18, &s.dataMemSize)
	stateSinkObject.Save(19, &s.lost)
}

func (s *segment) afterLoad() {}

// +checklocksignore
func (s *segment) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.segmentEntry)
	stateSourceObject.Load(1, &s.segmentRefs)
	stateSourceObject.Load(2, &s.ep)
	stateSourceObject.Load(3, &s.qFlags)
	stateSourceObject.Load(4, &s.pkt)
	stateSourceObject.Load(5, &s.sequenceNumber)
	stateSourceObject.Load(6, &s.ackNumber)
	stateSourceObject.Load(7, &s.flags)
	stateSourceObject.Load(8, &s.window)
	stateSourceObject.Load(9, &s.csum)
	stateSourceObject.Load(10, &s.csumValid)
	stateSourceObject.Load(11, &s.parsedOptions)
	stateSourceObject.Load(13, &s.hasNewSACKInfo)
	stateSourceObject.Load(14, &s.rcvdTime)
	stateSourceObject.Load(15, &s.xmitTime)
	stateSourceObject.Load(16, &s.xmitCount)
	stateSourceObject.Load(17, &s.acked)
	stateSourceObject.Load(18, &s.dataMemSize)
	stateSourceObject.Load(19, &s.lost)
	stateSourceObject.LoadValue(12, new([]byte), func(y any) { s.loadOptions(y.([]byte)) })
}

func (q *segmentQueue) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentQueue"
}

func (q *segmentQueue) StateFields() []string {
	return []string{
		"list",
		"ep",
		"frozen",
	}
}

func (q *segmentQueue) beforeSave() {}

// +checklocksignore
func (q *segmentQueue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.list)
	stateSinkObject.Save(1, &q.ep)
	stateSinkObject.Save(2, &q.frozen)
}

func (q *segmentQueue) afterLoad() {}

// +checklocksignore
func (q *segmentQueue) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadWait(0, &q.list)
	stateSourceObject.Load(1, &q.ep)
	stateSourceObject.Load(2, &q.frozen)
}

func (s *sender) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.sender"
}

func (s *sender) StateFields() []string {
	return []string{
		"TCPSenderState",
		"ep",
		"lr",
		"firstRetransmittedSegXmitTime",
		"writeNext",
		"writeList",
		"rtt",
		"minRTO",
		"maxRTO",
		"maxRetries",
		"gso",
		"state",
		"cc",
		"rc",
		"spuriousRecovery",
		"retransmitTS",
	}
}

func (s *sender) beforeSave() {}

// +checklocksignore
func (s *sender) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.TCPSenderState)
	stateSinkObject.Save(1, &s.ep)
	stateSinkObject.Save(2, &s.lr)
	stateSinkObject.Save(3, &s.firstRetransmittedSegXmitTime)
	stateSinkObject.Save(4, &s.writeNext)
	stateSinkObject.Save(5, &s.writeList)
	stateSinkObject.Save(6, &s.rtt)
	stateSinkObject.Save(7, &s.minRTO)
	stateSinkObject.Save(8, &s.maxRTO)
	stateSinkObject.Save(9, &s.maxRetries)
	stateSinkObject.Save(10, &s.gso)
	stateSinkObject.Save(11, &s.state)
	stateSinkObject.Save(12, &s.cc)
	stateSinkObject.Save(13, &s.rc)
	stateSinkObject.Save(14, &s.spuriousRecovery)
	stateSinkObject.Save(15, &s.retransmitTS)
}

func (s *sender) afterLoad() {}

// +checklocksignore
func (s *sender) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.TCPSenderState)
	stateSourceObject.Load(1, &s.ep)
	stateSourceObject.Load(2, &s.lr)
	stateSourceObject.Load(3, &s.firstRetransmittedSegXmitTime)
	stateSourceObject.Load(4, &s.writeNext)
	stateSourceObject.Load(5, &s.writeList)
	stateSourceObject.Load(6, &s.rtt)
	stateSourceObject.Load(7, &s.minRTO)
	stateSourceObject.Load(8, &s.maxRTO)
	stateSourceObject.Load(9, &s.maxRetries)
	stateSourceObject.Load(10, &s.gso)
	stateSourceObject.Load(11, &s.state)
	stateSourceObject.Load(12, &s.cc)
	stateSourceObject.Load(13, &s.rc)
	stateSourceObject.Load(14, &s.spuriousRecovery)
	stateSourceObject.Load(15, &s.retransmitTS)
}

func (r *rtt) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.rtt"
}

func (r *rtt) StateFields() []string {
	return []string{
		"TCPRTTState",
	}
}

func (r *rtt) beforeSave() {}

// +checklocksignore
func (r *rtt) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.TCPRTTState)
}

func (r *rtt) afterLoad() {}

// +checklocksignore
func (r *rtt) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.TCPRTTState)
}

func (l *endpointList) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.endpointList"
}

func (l *endpointList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *endpointList) beforeSave() {}

// +checklocksignore
func (l *endpointList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *endpointList) afterLoad() {}

// +checklocksignore
func (l *endpointList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *endpointEntry) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.endpointEntry"
}

func (e *endpointEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *endpointEntry) beforeSave() {}

// +checklocksignore
func (e *endpointEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *endpointEntry) afterLoad() {}

// +checklocksignore
func (e *endpointEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (l *segmentList) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentList"
}

func (l *segmentList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *segmentList) beforeSave() {}

// +checklocksignore
func (l *segmentList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *segmentList) afterLoad() {}

// +checklocksignore
func (l *segmentList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *segmentEntry) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentEntry"
}

func (e *segmentEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *segmentEntry) beforeSave() {}

// +checklocksignore
func (e *segmentEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *segmentEntry) afterLoad() {}

// +checklocksignore
func (e *segmentEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (r *segmentRefs) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentRefs"
}

func (r *segmentRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *segmentRefs) beforeSave() {}

// +checklocksignore
func (r *segmentRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *segmentRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*acceptQueue)(nil))
	state.Register((*handshake)(nil))
	state.Register((*cubicState)(nil))
	state.Register((*SACKInfo)(nil))
	state.Register((*ReceiveErrors)(nil))
	state.Register((*SendErrors)(nil))
	state.Register((*Stats)(nil))
	state.Register((*sndQueueInfo)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*keepalive)(nil))
	state.Register((*rackControl)(nil))
	state.Register((*receiver)(nil))
	state.Register((*renoState)(nil))
	state.Register((*renoRecovery)(nil))
	state.Register((*sackRecovery)(nil))
	state.Register((*SACKScoreboard)(nil))
	state.Register((*segment)(nil))
	state.Register((*segmentQueue)(nil))
	state.Register((*sender)(nil))
	state.Register((*rtt)(nil))
	state.Register((*endpointList)(nil))
	state.Register((*endpointEntry)(nil))
	state.Register((*segmentList)(nil))
	state.Register((*segmentEntry)(nil))
	state.Register((*segmentRefs)(nil))
}
