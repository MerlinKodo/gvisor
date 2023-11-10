// automatically generated by stateify.

package semaphore

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (r *Registry) StateTypeName() string {
	return "pkg/sentry/kernel/semaphore.Registry"
}

func (r *Registry) StateFields() []string {
	return []string{
		"reg",
		"indexes",
	}
}

func (r *Registry) beforeSave() {}

// +checklocksignore
func (r *Registry) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.reg)
	stateSinkObject.Save(1, &r.indexes)
}

func (r *Registry) afterLoad() {}

// +checklocksignore
func (r *Registry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.reg)
	stateSourceObject.Load(1, &r.indexes)
}

func (s *Set) StateTypeName() string {
	return "pkg/sentry/kernel/semaphore.Set"
}

func (s *Set) StateFields() []string {
	return []string{
		"registry",
		"obj",
		"opTime",
		"changeTime",
		"sems",
		"dead",
	}
}

func (s *Set) beforeSave() {}

// +checklocksignore
func (s *Set) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.registry)
	stateSinkObject.Save(1, &s.obj)
	stateSinkObject.Save(2, &s.opTime)
	stateSinkObject.Save(3, &s.changeTime)
	stateSinkObject.Save(4, &s.sems)
	stateSinkObject.Save(5, &s.dead)
}

func (s *Set) afterLoad() {}

// +checklocksignore
func (s *Set) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.registry)
	stateSourceObject.Load(1, &s.obj)
	stateSourceObject.Load(2, &s.opTime)
	stateSourceObject.Load(3, &s.changeTime)
	stateSourceObject.Load(4, &s.sems)
	stateSourceObject.Load(5, &s.dead)
}

func (s *sem) StateTypeName() string {
	return "pkg/sentry/kernel/semaphore.sem"
}

func (s *sem) StateFields() []string {
	return []string{
		"value",
		"pid",
	}
}

func (s *sem) beforeSave() {}

// +checklocksignore
func (s *sem) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	if !state.IsZeroValue(&s.waiters) {
		state.Failf("waiters is %#v, expected zero", &s.waiters)
	}
	stateSinkObject.Save(0, &s.value)
	stateSinkObject.Save(1, &s.pid)
}

func (s *sem) afterLoad() {}

// +checklocksignore
func (s *sem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.value)
	stateSourceObject.Load(1, &s.pid)
}

func (w *waiter) StateTypeName() string {
	return "pkg/sentry/kernel/semaphore.waiter"
}

func (w *waiter) StateFields() []string {
	return []string{
		"waiterEntry",
		"value",
		"ch",
	}
}

func (w *waiter) beforeSave() {}

// +checklocksignore
func (w *waiter) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.waiterEntry)
	stateSinkObject.Save(1, &w.value)
	stateSinkObject.Save(2, &w.ch)
}

func (w *waiter) afterLoad() {}

// +checklocksignore
func (w *waiter) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.waiterEntry)
	stateSourceObject.Load(1, &w.value)
	stateSourceObject.Load(2, &w.ch)
}

func (l *waiterList) StateTypeName() string {
	return "pkg/sentry/kernel/semaphore.waiterList"
}

func (l *waiterList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *waiterList) beforeSave() {}

// +checklocksignore
func (l *waiterList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *waiterList) afterLoad() {}

// +checklocksignore
func (l *waiterList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *waiterEntry) StateTypeName() string {
	return "pkg/sentry/kernel/semaphore.waiterEntry"
}

func (e *waiterEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *waiterEntry) beforeSave() {}

// +checklocksignore
func (e *waiterEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *waiterEntry) afterLoad() {}

// +checklocksignore
func (e *waiterEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*Registry)(nil))
	state.Register((*Set)(nil))
	state.Register((*sem)(nil))
	state.Register((*waiter)(nil))
	state.Register((*waiterList)(nil))
	state.Register((*waiterEntry)(nil))
}
