// automatically generated by stateify.

package tmpfs

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (l *dentryList) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.dentryList"
}

func (l *dentryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *dentryList) beforeSave() {}

// +checklocksignore
func (l *dentryList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *dentryList) afterLoad() {}

// +checklocksignore
func (l *dentryList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *dentryEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.dentryEntry"
}

func (e *dentryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *dentryEntry) beforeSave() {}

// +checklocksignore
func (e *dentryEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *dentryEntry) afterLoad() {}

// +checklocksignore
func (e *dentryEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (d *deviceFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.deviceFile"
}

func (d *deviceFile) StateFields() []string {
	return []string{
		"inode",
		"kind",
		"major",
		"minor",
	}
}

func (d *deviceFile) beforeSave() {}

// +checklocksignore
func (d *deviceFile) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.inode)
	stateSinkObject.Save(1, &d.kind)
	stateSinkObject.Save(2, &d.major)
	stateSinkObject.Save(3, &d.minor)
}

func (d *deviceFile) afterLoad() {}

// +checklocksignore
func (d *deviceFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.inode)
	stateSourceObject.Load(1, &d.kind)
	stateSourceObject.Load(2, &d.major)
	stateSourceObject.Load(3, &d.minor)
}

func (dir *directory) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.directory"
}

func (dir *directory) StateFields() []string {
	return []string{
		"dentry",
		"inode",
		"childMap",
		"numChildren",
		"childList",
	}
}

func (dir *directory) beforeSave() {}

// +checklocksignore
func (dir *directory) StateSave(stateSinkObject state.Sink) {
	dir.beforeSave()
	stateSinkObject.Save(0, &dir.dentry)
	stateSinkObject.Save(1, &dir.inode)
	stateSinkObject.Save(2, &dir.childMap)
	stateSinkObject.Save(3, &dir.numChildren)
	stateSinkObject.Save(4, &dir.childList)
}

func (dir *directory) afterLoad() {}

// +checklocksignore
func (dir *directory) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dir.dentry)
	stateSourceObject.Load(1, &dir.inode)
	stateSourceObject.Load(2, &dir.childMap)
	stateSourceObject.Load(3, &dir.numChildren)
	stateSourceObject.Load(4, &dir.childList)
}

func (fd *directoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.directoryFD"
}

func (fd *directoryFD) StateFields() []string {
	return []string{
		"fileDescription",
		"DirectoryFileDescriptionDefaultImpl",
		"iter",
		"off",
	}
}

func (fd *directoryFD) beforeSave() {}

// +checklocksignore
func (fd *directoryFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.iter)
	stateSinkObject.Save(3, &fd.off)
}

func (fd *directoryFD) afterLoad() {}

// +checklocksignore
func (fd *directoryFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.iter)
	stateSourceObject.Load(3, &fd.off)
}

func (r *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.inodeRefs"
}

func (r *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *inodeRefs) beforeSave() {}

// +checklocksignore
func (r *inodeRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *inodeRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (n *namedPipe) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.namedPipe"
}

func (n *namedPipe) StateFields() []string {
	return []string{
		"inode",
		"pipe",
	}
}

func (n *namedPipe) beforeSave() {}

// +checklocksignore
func (n *namedPipe) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.inode)
	stateSinkObject.Save(1, &n.pipe)
}

func (n *namedPipe) afterLoad() {}

// +checklocksignore
func (n *namedPipe) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.inode)
	stateSourceObject.Load(1, &n.pipe)
}

func (rf *regularFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.regularFile"
}

func (rf *regularFile) StateFields() []string {
	return []string{
		"inode",
		"memoryUsageKind",
		"mappings",
		"writableMappingPages",
		"data",
		"seals",
		"size",
	}
}

func (rf *regularFile) beforeSave() {}

// +checklocksignore
func (rf *regularFile) StateSave(stateSinkObject state.Sink) {
	rf.beforeSave()
	stateSinkObject.Save(0, &rf.inode)
	stateSinkObject.Save(1, &rf.memoryUsageKind)
	stateSinkObject.Save(2, &rf.mappings)
	stateSinkObject.Save(3, &rf.writableMappingPages)
	stateSinkObject.Save(4, &rf.data)
	stateSinkObject.Save(5, &rf.seals)
	stateSinkObject.Save(6, &rf.size)
}

func (rf *regularFile) afterLoad() {}

// +checklocksignore
func (rf *regularFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rf.inode)
	stateSourceObject.Load(1, &rf.memoryUsageKind)
	stateSourceObject.Load(2, &rf.mappings)
	stateSourceObject.Load(3, &rf.writableMappingPages)
	stateSourceObject.Load(4, &rf.data)
	stateSourceObject.Load(5, &rf.seals)
	stateSourceObject.Load(6, &rf.size)
}

func (fd *regularFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.regularFileFD"
}

func (fd *regularFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"off",
	}
}

func (fd *regularFileFD) beforeSave() {}

// +checklocksignore
func (fd *regularFileFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.off)
}

func (fd *regularFileFD) afterLoad() {}

// +checklocksignore
func (fd *regularFileFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.off)
}

func (s *socketFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.socketFile"
}

func (s *socketFile) StateFields() []string {
	return []string{
		"inode",
		"ep",
	}
}

func (s *socketFile) beforeSave() {}

// +checklocksignore
func (s *socketFile) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.inode)
	stateSinkObject.Save(1, &s.ep)
}

func (s *socketFile) afterLoad() {}

// +checklocksignore
func (s *socketFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.inode)
	stateSourceObject.Load(1, &s.ep)
}

func (s *symlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.symlink"
}

func (s *symlink) StateFields() []string {
	return []string{
		"inode",
		"target",
	}
}

func (s *symlink) beforeSave() {}

// +checklocksignore
func (s *symlink) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.inode)
	stateSinkObject.Save(1, &s.target)
}

func (s *symlink) afterLoad() {}

// +checklocksignore
func (s *symlink) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.inode)
	stateSourceObject.Load(1, &s.target)
}

func (fstype *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.FilesystemType"
}

func (fstype *FilesystemType) StateFields() []string {
	return []string{}
}

func (fstype *FilesystemType) beforeSave() {}

// +checklocksignore
func (fstype *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fstype.beforeSave()
}

func (fstype *FilesystemType) afterLoad() {}

// +checklocksignore
func (fstype *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"privateMF",
		"mfp",
		"clock",
		"devMinor",
		"mopts",
		"usage",
		"nextInoMinusOne",
		"root",
		"maxFilenameLen",
		"maxSizeInPages",
		"pagesUsed",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.vfsfs)
	stateSinkObject.Save(1, &fs.privateMF)
	stateSinkObject.Save(2, &fs.mfp)
	stateSinkObject.Save(3, &fs.clock)
	stateSinkObject.Save(4, &fs.devMinor)
	stateSinkObject.Save(5, &fs.mopts)
	stateSinkObject.Save(6, &fs.usage)
	stateSinkObject.Save(7, &fs.nextInoMinusOne)
	stateSinkObject.Save(8, &fs.root)
	stateSinkObject.Save(9, &fs.maxFilenameLen)
	stateSinkObject.Save(10, &fs.maxSizeInPages)
	stateSinkObject.Save(11, &fs.pagesUsed)
}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.vfsfs)
	stateSourceObject.Load(1, &fs.privateMF)
	stateSourceObject.Load(2, &fs.mfp)
	stateSourceObject.Load(3, &fs.clock)
	stateSourceObject.Load(4, &fs.devMinor)
	stateSourceObject.Load(5, &fs.mopts)
	stateSourceObject.Load(6, &fs.usage)
	stateSourceObject.Load(7, &fs.nextInoMinusOne)
	stateSourceObject.Load(8, &fs.root)
	stateSourceObject.Load(9, &fs.maxFilenameLen)
	stateSourceObject.Load(10, &fs.maxSizeInPages)
	stateSourceObject.Load(11, &fs.pagesUsed)
	stateSourceObject.AfterLoad(fs.afterLoad)
}

func (f *FilesystemOpts) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.FilesystemOpts"
}

func (f *FilesystemOpts) StateFields() []string {
	return []string{
		"RootFileType",
		"RootSymlinkTarget",
		"FilesystemType",
		"Usage",
		"MaxFilenameLen",
		"FilestoreFD",
	}
}

func (f *FilesystemOpts) beforeSave() {}

// +checklocksignore
func (f *FilesystemOpts) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.RootFileType)
	stateSinkObject.Save(1, &f.RootSymlinkTarget)
	stateSinkObject.Save(2, &f.FilesystemType)
	stateSinkObject.Save(3, &f.Usage)
	stateSinkObject.Save(4, &f.MaxFilenameLen)
	stateSinkObject.Save(5, &f.FilestoreFD)
}

func (f *FilesystemOpts) afterLoad() {}

// +checklocksignore
func (f *FilesystemOpts) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.RootFileType)
	stateSourceObject.Load(1, &f.RootSymlinkTarget)
	stateSourceObject.Load(2, &f.FilesystemType)
	stateSourceObject.Load(3, &f.Usage)
	stateSourceObject.Load(4, &f.MaxFilenameLen)
	stateSourceObject.Load(5, &f.FilestoreFD)
}

func (d *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.dentry"
}

func (d *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"parent",
		"name",
		"dentryEntry",
		"inode",
	}
}

func (d *dentry) beforeSave() {}

// +checklocksignore
func (d *dentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.vfsd)
	stateSinkObject.Save(1, &d.parent)
	stateSinkObject.Save(2, &d.name)
	stateSinkObject.Save(3, &d.dentryEntry)
	stateSinkObject.Save(4, &d.inode)
}

func (d *dentry) afterLoad() {}

// +checklocksignore
func (d *dentry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsd)
	stateSourceObject.Load(1, &d.parent)
	stateSourceObject.Load(2, &d.name)
	stateSourceObject.Load(3, &d.dentryEntry)
	stateSourceObject.Load(4, &d.inode)
}

func (i *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.inode"
}

func (i *inode) StateFields() []string {
	return []string{
		"fs",
		"refs",
		"xattrs",
		"mode",
		"nlink",
		"uid",
		"gid",
		"ino",
		"atime",
		"ctime",
		"mtime",
		"locks",
		"watches",
		"impl",
	}
}

func (i *inode) beforeSave() {}

// +checklocksignore
func (i *inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.fs)
	stateSinkObject.Save(1, &i.refs)
	stateSinkObject.Save(2, &i.xattrs)
	stateSinkObject.Save(3, &i.mode)
	stateSinkObject.Save(4, &i.nlink)
	stateSinkObject.Save(5, &i.uid)
	stateSinkObject.Save(6, &i.gid)
	stateSinkObject.Save(7, &i.ino)
	stateSinkObject.Save(8, &i.atime)
	stateSinkObject.Save(9, &i.ctime)
	stateSinkObject.Save(10, &i.mtime)
	stateSinkObject.Save(11, &i.locks)
	stateSinkObject.Save(12, &i.watches)
	stateSinkObject.Save(13, &i.impl)
}

func (i *inode) afterLoad() {}

// +checklocksignore
func (i *inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.fs)
	stateSourceObject.Load(1, &i.refs)
	stateSourceObject.Load(2, &i.xattrs)
	stateSourceObject.Load(3, &i.mode)
	stateSourceObject.Load(4, &i.nlink)
	stateSourceObject.Load(5, &i.uid)
	stateSourceObject.Load(6, &i.gid)
	stateSourceObject.Load(7, &i.ino)
	stateSourceObject.Load(8, &i.atime)
	stateSourceObject.Load(9, &i.ctime)
	stateSourceObject.Load(10, &i.mtime)
	stateSourceObject.Load(11, &i.locks)
	stateSourceObject.Load(12, &i.watches)
	stateSourceObject.Load(13, &i.impl)
}

func (fd *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.fileDescription"
}

func (fd *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
	}
}

func (fd *fileDescription) beforeSave() {}

// +checklocksignore
func (fd *fileDescription) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.LockFD)
}

func (fd *fileDescription) afterLoad() {}

// +checklocksignore
func (fd *fileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.LockFD)
}

func init() {
	state.Register((*dentryList)(nil))
	state.Register((*dentryEntry)(nil))
	state.Register((*deviceFile)(nil))
	state.Register((*directory)(nil))
	state.Register((*directoryFD)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*namedPipe)(nil))
	state.Register((*regularFile)(nil))
	state.Register((*regularFileFD)(nil))
	state.Register((*socketFile)(nil))
	state.Register((*symlink)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*FilesystemOpts)(nil))
	state.Register((*dentry)(nil))
	state.Register((*inode)(nil))
	state.Register((*fileDescription)(nil))
}
