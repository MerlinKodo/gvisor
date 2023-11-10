// automatically generated by stateify.

//go:build arm64
// +build arm64

package arch

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (r *Registers) StateTypeName() string {
	return "pkg/sentry/arch.Registers"
}

func (r *Registers) StateFields() []string {
	return []string{
		"PtraceRegs",
		"TPIDR_EL0",
	}
}

func (r *Registers) beforeSave() {}

// +checklocksignore
func (r *Registers) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.PtraceRegs)
	stateSinkObject.Save(1, &r.TPIDR_EL0)
}

func (r *Registers) afterLoad() {}

// +checklocksignore
func (r *Registers) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.PtraceRegs)
	stateSourceObject.Load(1, &r.TPIDR_EL0)
}

func (s *State) StateTypeName() string {
	return "pkg/sentry/arch.State"
}

func (s *State) StateFields() []string {
	return []string{
		"Regs",
		"fpState",
		"OrigR0",
	}
}

func (s *State) beforeSave() {}

// +checklocksignore
func (s *State) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Regs)
	stateSinkObject.Save(1, &s.fpState)
	stateSinkObject.Save(2, &s.OrigR0)
}

func (s *State) afterLoad() {}

// +checklocksignore
func (s *State) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Regs)
	stateSourceObject.LoadWait(1, &s.fpState)
	stateSourceObject.Load(2, &s.OrigR0)
}

func init() {
	state.Register((*Registers)(nil))
	state.Register((*State)(nil))
}
