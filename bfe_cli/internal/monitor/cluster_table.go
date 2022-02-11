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
	"github.com/bfenetworks/bfe/bfe_route"
	"github.com/bfenetworks/bfe/bfe_util/json"
)

var clusterTableVersion = NewClusterTableVersion()

var clusterTableVersionCmd = &cobra.Command{
	Use:   clusterTableVersion.Category.Usage(),
	Short: "Show cluster table version",
	Run:   clusterTableVersionRun,
}

func init() {
	base.BindMonitorFlags(clusterTableVersionCmd.Flags())
}

func clusterTableVersionRun(_ *cobra.Command, _ []string) {
	clusterTableVersion.Category.Request().EndBytes(clusterTableVersion.Category.EndBytes)
}

type ClusterTableVersion struct {
	Category base.ICategory
}

func NewClusterTableVersion() *ClusterTableVersion {
	s := &ClusterTableVersion{}
	c := base.NewMonitorCategory("cluster_table_version", s.Output)
	s.Category = c

	return s
}

func (v *ClusterTableVersion) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case base.FormatJSON:
		version := bfe_route.ClusterVersion{}
		var err error
		if err = json.Unmarshal(body, &version); err != nil {
			log.Error(err)
			return
		}

		if base.Pretty {
			body, err = json.MarshalIndent(version, "", "    ")
			if err != nil {
				log.Error(err)
				return
			}
		}
		log.Println(string(body))
	}
}
