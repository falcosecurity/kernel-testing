// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package matrix

import (
	"fmt"
	"io"
	"strings"
)

type ErrorReportKey struct {
	Kernel string
	Test   string
	Res    Result
}

func writeMDCodeBlock(w io.StringWriter, block string) {
	w.WriteString("```\n")
	w.WriteString(block + "\n")
	w.WriteString("```\n")
}

// ToMDSection example: archlinux-5.18 build-kernel-module will become
// "# archlinux-5.18 build-kernel-module"
func (m ErrorReportKey) ToMDSection() string {
	return "# " + m.Kernel + " " + m.Test + "\n\n"
}

// ToMDSectionLink example: archlinux-5.18 build-kernel-module will become
// "#archlinux-518-build-kernel-module"
func (m ErrorReportKey) ToMDSectionLink() string {
	key := fmt.Sprint("#" + m.Kernel + "-" + m.Test)
	// "." is not available, ie:
	// #archlinux-5.18-build-kernel-module should become
	// #archlinux-518-build-kernel-module
	return strings.Replace(key, ".", "", -1)
}

func (m ErrorReportKey) Dump(fW io.StringWriter) {
	fW.WriteString(m.ToMDSection())
	if m.Res.Skipped {
		fW.WriteString("Skipped Condition:\n")
		writeMDCodeBlock(fW, m.Res.FalseCondition)
	} else {
		fW.WriteString("Msg:\n")
		writeMDCodeBlock(fW, m.Res.Msg)
		fW.WriteString("Err:\n")
		if m.Res.StdErr != "" {
			writeMDCodeBlock(fW, m.Res.StdErr)
		} else {
			writeMDCodeBlock(fW, fmt.Sprintf("Exit Code: %d", m.Res.Rc))
		}
	}
	fW.WriteString("\n")
}

func newErrorReportKey(kernel, testName string, testRes Result) ErrorReportKey {
	return ErrorReportKey{
		Kernel: kernel,
		Test:   testName,
		Res:    testRes,
	}
}
