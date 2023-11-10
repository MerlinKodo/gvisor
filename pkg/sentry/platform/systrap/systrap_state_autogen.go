// automatically generated by stateify.

package systrap

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (l *contextList) StateTypeName() string {
	return "pkg/sentry/platform/systrap.contextList"
}

func (l *contextList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *contextList) beforeSave() {}

// +checklocksignore
func (l *contextList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *contextList) afterLoad() {}

// +checklocksignore
func (l *contextList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *contextEntry) StateTypeName() string {
	return "pkg/sentry/platform/systrap.contextEntry"
}

func (e *contextEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *contextEntry) beforeSave() {}

// +checklocksignore
func (e *contextEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *contextEntry) afterLoad() {}

// +checklocksignore
func (e *contextEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (r *subprocessRefs) StateTypeName() string {
	return "pkg/sentry/platform/systrap.subprocessRefs"
}

func (r *subprocessRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *subprocessRefs) beforeSave() {}

// +checklocksignore
func (r *subprocessRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *subprocessRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*contextList)(nil))
	state.Register((*contextEntry)(nil))
	state.Register((*subprocessRefs)(nil))
}
