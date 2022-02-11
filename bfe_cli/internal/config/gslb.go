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

package config

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/bfenetworks/bfe/bfe_cli/internal/base"
)

var gslbDataConf = NewGslbDataConf()

var gslbDataConfCmd = &cobra.Command{
	Use:   gslbDataConf.Category.Usage(),
	Short: "Trigger gslb data conf to reload",
	Run:   gslbDataConfRun,
}

func init() {
	base.BindConfigFlags(gslbDataConfCmd.Flags())
}

func gslbDataConfRun(_ *cobra.Command, _ []string) {
	gslbDataConf.Category.Request().EndBytes(gslbDataConf.Category.EndBytes)
}

type GslbDataConf struct {
	Category base.ICategory
}

func NewGslbDataConf() *GslbDataConf {
	s := &GslbDataConf{}
	c := base.NewConfigCategory("gslb_data_conf")
	s.Category = c

	return s
}
