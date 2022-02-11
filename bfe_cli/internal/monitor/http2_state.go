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

var http2State = NewHttp2State()

var http2StateCmd = &cobra.Command{
	Use:   http2State.Category.Usage(),
	Short: "Show http2 state",
	Run:   http2StateRun,
}

func init() {
	base.BindMetricFlags(http2StateCmd.Flags())
}

func http2StateRun(_ *cobra.Command, _ []string) {
	http2State.Category.Request().EndBytes(http2State.Category.EndBytes)
}

type Http2State struct {
	Category base.ICategory
}

func NewHttp2State() *Http2State {
	s := &Http2State{}
	c := base.NewMetricCategory("http2_state")
	s.Category = c

	return s
}
