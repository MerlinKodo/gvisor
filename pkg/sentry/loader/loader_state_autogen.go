// automatically generated by stateify.

package loader

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (v *VDSO) StateTypeName() string {
	return "pkg/sentry/loader.VDSO"
}

func (v *VDSO) StateFields() []string {
	return []string{
		"ParamPage",
		"vdso",
		"os",
		"arch",
		"phdrs",
	}
}

func (v *VDSO) beforeSave() {}

// +checklocksignore
func (v *VDSO) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	var phdrsValue []elfProgHeader
	phdrsValue = v.savePhdrs()
	stateSinkObject.SaveValue(4, phdrsValue)
	stateSinkObject.Save(0, &v.ParamPage)
	stateSinkObject.Save(1, &v.vdso)
	stateSinkObject.Save(2, &v.os)
	stateSinkObject.Save(3, &v.arch)
}

func (v *VDSO) afterLoad() {}

// +checklocksignore
func (v *VDSO) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.ParamPage)
	stateSourceObject.Load(1, &v.vdso)
	stateSourceObject.Load(2, &v.os)
	stateSourceObject.Load(3, &v.arch)
	stateSourceObject.LoadValue(4, new([]elfProgHeader), func(y any) { v.loadPhdrs(y.([]elfProgHeader)) })
}

func (e *elfProgHeader) StateTypeName() string {
	return "pkg/sentry/loader.elfProgHeader"
}

func (e *elfProgHeader) StateFields() []string {
	return []string{
		"Type",
		"Flags",
		"Off",
		"Vaddr",
		"Paddr",
		"Filesz",
		"Memsz",
		"Align",
	}
}

func (e *elfProgHeader) beforeSave() {}

// +checklocksignore
func (e *elfProgHeader) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Type)
	stateSinkObject.Save(1, &e.Flags)
	stateSinkObject.Save(2, &e.Off)
	stateSinkObject.Save(3, &e.Vaddr)
	stateSinkObject.Save(4, &e.Paddr)
	stateSinkObject.Save(5, &e.Filesz)
	stateSinkObject.Save(6, &e.Memsz)
	stateSinkObject.Save(7, &e.Align)
}

func (e *elfProgHeader) afterLoad() {}

// +checklocksignore
func (e *elfProgHeader) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Type)
	stateSourceObject.Load(1, &e.Flags)
	stateSourceObject.Load(2, &e.Off)
	stateSourceObject.Load(3, &e.Vaddr)
	stateSourceObject.Load(4, &e.Paddr)
	stateSourceObject.Load(5, &e.Filesz)
	stateSourceObject.Load(6, &e.Memsz)
	stateSourceObject.Load(7, &e.Align)
}

func init() {
	state.Register((*VDSO)(nil))
	state.Register((*elfProgHeader)(nil))
}
