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

package base

import (
	"fmt"
	"strings"
)

import (
	"github.com/baidu/go-lib/web-monitor/metrics"
	"github.com/parnurzeal/gorequest"
)

import (
	"github.com/bfenetworks/bfe/bfe_util/json"
)

type ICategory interface {
	Name() string
	Usage() string
	Request() *gorequest.SuperAgent
	EndBytes(response gorequest.Response, body []byte, errs []error)
}

// MonitorCategory monitor category
type MonitorCategory struct {
	name   string
	output func(body []byte)
}

func NewMonitorCategory(name string, output func(body []byte)) *MonitorCategory {
	return &MonitorCategory{name: name, output: output}
}

func (c *MonitorCategory) Usage() string {
	return strings.ReplaceAll(c.name, "_", "-")
}

func (c *MonitorCategory) Name() string {
	return c.name
}

func (c *MonitorCategory) SubURL(uri string) string {
	return fmt.Sprintf("%s/%s", baseURL("monitor"), uri)
}

func (c *MonitorCategory) Request() *gorequest.SuperAgent {
	var uri = c.name
	if Diff {
		uri += "_diff"
	}

	r := gorequest.New().SetDebug(Debug)
	r.Get(c.SubURL(uri)).Param("format", Format)

	return r
}

func (c *MonitorCategory) EndBytes(_ gorequest.Response, body []byte, errs []error) {
	if len(errs) != 0 {
		for _, err := range errs {
			Logger.Error(err)
		}
		return
	}

	if c.output != nil {
		c.output(body)
	}
}

// MetricCategory metric category
type MetricCategory struct {
	MonitorCategory
}

func NewMetricCategory(name string) *MetricCategory {
	c := NewMonitorCategory(name, OutputMetricsData)
	return &MetricCategory{*c}
}

func OutputMetricsData(body []byte) {
	switch Format {
	default:
		Logger.Print(string(body))
	case FormatJSON:
		data := metrics.MetricsData{}
		var err error
		if err = json.Unmarshal(body, &data); err != nil {
			Logger.Error(err)
			return
		}

		if Pretty {
			body, err = json.MarshalIndent(data, "", "    ")
			if err != nil {
				Logger.Error(err)
				return
			}
		}
		Logger.Println(string(body))
	}
}

// ConfigCategory configuration category
type ConfigCategory struct {
	name string
}

func NewConfigCategory(name string) *ConfigCategory {
	return &ConfigCategory{name: name}
}

func (c *ConfigCategory) Usage() string {
	return strings.ReplaceAll(c.name, "_", "-")
}

func (c *ConfigCategory) Name() string {
	return c.name
}

func (c *ConfigCategory) SubURL(uri string) string {
	return fmt.Sprintf("%s/%s", baseURL("reload"), uri)
}

func (c *ConfigCategory) Request() *gorequest.SuperAgent {
	r := gorequest.New().SetDebug(Debug)
	r.Get(c.SubURL(c.name)).Param("path", Path)

	return r
}

func (c *ConfigCategory) EndBytes(_ gorequest.Response, body []byte, errs []error) {
	if len(errs) != 0 {
		for _, err := range errs {
			Logger.Error(err)
		}
		return
	}

	c.Output(body)
}

func (c *ConfigCategory) Output(body []byte) {
	var ret = struct {
		Error string `json:"error"`
	}{}

	var err error
	if err = json.Unmarshal(body, &ret); err != nil {
		Logger.Error(err)
		return
	}

	if Pretty {
		body, err = json.MarshalIndent(ret, "", "    ")
		if err != nil {
			Logger.Error(err)
			return
		}
	}
	Logger.Println(string(body))
}
