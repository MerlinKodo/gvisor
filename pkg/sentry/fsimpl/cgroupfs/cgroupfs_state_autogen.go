// automatically generated by stateify.

package cgroupfs

import (
	"github.com/metacubex/gvisor/pkg/state"
)

func (c *controllerCommon) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.controllerCommon"
}

func (c *controllerCommon) StateFields() []string {
	return []string{
		"ty",
		"fs",
		"parent",
	}
}

func (c *controllerCommon) beforeSave() {}

// +checklocksignore
func (c *controllerCommon) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.ty)
	stateSinkObject.Save(1, &c.fs)
	stateSinkObject.Save(2, &c.parent)
}

func (c *controllerCommon) afterLoad() {}

// +checklocksignore
func (c *controllerCommon) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.ty)
	stateSourceObject.Load(1, &c.fs)
	stateSourceObject.Load(2, &c.parent)
}

func (c *cgroupInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cgroupInode"
}

func (c *cgroupInode) StateFields() []string {
	return []string{
		"dir",
		"id",
		"controllers",
		"ts",
	}
}

func (c *cgroupInode) beforeSave() {}

// +checklocksignore
func (c *cgroupInode) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.dir)
	stateSinkObject.Save(1, &c.id)
	stateSinkObject.Save(2, &c.controllers)
	stateSinkObject.Save(3, &c.ts)
}

func (c *cgroupInode) afterLoad() {}

// +checklocksignore
func (c *cgroupInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.dir)
	stateSourceObject.Load(1, &c.id)
	stateSourceObject.Load(2, &c.controllers)
	stateSourceObject.Load(3, &c.ts)
}

func (d *cgroupProcsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cgroupProcsData"
}

func (d *cgroupProcsData) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (d *cgroupProcsData) beforeSave() {}

// +checklocksignore
func (d *cgroupProcsData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cgroupInode)
}

func (d *cgroupProcsData) afterLoad() {}

// +checklocksignore
func (d *cgroupProcsData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cgroupInode)
}

func (d *tasksData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.tasksData"
}

func (d *tasksData) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (d *tasksData) beforeSave() {}

// +checklocksignore
func (d *tasksData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cgroupInode)
}

func (d *tasksData) afterLoad() {}

// +checklocksignore
func (d *tasksData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cgroupInode)
}

func (fsType *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.FilesystemType"
}

func (fsType *FilesystemType) StateFields() []string {
	return []string{}
}

func (fsType *FilesystemType) beforeSave() {}

// +checklocksignore
func (fsType *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fsType.beforeSave()
}

func (fsType *FilesystemType) afterLoad() {}

// +checklocksignore
func (fsType *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (i *InitialCgroup) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.InitialCgroup"
}

func (i *InitialCgroup) StateFields() []string {
	return []string{
		"Path",
		"SetOwner",
		"UID",
		"GID",
		"SetMode",
		"Mode",
	}
}

func (i *InitialCgroup) beforeSave() {}

// +checklocksignore
func (i *InitialCgroup) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Path)
	stateSinkObject.Save(1, &i.SetOwner)
	stateSinkObject.Save(2, &i.UID)
	stateSinkObject.Save(3, &i.GID)
	stateSinkObject.Save(4, &i.SetMode)
	stateSinkObject.Save(5, &i.Mode)
}

func (i *InitialCgroup) afterLoad() {}

// +checklocksignore
func (i *InitialCgroup) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Path)
	stateSourceObject.Load(1, &i.SetOwner)
	stateSourceObject.Load(2, &i.UID)
	stateSourceObject.Load(3, &i.GID)
	stateSourceObject.Load(4, &i.SetMode)
	stateSourceObject.Load(5, &i.Mode)
}

func (i *InternalData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.InternalData"
}

func (i *InternalData) StateFields() []string {
	return []string{
		"DefaultControlValues",
		"InitialCgroup",
	}
}

func (i *InternalData) beforeSave() {}

// +checklocksignore
func (i *InternalData) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.DefaultControlValues)
	stateSinkObject.Save(1, &i.InitialCgroup)
}

func (i *InternalData) afterLoad() {}

// +checklocksignore
func (i *InternalData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.DefaultControlValues)
	stateSourceObject.Load(1, &i.InitialCgroup)
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
		"hierarchyID",
		"hierarchyName",
		"controllers",
		"kcontrollers",
		"numCgroups",
		"root",
		"effectiveRoot",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Filesystem)
	stateSinkObject.Save(1, &fs.devMinor)
	stateSinkObject.Save(2, &fs.hierarchyID)
	stateSinkObject.Save(3, &fs.hierarchyName)
	stateSinkObject.Save(4, &fs.controllers)
	stateSinkObject.Save(5, &fs.kcontrollers)
	stateSinkObject.Save(6, &fs.numCgroups)
	stateSinkObject.Save(7, &fs.root)
	stateSinkObject.Save(8, &fs.effectiveRoot)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Filesystem)
	stateSourceObject.Load(1, &fs.devMinor)
	stateSourceObject.Load(2, &fs.hierarchyID)
	stateSourceObject.Load(3, &fs.hierarchyName)
	stateSourceObject.Load(4, &fs.controllers)
	stateSourceObject.Load(5, &fs.kcontrollers)
	stateSourceObject.Load(6, &fs.numCgroups)
	stateSourceObject.Load(7, &fs.root)
	stateSourceObject.Load(8, &fs.effectiveRoot)
}

func (i *implStatFS) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.implStatFS"
}

func (i *implStatFS) StateFields() []string {
	return []string{}
}

func (i *implStatFS) beforeSave() {}

// +checklocksignore
func (i *implStatFS) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *implStatFS) afterLoad() {}

// +checklocksignore
func (i *implStatFS) StateLoad(stateSourceObject state.Source) {
}

func (d *dir) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.dir"
}

func (d *dir) StateFields() []string {
	return []string{
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeDirectoryNoNewChildren",
		"InodeNoopRefCount",
		"InodeNotAnonymous",
		"InodeNotSymlink",
		"InodeWatches",
		"OrderedChildren",
		"implStatFS",
		"locks",
		"fs",
		"cgi",
	}
}

func (d *dir) beforeSave() {}

// +checklocksignore
func (d *dir) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.InodeAlwaysValid)
	stateSinkObject.Save(1, &d.InodeAttrs)
	stateSinkObject.Save(2, &d.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(3, &d.InodeNoopRefCount)
	stateSinkObject.Save(4, &d.InodeNotAnonymous)
	stateSinkObject.Save(5, &d.InodeNotSymlink)
	stateSinkObject.Save(6, &d.InodeWatches)
	stateSinkObject.Save(7, &d.OrderedChildren)
	stateSinkObject.Save(8, &d.implStatFS)
	stateSinkObject.Save(9, &d.locks)
	stateSinkObject.Save(10, &d.fs)
	stateSinkObject.Save(11, &d.cgi)
}

func (d *dir) afterLoad() {}

// +checklocksignore
func (d *dir) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.InodeAlwaysValid)
	stateSourceObject.Load(1, &d.InodeAttrs)
	stateSourceObject.Load(2, &d.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(3, &d.InodeNoopRefCount)
	stateSourceObject.Load(4, &d.InodeNotAnonymous)
	stateSourceObject.Load(5, &d.InodeNotSymlink)
	stateSourceObject.Load(6, &d.InodeWatches)
	stateSourceObject.Load(7, &d.OrderedChildren)
	stateSourceObject.Load(8, &d.implStatFS)
	stateSourceObject.Load(9, &d.locks)
	stateSourceObject.Load(10, &d.fs)
	stateSourceObject.Load(11, &d.cgi)
}

func (f *controllerFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.controllerFile"
}

func (f *controllerFile) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"implStatFS",
		"allowBackgroundAccess",
	}
}

func (f *controllerFile) beforeSave() {}

// +checklocksignore
func (f *controllerFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.DynamicBytesFile)
	stateSinkObject.Save(1, &f.implStatFS)
	stateSinkObject.Save(2, &f.allowBackgroundAccess)
}

func (f *controllerFile) afterLoad() {}

// +checklocksignore
func (f *controllerFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.DynamicBytesFile)
	stateSourceObject.Load(1, &f.implStatFS)
	stateSourceObject.Load(2, &f.allowBackgroundAccess)
}

func (f *staticControllerFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.staticControllerFile"
}

func (f *staticControllerFile) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"StaticData",
	}
}

func (f *staticControllerFile) beforeSave() {}

// +checklocksignore
func (f *staticControllerFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.DynamicBytesFile)
	stateSinkObject.Save(1, &f.StaticData)
}

func (f *staticControllerFile) afterLoad() {}

// +checklocksignore
func (f *staticControllerFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.DynamicBytesFile)
	stateSourceObject.Load(1, &f.StaticData)
}

func (f *stubControllerFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.stubControllerFile"
}

func (f *stubControllerFile) StateFields() []string {
	return []string{
		"controllerFile",
		"data",
	}
}

func (f *stubControllerFile) beforeSave() {}

// +checklocksignore
func (f *stubControllerFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.controllerFile)
	stateSinkObject.Save(1, &f.data)
}

func (f *stubControllerFile) afterLoad() {}

// +checklocksignore
func (f *stubControllerFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.controllerFile)
	stateSourceObject.Load(1, &f.data)
}

func (c *cpuController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuController"
}

func (c *cpuController) StateFields() []string {
	return []string{
		"controllerCommon",
		"controllerStateless",
		"controllerNoResource",
		"cfsPeriod",
		"cfsQuota",
		"shares",
	}
}

func (c *cpuController) beforeSave() {}

// +checklocksignore
func (c *cpuController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.controllerStateless)
	stateSinkObject.Save(2, &c.controllerNoResource)
	stateSinkObject.Save(3, &c.cfsPeriod)
	stateSinkObject.Save(4, &c.cfsQuota)
	stateSinkObject.Save(5, &c.shares)
}

func (c *cpuController) afterLoad() {}

// +checklocksignore
func (c *cpuController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.controllerStateless)
	stateSourceObject.Load(2, &c.controllerNoResource)
	stateSourceObject.Load(3, &c.cfsPeriod)
	stateSourceObject.Load(4, &c.cfsQuota)
	stateSourceObject.Load(5, &c.shares)
}

func (c *cpuacctController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctController"
}

func (c *cpuacctController) StateFields() []string {
	return []string{
		"controllerCommon",
		"controllerNoResource",
		"taskCommittedCharges",
		"usage",
	}
}

func (c *cpuacctController) beforeSave() {}

// +checklocksignore
func (c *cpuacctController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.controllerNoResource)
	stateSinkObject.Save(2, &c.taskCommittedCharges)
	stateSinkObject.Save(3, &c.usage)
}

func (c *cpuacctController) afterLoad() {}

// +checklocksignore
func (c *cpuacctController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.controllerNoResource)
	stateSourceObject.Load(2, &c.taskCommittedCharges)
	stateSourceObject.Load(3, &c.usage)
}

func (c *cpuacctCgroup) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctCgroup"
}

func (c *cpuacctCgroup) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (c *cpuacctCgroup) beforeSave() {}

// +checklocksignore
func (c *cpuacctCgroup) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.cgroupInode)
}

func (c *cpuacctCgroup) afterLoad() {}

// +checklocksignore
func (c *cpuacctCgroup) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.cgroupInode)
}

func (d *cpuacctStatData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctStatData"
}

func (d *cpuacctStatData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctStatData) beforeSave() {}

// +checklocksignore
func (d *cpuacctStatData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctStatData) afterLoad() {}

// +checklocksignore
func (d *cpuacctStatData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctUsageData"
}

func (d *cpuacctUsageData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctUsageData) beforeSave() {}

// +checklocksignore
func (d *cpuacctUsageData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageData) afterLoad() {}

// +checklocksignore
func (d *cpuacctUsageData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageUserData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctUsageUserData"
}

func (d *cpuacctUsageUserData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctUsageUserData) beforeSave() {}

// +checklocksignore
func (d *cpuacctUsageUserData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageUserData) afterLoad() {}

// +checklocksignore
func (d *cpuacctUsageUserData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageSysData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctUsageSysData"
}

func (d *cpuacctUsageSysData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctUsageSysData) beforeSave() {}

// +checklocksignore
func (d *cpuacctUsageSysData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageSysData) afterLoad() {}

// +checklocksignore
func (d *cpuacctUsageSysData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (c *cpusetController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpusetController"
}

func (c *cpusetController) StateFields() []string {
	return []string{
		"controllerCommon",
		"controllerStateless",
		"controllerNoResource",
		"maxCpus",
		"maxMems",
		"cpus",
		"mems",
	}
}

func (c *cpusetController) beforeSave() {}

// +checklocksignore
func (c *cpusetController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.controllerStateless)
	stateSinkObject.Save(2, &c.controllerNoResource)
	stateSinkObject.Save(3, &c.maxCpus)
	stateSinkObject.Save(4, &c.maxMems)
	stateSinkObject.Save(5, &c.cpus)
	stateSinkObject.Save(6, &c.mems)
}

func (c *cpusetController) afterLoad() {}

// +checklocksignore
func (c *cpusetController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.controllerStateless)
	stateSourceObject.Load(2, &c.controllerNoResource)
	stateSourceObject.Load(3, &c.maxCpus)
	stateSourceObject.Load(4, &c.maxMems)
	stateSourceObject.Load(5, &c.cpus)
	stateSourceObject.Load(6, &c.mems)
}

func (d *cpusData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpusData"
}

func (d *cpusData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *cpusData) beforeSave() {}

// +checklocksignore
func (d *cpusData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *cpusData) afterLoad() {}

// +checklocksignore
func (d *cpusData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (d *memsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.memsData"
}

func (d *memsData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *memsData) beforeSave() {}

// +checklocksignore
func (d *memsData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *memsData) afterLoad() {}

// +checklocksignore
func (d *memsData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (d *deviceID) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.deviceID"
}

func (d *deviceID) StateFields() []string {
	return []string{
		"controllerType",
		"major",
		"minor",
	}
}

func (d *deviceID) beforeSave() {}

// +checklocksignore
func (d *deviceID) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.controllerType)
	stateSinkObject.Save(1, &d.major)
	stateSinkObject.Save(2, &d.minor)
}

func (d *deviceID) afterLoad() {}

// +checklocksignore
func (d *deviceID) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.controllerType)
	stateSourceObject.Load(1, &d.major)
	stateSourceObject.Load(2, &d.minor)
}

func (c *devicesController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.devicesController"
}

func (c *devicesController) StateFields() []string {
	return []string{
		"controllerCommon",
		"controllerStateless",
		"controllerNoResource",
		"defaultAllow",
		"deviceRules",
	}
}

func (c *devicesController) beforeSave() {}

// +checklocksignore
func (c *devicesController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.controllerStateless)
	stateSinkObject.Save(2, &c.controllerNoResource)
	stateSinkObject.Save(3, &c.defaultAllow)
	stateSinkObject.Save(4, &c.deviceRules)
}

func (c *devicesController) afterLoad() {}

// +checklocksignore
func (c *devicesController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.controllerStateless)
	stateSourceObject.Load(2, &c.controllerNoResource)
	stateSourceObject.Load(3, &c.defaultAllow)
	stateSourceObject.Load(4, &c.deviceRules)
}

func (d *allowedDevicesData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.allowedDevicesData"
}

func (d *allowedDevicesData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *allowedDevicesData) beforeSave() {}

// +checklocksignore
func (d *allowedDevicesData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *allowedDevicesData) afterLoad() {}

// +checklocksignore
func (d *allowedDevicesData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (d *deniedDevicesData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.deniedDevicesData"
}

func (d *deniedDevicesData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *deniedDevicesData) beforeSave() {}

// +checklocksignore
func (d *deniedDevicesData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *deniedDevicesData) afterLoad() {}

// +checklocksignore
func (d *deniedDevicesData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (d *controlledDevicesData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.controlledDevicesData"
}

func (d *controlledDevicesData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *controlledDevicesData) beforeSave() {}

// +checklocksignore
func (d *controlledDevicesData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *controlledDevicesData) afterLoad() {}

// +checklocksignore
func (d *controlledDevicesData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (r *dirRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.dirRefs"
}

func (r *dirRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *dirRefs) beforeSave() {}

// +checklocksignore
func (r *dirRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *dirRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (c *jobController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.jobController"
}

func (c *jobController) StateFields() []string {
	return []string{
		"controllerCommon",
		"controllerStateless",
		"controllerNoResource",
		"id",
	}
}

func (c *jobController) beforeSave() {}

// +checklocksignore
func (c *jobController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.controllerStateless)
	stateSinkObject.Save(2, &c.controllerNoResource)
	stateSinkObject.Save(3, &c.id)
}

func (c *jobController) afterLoad() {}

// +checklocksignore
func (c *jobController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.controllerStateless)
	stateSourceObject.Load(2, &c.controllerNoResource)
	stateSourceObject.Load(3, &c.id)
}

func (c *memoryController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.memoryController"
}

func (c *memoryController) StateFields() []string {
	return []string{
		"controllerCommon",
		"controllerNoResource",
		"limitBytes",
		"softLimitBytes",
		"moveChargeAtImmigrate",
		"pressureLevel",
		"memCg",
	}
}

func (c *memoryController) beforeSave() {}

// +checklocksignore
func (c *memoryController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.controllerNoResource)
	stateSinkObject.Save(2, &c.limitBytes)
	stateSinkObject.Save(3, &c.softLimitBytes)
	stateSinkObject.Save(4, &c.moveChargeAtImmigrate)
	stateSinkObject.Save(5, &c.pressureLevel)
	stateSinkObject.Save(6, &c.memCg)
}

func (c *memoryController) afterLoad() {}

// +checklocksignore
func (c *memoryController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.controllerNoResource)
	stateSourceObject.Load(2, &c.limitBytes)
	stateSourceObject.Load(3, &c.softLimitBytes)
	stateSourceObject.Load(4, &c.moveChargeAtImmigrate)
	stateSourceObject.Load(5, &c.pressureLevel)
	stateSourceObject.Load(6, &c.memCg)
}

func (memCg *memoryCgroup) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.memoryCgroup"
}

func (memCg *memoryCgroup) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (memCg *memoryCgroup) beforeSave() {}

// +checklocksignore
func (memCg *memoryCgroup) StateSave(stateSinkObject state.Sink) {
	memCg.beforeSave()
	stateSinkObject.Save(0, &memCg.cgroupInode)
}

func (memCg *memoryCgroup) afterLoad() {}

// +checklocksignore
func (memCg *memoryCgroup) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &memCg.cgroupInode)
}

func (d *memoryUsageInBytesData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.memoryUsageInBytesData"
}

func (d *memoryUsageInBytesData) StateFields() []string {
	return []string{
		"memCg",
	}
}

func (d *memoryUsageInBytesData) beforeSave() {}

// +checklocksignore
func (d *memoryUsageInBytesData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.memCg)
}

func (d *memoryUsageInBytesData) afterLoad() {}

// +checklocksignore
func (d *memoryUsageInBytesData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.memCg)
}

func (c *pidsController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.pidsController"
}

func (c *pidsController) StateFields() []string {
	return []string{
		"controllerCommon",
		"isRoot",
		"pendingTotal",
		"pendingPool",
		"committed",
		"max",
	}
}

func (c *pidsController) beforeSave() {}

// +checklocksignore
func (c *pidsController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.isRoot)
	stateSinkObject.Save(2, &c.pendingTotal)
	stateSinkObject.Save(3, &c.pendingPool)
	stateSinkObject.Save(4, &c.committed)
	stateSinkObject.Save(5, &c.max)
}

func (c *pidsController) afterLoad() {}

// +checklocksignore
func (c *pidsController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.isRoot)
	stateSourceObject.Load(2, &c.pendingTotal)
	stateSourceObject.Load(3, &c.pendingPool)
	stateSourceObject.Load(4, &c.committed)
	stateSourceObject.Load(5, &c.max)
}

func (d *pidsCurrentData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.pidsCurrentData"
}

func (d *pidsCurrentData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *pidsCurrentData) beforeSave() {}

// +checklocksignore
func (d *pidsCurrentData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *pidsCurrentData) afterLoad() {}

// +checklocksignore
func (d *pidsCurrentData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (d *pidsMaxData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.pidsMaxData"
}

func (d *pidsMaxData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *pidsMaxData) beforeSave() {}

// +checklocksignore
func (d *pidsMaxData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *pidsMaxData) afterLoad() {}

// +checklocksignore
func (d *pidsMaxData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func init() {
	state.Register((*controllerCommon)(nil))
	state.Register((*cgroupInode)(nil))
	state.Register((*cgroupProcsData)(nil))
	state.Register((*tasksData)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*InitialCgroup)(nil))
	state.Register((*InternalData)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*implStatFS)(nil))
	state.Register((*dir)(nil))
	state.Register((*controllerFile)(nil))
	state.Register((*staticControllerFile)(nil))
	state.Register((*stubControllerFile)(nil))
	state.Register((*cpuController)(nil))
	state.Register((*cpuacctController)(nil))
	state.Register((*cpuacctCgroup)(nil))
	state.Register((*cpuacctStatData)(nil))
	state.Register((*cpuacctUsageData)(nil))
	state.Register((*cpuacctUsageUserData)(nil))
	state.Register((*cpuacctUsageSysData)(nil))
	state.Register((*cpusetController)(nil))
	state.Register((*cpusData)(nil))
	state.Register((*memsData)(nil))
	state.Register((*deviceID)(nil))
	state.Register((*devicesController)(nil))
	state.Register((*allowedDevicesData)(nil))
	state.Register((*deniedDevicesData)(nil))
	state.Register((*controlledDevicesData)(nil))
	state.Register((*dirRefs)(nil))
	state.Register((*jobController)(nil))
	state.Register((*memoryController)(nil))
	state.Register((*memoryCgroup)(nil))
	state.Register((*memoryUsageInBytesData)(nil))
	state.Register((*pidsController)(nil))
	state.Register((*pidsCurrentData)(nil))
	state.Register((*pidsMaxData)(nil))
}
