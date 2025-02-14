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

//go:build linux && debug
// +build linux,debug

package platforms

import (
	// Import platforms that runsc might use.

	// The KVM platform is not included because it's incompatible with debug
	// builds. Unoptimized functions grow the stack too much and fail the nosplit
	// check.
	_ "github.com/MerlinKodo/gvisor/pkg/sentry/platform/ptrace"
	_ "github.com/MerlinKodo/gvisor/pkg/sentry/platform/systrap"
)
