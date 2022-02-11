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

var balState = NewBalState()

var balStateCmd = &cobra.Command{
	Use:   balState.Category.Usage(),
	Short: "Show bal state",
	Run:   balStateRun,
}

func init() {
	base.BindMetricFlags(balStateCmd.Flags())
}

func balStateRun(_ *cobra.Command, _ []string) {
	balState.Category.Request().EndBytes(balState.Category.EndBytes)
}

type BalState struct {
	Category base.ICategory
}

func NewBalState() *BalState {
	s := BalState{}
	c := base.NewMetricCategory("bal_state")
	s.Category = c
	return &s
}
