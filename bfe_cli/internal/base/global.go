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

import "github.com/spf13/pflag"

var (
	Server      string
	MonitorPort uint16
	Debug       bool
)

func BindGlobalFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&Server, "server", "s", "localhost", "bfe server address")
	flags.Uint16VarP(&MonitorPort, "port", "p", 8421, "bfe server monitor port")
	flags.BoolVar(&Debug, "debug", false, "debug request bfe server")
}
