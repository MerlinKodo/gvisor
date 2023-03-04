// Copyright 2023 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sysmsg

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/metacubex/gvisor/pkg/cpuid"
)

// SighandlerBlob contains the compiled code of the sysmsg signal handler.
//
//go:embed sighandler.built-in.amd64.bin
var SighandlerBlob []byte

// ArchState defines variables specific to the architecture being
// used.
type ArchState struct {
	xsaveMode uint32
	fpLen     uint32
	fsgsbase  uint32
}

const (
	fxsave = iota
	xsave
	xsaveopt
	xsavec
)

// Init initializes the arch specific state.
func (s *ArchState) Init() {
	fs := cpuid.HostFeatureSet()

	fpLenUint, _ := fs.ExtendedStateSize()
	s.fpLen = uint32(fpLenUint)
	if fs.UseXsavec() {
		s.xsaveMode = xsavec
	} else if fs.UseXsaveopt() {
		s.xsaveMode = xsaveopt
	} else if fs.UseXsave() {
		s.xsaveMode = xsave
	} else {
		s.xsaveMode = fxsave
	}

	if fs.UseFSGSBASE() {
		s.fsgsbase = 1
	}
}

func (s *ArchState) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "sysmsg.ArchState{")
	fmt.Fprintf(&b, " xsaveMode %d", s.xsaveMode)
	fmt.Fprintf(&b, " fpLen %d", s.fpLen)
	b.WriteString(" }")

	return b.String()
}
