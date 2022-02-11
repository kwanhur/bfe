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
	"strings"
)

import (
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

import (
	"github.com/bfenetworks/bfe/bfe_cli/internal/base"
)

var protos []string

var tlsConf = NewTlsConf()
var tlsSessionTicketKey = NewTlsSessionTicketKey()

var tlsConfCmd = &cobra.Command{
	Use:   tlsConf.Category.Usage(),
	Short: "Trigger tls conf to reload",
	Run:   tlsConfRun,
}

var tlsSessionTicketKeyCmd = &cobra.Command{
	Use:   tlsSessionTicketKey.Category.Usage(),
	Short: "Trigger tls session ticket key to reload",
	Run:   tlsSessionTicketKeyRun,
}

func init() {
	base.BindConfigFlags(tlsConfCmd.Flags())
	tlsConfCmd.Flags().StringSliceVar(&protos, "protos", protos, "specify tls protocol[spdy|h2] to enable or disable[+|-]")

	base.BindPrettyFlag(tlsSessionTicketKeyCmd.Flags())
}

func tlsConfRun(_ *cobra.Command, _ []string) {
	tlsConf.Request().EndBytes(tlsConf.Category.EndBytes)
}

type TlsConf struct {
	Category base.ICategory
}

func NewTlsConf() *TlsConf {
	s := &TlsConf{}
	c := base.NewConfigCategory("tls_conf")
	s.Category = c

	return s
}

func (s *TlsConf) Request() *gorequest.SuperAgent {
	return s.Category.Request().Param("enable", strings.Join(protos, ","))
}

func tlsSessionTicketKeyRun(_ *cobra.Command, _ []string) {
	tlsConf.Request().EndBytes(tlsConf.Category.EndBytes)
}

type TlsSessionTicketKey struct {
	Category base.ICategory
}

func NewTlsSessionTicketKey() *TlsSessionTicketKey {
	s := &TlsSessionTicketKey{}
	c := base.NewConfigCategory("tls_session_ticket_key")
	s.Category = c

	return s
}
