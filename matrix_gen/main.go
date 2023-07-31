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
