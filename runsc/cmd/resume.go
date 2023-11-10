// Copyright 2018 The gVisor Authors.
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

package cmd

import (
	"context"

	"github.com/MerlinKodo/gvisor/runsc/cmd/util"
	"github.com/MerlinKodo/gvisor/runsc/config"
	"github.com/MerlinKodo/gvisor/runsc/container"
	"github.com/MerlinKodo/gvisor/runsc/flag"
	"github.com/google/subcommands"
)

// Resume implements subcommands.Command for the "resume" command.
type Resume struct{}

// Name implements subcommands.Command.Name.
func (*Resume) Name() string {
	return "resume"
}

// Synopsis implements subcommands.Command.Synopsis.
func (*Resume) Synopsis() string {
	return "Resume unpauses a paused container"
}

// Usage implements subcommands.Command.Usage.
func (*Resume) Usage() string {
	return `resume <container id> - resume a paused container.
`
}

// SetFlags implements subcommands.Command.SetFlags.
func (r *Resume) SetFlags(*flag.FlagSet) {
}

// Execute implements subcommands.Command.Execute.
func (r *Resume) Execute(_ context.Context, f *flag.FlagSet, args ...any) subcommands.ExitStatus {
	if f.NArg() != 1 {
		f.Usage()
		return subcommands.ExitUsageError
	}

	id := f.Arg(0)
	conf := args[0].(*config.Config)

	cont, err := container.Load(conf.RootDir, container.FullID{ContainerID: id}, container.LoadOpts{})
	if err != nil {
		util.Fatalf("loading container: %v", err)
	}

	if err := cont.Resume(); err != nil {
		util.Fatalf("resume failed: %v", err)
	}

	return subcommands.ExitSuccess
}
