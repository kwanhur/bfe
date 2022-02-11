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
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

import (
	"github.com/bfenetworks/bfe/bfe_balance"
	"github.com/bfenetworks/bfe/bfe_cli/internal/base"
	"github.com/bfenetworks/bfe/bfe_util/json"
)

var cluster string
var balTableStatus = NewBalTableStatus()
var balTableVersion = NewBalTableVersion()

var balTableStatusCmd = &cobra.Command{
	Use:   balTableStatus.Category.Usage(),
	Short: "Show bal table status",
	Run:   balTableStatusRun,
}

var balTableVersionCmd = &cobra.Command{
	Use:   balTableVersion.Category.Usage(),
	Short: "Show bal table version",
	Run:   balTableVersionRun,
}

func init() {
	base.BindMonitorFlags(balTableStatusCmd.Flags())
	base.BindMonitorFlags(balTableVersionCmd.Flags())

	balTableStatusCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "specify cluster name")
}

func balTableStatusRun(_ *cobra.Command, _ []string) {
	balTableStatus.Category.Request().EndBytes(balTableStatus.Category.EndBytes)
}

type BalTableStatus struct {
	Category base.ICategory
}

func NewBalTableStatus() *BalTableStatus {
	s := &BalTableStatus{}
	c := base.NewMonitorCategory("bal_table_status", s.Output)
	s.Category = c

	return s
}

func (s *BalTableStatus) Request() *gorequest.SuperAgent {
	return s.Category.Request().Param("cluster_name", cluster)
}

func (s *BalTableStatus) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case base.FormatJSON:
		var state interface{}
		if cluster == "" {
			state = bfe_balance.BalTableState{}
		} else {
			state = struct {
				Status string `json:"status"`
			}{}
		}
		var err error
		if err = json.Unmarshal(body, &state); err != nil {
			log.Error(err)
			return
		}

		if base.Pretty {
			body, err = json.MarshalIndent(state, "", "    ")
			if err != nil {
				log.Error(err)
				return
			}
		}
		log.Println(string(body))
	}
}

func balTableVersionRun(_ *cobra.Command, _ []string) {
	balTableVersion.Category.Request().EndBytes(balTableVersion.Category.EndBytes)
}

type BalTableVersion struct {
	Category base.ICategory
}

func NewBalTableVersion() *BalTableVersion {
	s := &BalTableVersion{}
	c := base.NewMonitorCategory("bal_table_version", s.Output)
	s.Category = c

	return s
}

func (v *BalTableVersion) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case base.FormatJSON:
		version := bfe_balance.BalVersion{}
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
