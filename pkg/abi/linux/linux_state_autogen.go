// automatically generated by stateify.

package linux

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (i *IOEvent) StateTypeName() string {
	return "pkg/abi/linux.IOEvent"
}

func (i *IOEvent) StateFields() []string {
	return []string{
		"Data",
		"Obj",
		"Result",
		"Result2",
	}
}

func (i *IOEvent) beforeSave() {}

// +checklocksignore
func (i *IOEvent) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Data)
	stateSinkObject.Save(1, &i.Obj)
	stateSinkObject.Save(2, &i.Result)
	stateSinkObject.Save(3, &i.Result2)
}

func (i *IOEvent) afterLoad() {}

// +checklocksignore
func (i *IOEvent) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Data)
	stateSourceObject.Load(1, &i.Obj)
	stateSourceObject.Load(2, &i.Result)
	stateSourceObject.Load(3, &i.Result2)
}

func (b *BPFInstruction) StateTypeName() string {
	return "pkg/abi/linux.BPFInstruction"
}

func (b *BPFInstruction) StateFields() []string {
	return []string{
		"OpCode",
		"JumpIfTrue",
		"JumpIfFalse",
		"K",
	}
}

func (b *BPFInstruction) beforeSave() {}

// +checklocksignore
func (b *BPFInstruction) StateSave(stateSinkObject state.Sink) {
	b.beforeSave()
	stateSinkObject.Save(0, &b.OpCode)
	stateSinkObject.Save(1, &b.JumpIfTrue)
	stateSinkObject.Save(2, &b.JumpIfFalse)
	stateSinkObject.Save(3, &b.K)
}

func (b *BPFInstruction) afterLoad() {}

// +checklocksignore
func (b *BPFInstruction) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &b.OpCode)
	stateSourceObject.Load(1, &b.JumpIfTrue)
	stateSourceObject.Load(2, &b.JumpIfFalse)
	stateSourceObject.Load(3, &b.K)
}

func (c *ControlMessageIPPacketInfo) StateTypeName() string {
	return "pkg/abi/linux.ControlMessageIPPacketInfo"
}

func (c *ControlMessageIPPacketInfo) StateFields() []string {
	return []string{
		"NIC",
		"LocalAddr",
		"DestinationAddr",
	}
}

func (c *ControlMessageIPPacketInfo) beforeSave() {}

// +checklocksignore
func (c *ControlMessageIPPacketInfo) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.NIC)
	stateSinkObject.Save(1, &c.LocalAddr)
	stateSinkObject.Save(2, &c.DestinationAddr)
}

func (c *ControlMessageIPPacketInfo) afterLoad() {}

// +checklocksignore
func (c *ControlMessageIPPacketInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.NIC)
	stateSourceObject.Load(1, &c.LocalAddr)
	stateSourceObject.Load(2, &c.DestinationAddr)
}

func (t *KernelTermios) StateTypeName() string {
	return "pkg/abi/linux.KernelTermios"
}

func (t *KernelTermios) StateFields() []string {
	return []string{
		"InputFlags",
		"OutputFlags",
		"ControlFlags",
		"LocalFlags",
		"LineDiscipline",
		"ControlCharacters",
		"InputSpeed",
		"OutputSpeed",
	}
}

func (t *KernelTermios) beforeSave() {}

// +checklocksignore
func (t *KernelTermios) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.InputFlags)
	stateSinkObject.Save(1, &t.OutputFlags)
	stateSinkObject.Save(2, &t.ControlFlags)
	stateSinkObject.Save(3, &t.LocalFlags)
	stateSinkObject.Save(4, &t.LineDiscipline)
	stateSinkObject.Save(5, &t.ControlCharacters)
	stateSinkObject.Save(6, &t.InputSpeed)
	stateSinkObject.Save(7, &t.OutputSpeed)
}

func (t *KernelTermios) afterLoad() {}

// +checklocksignore
func (t *KernelTermios) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.InputFlags)
	stateSourceObject.Load(1, &t.OutputFlags)
	stateSourceObject.Load(2, &t.ControlFlags)
	stateSourceObject.Load(3, &t.LocalFlags)
	stateSourceObject.Load(4, &t.LineDiscipline)
	stateSourceObject.Load(5, &t.ControlCharacters)
	stateSourceObject.Load(6, &t.InputSpeed)
	stateSourceObject.Load(7, &t.OutputSpeed)
}

func (w *WindowSize) StateTypeName() string {
	return "pkg/abi/linux.WindowSize"
}

func (w *WindowSize) StateFields() []string {
	return []string{
		"Rows",
		"Cols",
	}
}

func (w *WindowSize) beforeSave() {}

// +checklocksignore
func (w *WindowSize) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.Rows)
	stateSinkObject.Save(1, &w.Cols)
}

func (w *WindowSize) afterLoad() {}

// +checklocksignore
func (w *WindowSize) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.Rows)
	stateSourceObject.Load(1, &w.Cols)
}

func init() {
	state.Register((*IOEvent)(nil))
	state.Register((*BPFInstruction)(nil))
	state.Register((*ControlMessageIPPacketInfo)(nil))
	state.Register((*KernelTermios)(nil))
	state.Register((*WindowSize)(nil))
}