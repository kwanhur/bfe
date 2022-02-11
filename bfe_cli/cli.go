// Copyright 2022 The BFE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/bfenetworks/bfe/bfe_cli/internal/base"
	"github.com/bfenetworks/bfe/bfe_cli/internal/config"
	"github.com/bfenetworks/bfe/bfe_cli/internal/module"
	"github.com/bfenetworks/bfe/bfe_cli/internal/monitor"
)

var log = base.Logger
var commit string

var bfecli = &cobra.Command{
	Use:   "bfecli",
	Short: "bfecli: commandline to interact with bfe server",
	Long: `bfecli: a friendly commandline to interact with bfe server,like:
- list supported modules
- fetch runtime metrics
- reload configuration`,
	Aliases: []string{"bfectl"},
	Run: func(cmd *cobra.Command, args []string) {
		if base.Version {
			showVersion()
			return
		}

		if base.Verbose {
			showVerbose()
		}
	},
}

func init() {
	bfecli.AddCommand(versionCmd, verboseCmd)
	bfecli.AddCommand(monitor.MonCmd)
	bfecli.AddCommand(module.ModCmd)
	bfecli.AddCommand(config.ReloadCmd)
}

func init() {
	global := bfecli.PersistentFlags()
	base.BindGlobalFlags(global)
	base.BindVerFlags(bfecli.Flags())
}

func main() {
	if err := bfecli.Execute(); err != nil {
		log.Fatal(err)
	}
}
