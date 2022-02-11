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

package base

import (
	"github.com/spf13/pflag"
)

const FormatJSON = "json"

var (
	Format = "json"
	Path   = ""

	Pretty bool
	Diff   bool

	Version bool
	Verbose bool
)

// BindPrettyFlag add Pretty flag into sub-commands
func BindPrettyFlag(flags *pflag.FlagSet) {
	flags.BoolVar(&Pretty, "pretty", false, "show json text pretty")
}

// BindMonitorFlags add common flags into sub-commands, include Pretty Format
func BindMonitorFlags(flags *pflag.FlagSet) {
	BindPrettyFlag(flags)
	flags.StringVar(&Format, "format", "json", "specify text format")
}

// BindMetricFlags add common flags into sub-commands, include Pretty Format Diff
func BindMetricFlags(flags *pflag.FlagSet) {
	BindMonitorFlags(flags)
	flags.BoolVar(&Diff, "diff", false, "show state metric diff data")
}

// BindConfigFlags add common flags into sub-commands, include Pretty Path
func BindConfigFlags(flags *pflag.FlagSet) {
	BindPrettyFlag(flags)
	flags.StringVar(&Path, "path", "", "specify root path of bfe conf")
}

// BindVerFlags add Version Verbose flags into sub-commands
func BindVerFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&Version, "version", "v", false, "Show version number of bfecli")
	flags.BoolVarP(&Verbose, "verbose", "V", false, "Show verbose information of bfecli")
}
