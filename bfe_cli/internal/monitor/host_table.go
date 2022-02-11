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

var hostTableStatus = NewHostTableStatus()
var hostTableVersion = NewHostTableVersion()

var hostTableStatusCmd = &cobra.Command{
	Use:   hostTableStatus.Category.Usage(),
	Short: "Show host table status",
	Run:   hostTableStatusRun,
}

var hostTableVersionCmd = &cobra.Command{
	Use:   hostTableVersion.Category.Usage(),
	Short: "Show host table version",
	Run:   hostTableVersionRun,
}

func init() {
	base.BindMonitorFlags(hostTableStatusCmd.Flags())
	base.BindMonitorFlags(hostTableVersionCmd.Flags())
}

func hostTableStatusRun(_ *cobra.Command, _ []string) {
	hostTableStatus.Category.Request().EndBytes(hostTableStatus.Category.EndBytes)
}

type HostTableStatus struct {
	Category base.ICategory
}

func NewHostTableStatus() *HostTableStatus {
	s := &HostTableStatus{}
	c := base.NewMonitorCategory("host_table_status", s.Output)
	s.Category = c

	return s
}

func (s *HostTableStatus) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case base.FormatJSON:
		status := bfe_route.Status{}
		var err error
		if err = json.Unmarshal(body, &status); err != nil {
			log.Error(err)
			return
		}

		if base.Pretty {
			body, err = json.MarshalIndent(status, "", "    ")
			if err != nil {
				log.Error(err)
				return
			}
		}
		log.Println(string(body))
	}
}

func hostTableVersionRun(_ *cobra.Command, _ []string) {
	hostTableVersion.Category.Request().EndBytes(hostTableVersion.Category.EndBytes)
}

type HostTableVersion struct {
	Category base.ICategory
}

func NewHostTableVersion() *HostTableVersion {
	s := &HostTableVersion{}
	c := base.NewMonitorCategory("host_table_version", s.Output)
	s.Category = c

	return s
}

func (v *HostTableVersion) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case base.FormatJSON:
		versions := bfe_route.Versions{}
		var err error
		if err = json.Unmarshal(body, &versions); err != nil {
			log.Error(err)
			return
		}

		if base.Pretty {
			body, err = json.MarshalIndent(versions, "", "    ")
			if err != nil {
				log.Error(err)
				return
			}
		}
		log.Println(string(body))
	}
}
