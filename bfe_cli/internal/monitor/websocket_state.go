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

var websocketState = NewWebsocketState()

var websocketStateCmd = &cobra.Command{
	Use:   websocketState.Category.Usage(),
	Short: "Show websocket state",
	Run:   websocketStateRun,
}

func init() {
	base.BindMetricFlags(websocketStateCmd.Flags())
}

func websocketStateRun(_ *cobra.Command, _ []string) {
	websocketState.Category.Request().EndBytes(websocketState.Category.EndBytes)
}

type WebsocketState struct {
	Category base.ICategory
}

func NewWebsocketState() *WebsocketState {
	s := &WebsocketState{}
	c := base.NewMetricCategory("websocket_state")
	s.Category = c

	return s
}
