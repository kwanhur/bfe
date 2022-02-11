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

var proxyProtocolState = NewProxyProtocolState()

var proxyProtocolStateCmd = &cobra.Command{
	Use:   proxyProtocolState.Category.Usage(),
	Short: "Show proxy protocol state",
	Run:   proxyProtocolStateRun,
}

func init() {
	base.BindMetricFlags(proxyProtocolStateCmd.Flags())
}

func proxyProtocolStateRun(_ *cobra.Command, _ []string) {
	proxyProtocolState.Category.Request().EndBytes(proxyProtocolState.Category.EndBytes)
}

type ProxyProtocolState struct {
	Category base.ICategory
}

func NewProxyProtocolState() *ProxyProtocolState {
	s := &ProxyProtocolState{}
	c := base.NewMetricCategory("proxy_protocol_state")
	s.Category = c

	return s
}
