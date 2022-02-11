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

package module

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/bfenetworks/bfe/bfe_cli/internal/base"
	"github.com/bfenetworks/bfe/bfe_util/json"
)

var log = base.Logger

var moduleHandlers = NewModuleHandlers()

var moduleStatus = NewModuleStatus()

var moduleStatusCmd = &cobra.Command{
	Use:     moduleStatus.Category.Usage(),
	Short:   "Show module status",
	Run:     moduleStatusRun,
	Aliases: []string{"status"},
}

var moduleHandlersCmd = &cobra.Command{
	Use:     moduleHandlers.Category.Usage(),
	Short:   "Show module handlers",
	Run:     moduleHandlersRun,
	Aliases: []string{"ls", "handlers"},
}

func init() {
	base.BindPrettyFlag(moduleHandlersCmd.Flags())
	base.BindPrettyFlag(moduleStatusCmd.Flags())
}

func moduleHandlersRun(_ *cobra.Command, _ []string) {
	moduleHandlers.Category.Request().EndBytes(moduleHandlers.Category.EndBytes)
}

type ModuleHandlers struct {
	Category base.ICategory
}

func NewModuleHandlers() *ModuleHandlers {
	s := &ModuleHandlers{}
	c := base.NewMonitorCategory("module_handlers", s.Output)
	s.Category = c

	return s
}

func (s *ModuleHandlers) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case "json":
		data := make(map[string][]string)
		var err error
		if err = json.Unmarshal(body, &data); err != nil {
			log.Error(err)
			return
		}

		if base.Pretty {
			body, err = json.MarshalIndent(data, "", "    ")
			if err != nil {
				log.Error(err)
				return
			}
		}
		log.Println(string(body))
	}
}

func moduleStatusRun(_ *cobra.Command, _ []string) {
	moduleStatus.Category.Request().EndBytes(moduleStatus.Category.EndBytes)
}

type ModuleStatus struct {
	Category base.ICategory
}

func NewModuleStatus() *ModuleStatus {
	s := &ModuleStatus{}
	c := base.NewMonitorCategory("module_status", s.Output)
	s.Category = c

	return s
}

func (s *ModuleStatus) Output(body []byte) {
	switch base.Format {
	default:
		log.Print(string(body))
	case base.FormatJSON:
		data := struct {
			Available []string `json:"available"`
			Enabled   []string `json:"enabled"`
		}{}
		var err error
		if err = json.Unmarshal(body, &data); err != nil {
			log.Error(err)
			return
		}

		if base.Pretty {
			body, err = json.MarshalIndent(data, "", "    ")
			if err != nil {
				log.Error(err)
				return
			}
		}
		log.Println(string(body))
	}
}
