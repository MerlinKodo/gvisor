// automatically generated by stateify.

package pgalloc

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (r *EvictableRange) StateTypeName() string {
	return "pkg/sentry/pgalloc.EvictableRange"
}

func (r *EvictableRange) StateFields() []string {
	return []string{
		"Start",
		"End",
	}
}

func (r *EvictableRange) beforeSave() {}

// +checklocksignore
func (r *EvictableRange) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Start)
	stateSinkObject.Save(1, &r.End)
}

func (r *EvictableRange) afterLoad() {}

// +checklocksignore
func (r *EvictableRange) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Start)
	stateSourceObject.Load(1, &r.End)
}

func (s *evictableRangeSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.evictableRangeSet"
}

func (s *evictableRangeSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *evictableRangeSet) beforeSave() {}

// +checklocksignore
func (s *evictableRangeSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *evictableRangeSegmentDataSlices
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *evictableRangeSet) afterLoad() {}

// +checklocksignore
func (s *evictableRangeSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*evictableRangeSegmentDataSlices), func(y any) { s.loadRoot(y.(*evictableRangeSegmentDataSlices)) })
}

func (n *evictableRangenode) StateTypeName() string {
	return "pkg/sentry/pgalloc.evictableRangenode"
}

func (n *evictableRangenode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *evictableRangenode) beforeSave() {}

// +checklocksignore
func (n *evictableRangenode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *evictableRangenode) afterLoad() {}

// +checklocksignore
func (n *evictableRangenode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (e *evictableRangeSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/pgalloc.evictableRangeSegmentDataSlices"
}

func (e *evictableRangeSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (e *evictableRangeSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (e *evictableRangeSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Start)
	stateSinkObject.Save(1, &e.End)
	stateSinkObject.Save(2, &e.Values)
}

func (e *evictableRangeSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (e *evictableRangeSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Start)
	stateSourceObject.Load(1, &e.End)
	stateSourceObject.Load(2, &e.Values)
}

func (u *usageInfo) StateTypeName() string {
	return "pkg/sentry/pgalloc.usageInfo"
}

func (u *usageInfo) StateFields() []string {
	return []string{
		"kind",
		"knownCommitted",
		"refs",
		"memCgID",
	}
}

func (u *usageInfo) beforeSave() {}

// +checklocksignore
func (u *usageInfo) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.kind)
	stateSinkObject.Save(1, &u.knownCommitted)
	stateSinkObject.Save(2, &u.refs)
	stateSinkObject.Save(3, &u.memCgID)
}

func (u *usageInfo) afterLoad() {}

// +checklocksignore
func (u *usageInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.kind)
	stateSourceObject.Load(1, &u.knownCommitted)
	stateSourceObject.Load(2, &u.refs)
	stateSourceObject.Load(3, &u.memCgID)
}

func (s *reclaimSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.reclaimSet"
}

func (s *reclaimSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *reclaimSet) beforeSave() {}

// +checklocksignore
func (s *reclaimSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *reclaimSegmentDataSlices
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *reclaimSet) afterLoad() {}

// +checklocksignore
func (s *reclaimSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*reclaimSegmentDataSlices), func(y any) { s.loadRoot(y.(*reclaimSegmentDataSlices)) })
}

func (n *reclaimnode) StateTypeName() string {
	return "pkg/sentry/pgalloc.reclaimnode"
}

func (n *reclaimnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *reclaimnode) beforeSave() {}

// +checklocksignore
func (n *reclaimnode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *reclaimnode) afterLoad() {}

// +checklocksignore
func (n *reclaimnode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (r *reclaimSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/pgalloc.reclaimSegmentDataSlices"
}

func (r *reclaimSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (r *reclaimSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (r *reclaimSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Start)
	stateSinkObject.Save(1, &r.End)
	stateSinkObject.Save(2, &r.Values)
}

func (r *reclaimSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (r *reclaimSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Start)
	stateSourceObject.Load(1, &r.End)
	stateSourceObject.Load(2, &r.Values)
}

func (s *usageSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.usageSet"
}

func (s *usageSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *usageSet) beforeSave() {}

// +checklocksignore
func (s *usageSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *usageSegmentDataSlices
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *usageSet) afterLoad() {}

// +checklocksignore
func (s *usageSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*usageSegmentDataSlices), func(y any) { s.loadRoot(y.(*usageSegmentDataSlices)) })
}

func (n *usagenode) StateTypeName() string {
	return "pkg/sentry/pgalloc.usagenode"
}

func (n *usagenode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *usagenode) beforeSave() {}

// +checklocksignore
func (n *usagenode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *usagenode) afterLoad() {}

// +checklocksignore
func (n *usagenode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (u *usageSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/pgalloc.usageSegmentDataSlices"
}

func (u *usageSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (u *usageSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (u *usageSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.Start)
	stateSinkObject.Save(1, &u.End)
	stateSinkObject.Save(2, &u.Values)
}

func (u *usageSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (u *usageSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.Start)
	stateSourceObject.Load(1, &u.End)
	stateSourceObject.Load(2, &u.Values)
}

func init() {
	state.Register((*EvictableRange)(nil))
	state.Register((*evictableRangeSet)(nil))
	state.Register((*evictableRangenode)(nil))
	state.Register((*evictableRangeSegmentDataSlices)(nil))
	state.Register((*usageInfo)(nil))
	state.Register((*reclaimSet)(nil))
	state.Register((*reclaimnode)(nil))
	state.Register((*reclaimSegmentDataSlices)(nil))
	state.Register((*usageSet)(nil))
	state.Register((*usagenode)(nil))
	state.Register((*usageSegmentDataSlices)(nil))
}
