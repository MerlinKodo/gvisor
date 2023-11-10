// automatically generated by stateify.

package eventfd

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (efd *EventFileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/eventfd.EventFileDescription"
}

func (efd *EventFileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"queue",
		"val",
		"semMode",
		"hostfd",
	}
}

func (efd *EventFileDescription) beforeSave() {}

// +checklocksignore
func (efd *EventFileDescription) StateSave(stateSinkObject state.Sink) {
	efd.beforeSave()
	stateSinkObject.Save(0, &efd.vfsfd)
	stateSinkObject.Save(1, &efd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &efd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &efd.NoLockFD)
	stateSinkObject.Save(4, &efd.queue)
	stateSinkObject.Save(5, &efd.val)
	stateSinkObject.Save(6, &efd.semMode)
	stateSinkObject.Save(7, &efd.hostfd)
}

func (efd *EventFileDescription) afterLoad() {}

// +checklocksignore
func (efd *EventFileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &efd.vfsfd)
	stateSourceObject.Load(1, &efd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &efd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &efd.NoLockFD)
	stateSourceObject.Load(4, &efd.queue)
	stateSourceObject.Load(5, &efd.val)
	stateSourceObject.Load(6, &efd.semMode)
	stateSourceObject.Load(7, &efd.hostfd)
}

func init() {
	state.Register((*EventFileDescription)(nil))
}
