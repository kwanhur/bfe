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
	"os"
)

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

const (
	errFmt = red + "%v" + red + "\n"
)

type printer struct {
}

func (p *printer) Error(err error) {
	_, _ = fmt.Fprintf(os.Stderr, errFmt, err)
}

func (p *printer) Fatal(err error) {
	_, _ = fmt.Fprintf(os.Stderr, errFmt, err)
	os.Exit(1)
}

func (p *printer) Println(a ...interface{}) {
	fmt.Println(a...)
}

func (p *printer) Print(a ...interface{}) {
	fmt.Print(a...)
}

var Logger = printer{}
