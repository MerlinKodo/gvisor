// automatically generated by stateify.

package buffer

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (b *Buffer) StateTypeName() string {
	return "pkg/buffer.Buffer"
}

func (b *Buffer) StateFields() []string {
	return []string{
		"data",
		"size",
	}
}

func (b *Buffer) beforeSave() {}

// +checklocksignore
func (b *Buffer) StateSave(stateSinkObject state.Sink) {
	b.beforeSave()
	var dataValue []byte
	dataValue = b.saveData()
	stateSinkObject.SaveValue(0, dataValue)
	stateSinkObject.Save(1, &b.size)
}

func (b *Buffer) afterLoad() {}

// +checklocksignore
func (b *Buffer) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(1, &b.size)
	stateSourceObject.LoadValue(0, new([]byte), func(y any) { b.loadData(y.([]byte)) })
}

func (c *chunk) StateTypeName() string {
	return "pkg/buffer.chunk"
}

func (c *chunk) StateFields() []string {
	return []string{
		"chunkRefs",
		"data",
	}
}

func (c *chunk) beforeSave() {}

// +checklocksignore
func (c *chunk) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.chunkRefs)
	stateSinkObject.Save(1, &c.data)
}

func (c *chunk) afterLoad() {}

// +checklocksignore
func (c *chunk) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.chunkRefs)
	stateSourceObject.Load(1, &c.data)
}

func (r *chunkRefs) StateTypeName() string {
	return "pkg/buffer.chunkRefs"
}

func (r *chunkRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *chunkRefs) beforeSave() {}

// +checklocksignore
func (r *chunkRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *chunkRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (v *View) StateTypeName() string {
	return "pkg/buffer.View"
}

func (v *View) StateFields() []string {
	return []string{
		"read",
		"write",
		"chunk",
	}
}

func (v *View) beforeSave() {}

// +checklocksignore
func (v *View) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.read)
	stateSinkObject.Save(1, &v.write)
	stateSinkObject.Save(2, &v.chunk)
}

func (v *View) afterLoad() {}

// +checklocksignore
func (v *View) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.read)
	stateSourceObject.Load(1, &v.write)
	stateSourceObject.Load(2, &v.chunk)
}

func (l *viewList) StateTypeName() string {
	return "pkg/buffer.viewList"
}

func (l *viewList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *viewList) beforeSave() {}

// +checklocksignore
func (l *viewList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *viewList) afterLoad() {}

// +checklocksignore
func (l *viewList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *viewEntry) StateTypeName() string {
	return "pkg/buffer.viewEntry"
}

func (e *viewEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *viewEntry) beforeSave() {}

// +checklocksignore
func (e *viewEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *viewEntry) afterLoad() {}

// +checklocksignore
func (e *viewEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*Buffer)(nil))
	state.Register((*chunk)(nil))
	state.Register((*chunkRefs)(nil))
	state.Register((*View)(nil))
	state.Register((*viewList)(nil))
	state.Register((*viewEntry)(nil))
}
