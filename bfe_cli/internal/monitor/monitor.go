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

package monitor

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/bfenetworks/bfe/bfe_cli/internal/base"
)

var log = base.Logger

// MonCmd monitor command
var MonCmd = &cobra.Command{
	Use:     "monitor",
	Short:   "Show bfe server monitor runtime",
	Aliases: []string{"mon"},
}

// init add sub-commands
func init() {
	MonCmd.AddCommand(hostTableStatusCmd, hostTableVersionCmd)
	MonCmd.AddCommand(clusterTableVersionCmd)
	MonCmd.AddCommand(balTableStatusCmd, balTableVersionCmd, balStateCmd)
	MonCmd.AddCommand(proxyStateCmd, proxyProtocolStateCmd, tlsStateCmd)
	MonCmd.AddCommand(spdyStateCmd, http2StateCmd, httpStateCmd, streamStateCmd, websocketStateCmd)
	MonCmd.AddCommand(proxyDelayCmd, proxyPostDelayCmd, proxyHandshakeDelayCmd, proxyHandshakeFullDelayCmd, proxyHandshakeResumeDelayCmd)
}
