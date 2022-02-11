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

var proxyDelay = NewProxyDelay()
var proxyPostDelay = NewProxyPostDelay()
var proxyHandshakeDelay = NewProxyHandshakeDelay()
var proxyHandshakeFullDelay = NewProxyHandshakeFullDelay()
var proxyHandshakeResumeDelay = NewProxyHandshakeResumeDelay()

var proxyDelayCmd = &cobra.Command{
	Use:   proxyDelay.Category.Usage(),
	Short: "Show proxy delay",
	Run:   proxyDelayRun,
}

var proxyPostDelayCmd = &cobra.Command{
	Use:   proxyPostDelay.Category.Usage(),
	Short: "Show proxy post delay",
	Run:   proxyPostDelayRun,
}

var proxyHandshakeDelayCmd = &cobra.Command{
	Use:   proxyHandshakeDelay.Category.Usage(),
	Short: "Show proxy handshake delay",
	Run:   proxyHandshakeDelayRun,
}

var proxyHandshakeFullDelayCmd = &cobra.Command{
	Use:   proxyHandshakeFullDelay.Category.Usage(),
	Short: "Show proxy handshake full delay",
	Run:   proxyHandshakeFullDelayRun,
}

var proxyHandshakeResumeDelayCmd = &cobra.Command{
	Use:   proxyHandshakeResumeDelay.Category.Usage(),
	Short: "Show proxy handshake resume delay",
	Run:   proxyHandshakeResumeDelayRun,
}

func init() {
	base.BindMonitorFlags(proxyDelayCmd.Flags())
	base.BindMonitorFlags(proxyPostDelayCmd.Flags())
	base.BindMonitorFlags(proxyHandshakeDelayCmd.Flags())
	base.BindMonitorFlags(proxyHandshakeFullDelayCmd.Flags())
	base.BindMonitorFlags(proxyHandshakeResumeDelayCmd.Flags())
}

func proxyDelayRun(_ *cobra.Command, _ []string) {
	proxyDelay.Category.Request().EndBytes(proxyDelay.Category.EndBytes)
}

type ProxyDelay struct {
	Category base.ICategory
}

func NewProxyDelay() *ProxyDelay {
	s := &ProxyDelay{}
	c := base.NewMetricCategory("proxy_delay")
	s.Category = c

	return s
}

func proxyPostDelayRun(_ *cobra.Command, _ []string) {
	proxyPostDelay.Category.Request().EndBytes(proxyPostDelay.Category.EndBytes)
}

type ProxyPostDelay struct {
	Category base.ICategory
}

func NewProxyPostDelay() *ProxyPostDelay {
	s := &ProxyPostDelay{}
	c := base.NewMetricCategory("proxy_post_delay")
	s.Category = c

	return s
}

func proxyHandshakeDelayRun(_ *cobra.Command, _ []string) {
	proxyHandshakeDelay.Category.Request().EndBytes(proxyHandshakeDelay.Category.EndBytes)
}

type ProxyHandshakeDelay struct {
	Category base.ICategory
}

func NewProxyHandshakeDelay() *ProxyHandshakeDelay {
	s := &ProxyHandshakeDelay{}
	c := base.NewMetricCategory("proxy_handshake_delay")
	s.Category = c

	return s
}

func proxyHandshakeFullDelayRun(_ *cobra.Command, _ []string) {
	proxyHandshakeFullDelay.Category.Request().EndBytes(proxyHandshakeFullDelay.Category.EndBytes)
}

type ProxyHandshakeFullDelay struct {
	Category base.ICategory
}

func NewProxyHandshakeFullDelay() *ProxyHandshakeFullDelay {
	s := &ProxyHandshakeFullDelay{}
	c := base.NewMetricCategory("proxy_handshake_full_delay")
	s.Category = c

	return s
}

func proxyHandshakeResumeDelayRun(_ *cobra.Command, _ []string) {
	proxyHandshakeResumeDelay.Category.Request().EndBytes(proxyHandshakeResumeDelay.Category.EndBytes)
}

type ProxyHandshakeResumeDelay struct {
	Category base.ICategory
}

func NewProxyHandshakeResumeDelay() *ProxyHandshakeResumeDelay {
	s := &ProxyHandshakeResumeDelay{}
	c := base.NewMetricCategory("proxy_handshake_resume_delay")
	s.Category = c

	return s
}
