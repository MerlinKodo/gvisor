package fsutil

import (
	"reflect"

	"github.com/metacubex/gvisor/pkg/sync"
	"github.com/metacubex/gvisor/pkg/sync/locking"
)

// Mutex is sync.Mutex with the correctness validator.
type mapsMutex struct {
	mu sync.Mutex
}

var mapsprefixIndex *locking.MutexClass

// lockNames is a list of user-friendly lock names.
// Populated in init.
var mapslockNames []string

// lockNameIndex is used as an index passed to NestedLock and NestedUnlock,
// refering to an index within lockNames.
// Values are specified using the "consts" field of go_template_instance.
type mapslockNameIndex int

// DO NOT REMOVE: The following function automatically replaced with lock index constants.
// LOCK_NAME_INDEX_CONSTANTS
const ()

// Lock locks m.
// +checklocksignore
func (m *mapsMutex) Lock() {
	locking.AddGLock(mapsprefixIndex, -1)
	m.mu.Lock()
}

// NestedLock locks m knowing that another lock of the same type is held.
// +checklocksignore
func (m *mapsMutex) NestedLock(i mapslockNameIndex) {
	locking.AddGLock(mapsprefixIndex, int(i))
	m.mu.Lock()
}

// Unlock unlocks m.
// +checklocksignore
func (m *mapsMutex) Unlock() {
	locking.DelGLock(mapsprefixIndex, -1)
	m.mu.Unlock()
}

// NestedUnlock unlocks m knowing that another lock of the same type is held.
// +checklocksignore
func (m *mapsMutex) NestedUnlock(i mapslockNameIndex) {
	locking.DelGLock(mapsprefixIndex, int(i))
	m.mu.Unlock()
}

// DO NOT REMOVE: The following function is automatically replaced.
func mapsinitLockNames() {}

func init() {
	mapsinitLockNames()
	mapsprefixIndex = locking.NewMutexClass(reflect.TypeOf(mapsMutex{}), mapslockNames)
}
