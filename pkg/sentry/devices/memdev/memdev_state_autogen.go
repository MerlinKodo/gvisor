// automatically generated by stateify.

package memdev

import (
	"github.com/MerlinKodo/gvisor/pkg/state"
)

func (f *fullDevice) StateTypeName() string {
	return "pkg/sentry/devices/memdev.fullDevice"
}

func (f *fullDevice) StateFields() []string {
	return []string{}
}

func (f *fullDevice) beforeSave() {}

// +checklocksignore
func (f *fullDevice) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *fullDevice) afterLoad() {}

// +checklocksignore
func (f *fullDevice) StateLoad(stateSourceObject state.Source) {
}

func (fd *fullFD) StateTypeName() string {
	return "pkg/sentry/devices/memdev.fullFD"
}

func (fd *fullFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
	}
}

func (fd *fullFD) beforeSave() {}

// +checklocksignore
func (fd *fullFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
}

func (fd *fullFD) afterLoad() {}

// +checklocksignore
func (fd *fullFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
}

func (n *nullDevice) StateTypeName() string {
	return "pkg/sentry/devices/memdev.nullDevice"
}

func (n *nullDevice) StateFields() []string {
	return []string{}
}

func (n *nullDevice) beforeSave() {}

// +checklocksignore
func (n *nullDevice) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
}

func (n *nullDevice) afterLoad() {}

// +checklocksignore
func (n *nullDevice) StateLoad(stateSourceObject state.Source) {
}

func (fd *nullFD) StateTypeName() string {
	return "pkg/sentry/devices/memdev.nullFD"
}

func (fd *nullFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
	}
}

func (fd *nullFD) beforeSave() {}

// +checklocksignore
func (fd *nullFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
}

func (fd *nullFD) afterLoad() {}

// +checklocksignore
func (fd *nullFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
}

func (r *randomDevice) StateTypeName() string {
	return "pkg/sentry/devices/memdev.randomDevice"
}

func (r *randomDevice) StateFields() []string {
	return []string{}
}

func (r *randomDevice) beforeSave() {}

// +checklocksignore
func (r *randomDevice) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
}

func (r *randomDevice) afterLoad() {}

// +checklocksignore
func (r *randomDevice) StateLoad(stateSourceObject state.Source) {
}

func (fd *randomFD) StateTypeName() string {
	return "pkg/sentry/devices/memdev.randomFD"
}

func (fd *randomFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"off",
	}
}

func (fd *randomFD) beforeSave() {}

// +checklocksignore
func (fd *randomFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
	stateSinkObject.Save(4, &fd.off)
}

func (fd *randomFD) afterLoad() {}

// +checklocksignore
func (fd *randomFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
	stateSourceObject.Load(4, &fd.off)
}

func (z *zeroDevice) StateTypeName() string {
	return "pkg/sentry/devices/memdev.zeroDevice"
}

func (z *zeroDevice) StateFields() []string {
	return []string{}
}

func (z *zeroDevice) beforeSave() {}

// +checklocksignore
func (z *zeroDevice) StateSave(stateSinkObject state.Sink) {
	z.beforeSave()
}

func (z *zeroDevice) afterLoad() {}

// +checklocksignore
func (z *zeroDevice) StateLoad(stateSourceObject state.Source) {
}

func (fd *zeroFD) StateTypeName() string {
	return "pkg/sentry/devices/memdev.zeroFD"
}

func (fd *zeroFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
	}
}

func (fd *zeroFD) beforeSave() {}

// +checklocksignore
func (fd *zeroFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
}

func (fd *zeroFD) afterLoad() {}

// +checklocksignore
func (fd *zeroFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
}

func init() {
	state.Register((*fullDevice)(nil))
	state.Register((*fullFD)(nil))
	state.Register((*nullDevice)(nil))
	state.Register((*nullFD)(nil))
	state.Register((*randomDevice)(nil))
	state.Register((*randomFD)(nil))
	state.Register((*zeroDevice)(nil))
	state.Register((*zeroFD)(nil))
}
