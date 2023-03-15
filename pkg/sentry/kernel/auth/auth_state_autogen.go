// automatically generated by stateify.

package auth

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (c *Credentials) StateTypeName() string {
	return "pkg/sentry/kernel/auth.Credentials"
}

func (c *Credentials) StateFields() []string {
	return []string{
		"RealKUID",
		"EffectiveKUID",
		"SavedKUID",
		"RealKGID",
		"EffectiveKGID",
		"SavedKGID",
		"ExtraKGIDs",
		"PermittedCaps",
		"InheritableCaps",
		"EffectiveCaps",
		"BoundingCaps",
		"KeepCaps",
		"UserNamespace",
	}
}

func (c *Credentials) beforeSave() {}

// +checklocksignore
func (c *Credentials) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.RealKUID)
	stateSinkObject.Save(1, &c.EffectiveKUID)
	stateSinkObject.Save(2, &c.SavedKUID)
	stateSinkObject.Save(3, &c.RealKGID)
	stateSinkObject.Save(4, &c.EffectiveKGID)
	stateSinkObject.Save(5, &c.SavedKGID)
	stateSinkObject.Save(6, &c.ExtraKGIDs)
	stateSinkObject.Save(7, &c.PermittedCaps)
	stateSinkObject.Save(8, &c.InheritableCaps)
	stateSinkObject.Save(9, &c.EffectiveCaps)
	stateSinkObject.Save(10, &c.BoundingCaps)
	stateSinkObject.Save(11, &c.KeepCaps)
	stateSinkObject.Save(12, &c.UserNamespace)
}

func (c *Credentials) afterLoad() {}

// +checklocksignore
func (c *Credentials) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.RealKUID)
	stateSourceObject.Load(1, &c.EffectiveKUID)
	stateSourceObject.Load(2, &c.SavedKUID)
	stateSourceObject.Load(3, &c.RealKGID)
	stateSourceObject.Load(4, &c.EffectiveKGID)
	stateSourceObject.Load(5, &c.SavedKGID)
	stateSourceObject.Load(6, &c.ExtraKGIDs)
	stateSourceObject.Load(7, &c.PermittedCaps)
	stateSourceObject.Load(8, &c.InheritableCaps)
	stateSourceObject.Load(9, &c.EffectiveCaps)
	stateSourceObject.Load(10, &c.BoundingCaps)
	stateSourceObject.Load(11, &c.KeepCaps)
	stateSourceObject.Load(12, &c.UserNamespace)
}

func (i *IDMapEntry) StateTypeName() string {
	return "pkg/sentry/kernel/auth.IDMapEntry"
}

func (i *IDMapEntry) StateFields() []string {
	return []string{
		"FirstID",
		"FirstParentID",
		"Length",
	}
}

func (i *IDMapEntry) beforeSave() {}

// +checklocksignore
func (i *IDMapEntry) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.FirstID)
	stateSinkObject.Save(1, &i.FirstParentID)
	stateSinkObject.Save(2, &i.Length)
}

func (i *IDMapEntry) afterLoad() {}

// +checklocksignore
func (i *IDMapEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.FirstID)
	stateSourceObject.Load(1, &i.FirstParentID)
	stateSourceObject.Load(2, &i.Length)
}

func (r *idMapRange) StateTypeName() string {
	return "pkg/sentry/kernel/auth.idMapRange"
}

func (r *idMapRange) StateFields() []string {
	return []string{
		"Start",
		"End",
	}
}

func (r *idMapRange) beforeSave() {}

// +checklocksignore
func (r *idMapRange) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Start)
	stateSinkObject.Save(1, &r.End)
}

func (r *idMapRange) afterLoad() {}

// +checklocksignore
func (r *idMapRange) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Start)
	stateSourceObject.Load(1, &r.End)
}

func (s *idMapSet) StateTypeName() string {
	return "pkg/sentry/kernel/auth.idMapSet"
}

func (s *idMapSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *idMapSet) beforeSave() {}

// +checklocksignore
func (s *idMapSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *idMapSegmentDataSlices
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *idMapSet) afterLoad() {}

// +checklocksignore
func (s *idMapSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*idMapSegmentDataSlices), func(y any) { s.loadRoot(y.(*idMapSegmentDataSlices)) })
}

func (n *idMapnode) StateTypeName() string {
	return "pkg/sentry/kernel/auth.idMapnode"
}

func (n *idMapnode) StateFields() []string {
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

func (n *idMapnode) beforeSave() {}

// +checklocksignore
func (n *idMapnode) StateSave(stateSinkObject state.Sink) {
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

func (n *idMapnode) afterLoad() {}

// +checklocksignore
func (n *idMapnode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (i *idMapSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/kernel/auth.idMapSegmentDataSlices"
}

func (i *idMapSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (i *idMapSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (i *idMapSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Start)
	stateSinkObject.Save(1, &i.End)
	stateSinkObject.Save(2, &i.Values)
}

func (i *idMapSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (i *idMapSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Start)
	stateSourceObject.Load(1, &i.End)
	stateSourceObject.Load(2, &i.Values)
}

func (ns *UserNamespace) StateTypeName() string {
	return "pkg/sentry/kernel/auth.UserNamespace"
}

func (ns *UserNamespace) StateFields() []string {
	return []string{
		"parent",
		"owner",
		"uidMapFromParent",
		"uidMapToParent",
		"gidMapFromParent",
		"gidMapToParent",
	}
}

func (ns *UserNamespace) beforeSave() {}

// +checklocksignore
func (ns *UserNamespace) StateSave(stateSinkObject state.Sink) {
	ns.beforeSave()
	stateSinkObject.Save(0, &ns.parent)
	stateSinkObject.Save(1, &ns.owner)
	stateSinkObject.Save(2, &ns.uidMapFromParent)
	stateSinkObject.Save(3, &ns.uidMapToParent)
	stateSinkObject.Save(4, &ns.gidMapFromParent)
	stateSinkObject.Save(5, &ns.gidMapToParent)
}

func (ns *UserNamespace) afterLoad() {}

// +checklocksignore
func (ns *UserNamespace) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ns.parent)
	stateSourceObject.Load(1, &ns.owner)
	stateSourceObject.Load(2, &ns.uidMapFromParent)
	stateSourceObject.Load(3, &ns.uidMapToParent)
	stateSourceObject.Load(4, &ns.gidMapFromParent)
	stateSourceObject.Load(5, &ns.gidMapToParent)
}

func init() {
	state.Register((*Credentials)(nil))
	state.Register((*IDMapEntry)(nil))
	state.Register((*idMapRange)(nil))
	state.Register((*idMapSet)(nil))
	state.Register((*idMapnode)(nil))
	state.Register((*idMapSegmentDataSlices)(nil))
	state.Register((*UserNamespace)(nil))
}
