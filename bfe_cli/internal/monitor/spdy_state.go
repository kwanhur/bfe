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

var spdyState = NewSpdyState()

var spdyStateCmd = &cobra.Command{
	Use:   spdyState.Category.Usage(),
	Short: "Show spdy state",
	Run:   spdyStateRun,
}

func init() {
	base.BindMetricFlags(spdyStateCmd.Flags())
}

func spdyStateRun(_ *cobra.Command, _ []string) {
	spdyState.Category.Request().EndBytes(spdyState.Category.EndBytes)
}

type SpdyState struct {
	Category base.ICategory
}

func NewSpdyState() *SpdyState {
	s := &SpdyState{}
	c := base.NewMetricCategory("spdy_state")
	s.Category = c

	return s
}
