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

package main

import (
	"flag"
	matrix "github.com/falcosecurity/kernel-testing/matrix_gen/pkg/matrix"
	"log"
)

var (
	rootFolder *string
	outputFile *string
)

func init() {
	rootFolder = flag.String("root-folder", "~/ansible_output", "ansible output root folder")
	outputFile = flag.String("output-file", "matrix.md", "output file where the generated matrix is stored")
}

func main() {
	flag.Parse()

	outputMatrix := matrix.NewOutput()
	err := outputMatrix.Loop(*rootFolder)
	if err != nil {
		log.Fatalf("failed to loop directory %s: %s", *rootFolder, err)
	}
	outputMatrix.Store(*outputFile)
}
