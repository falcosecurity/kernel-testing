package main

import (
	"encoding/json"
	"flag"
	"github.com/olekukonko/tablewriter"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	rootFolder *string
	outputFile *string
)

type testResult struct {
	Rc matrixEntryResult `json:"rc"`
}

type matrixEntryResult int

const (
	matrixEntryResultOK matrixEntryResult = iota
	matrixEntryResultFail
	matrixEntryResultSkip
)

type matrixEntry map[string]matrixEntryResult

type matrixOutput struct {
	entries  map[string]matrixEntry
	testList map[string]struct{}
}

func init() {
	rootFolder = flag.String("root-folder", "~/ansible_output", "ansible output root folder")
	outputFile = flag.String("output-file", "matrix.md", "output file where the generated matrix is stored")
}

func loadTestResult(path string) matrixEntryResult {
	file, _ := os.ReadFile(path)

	res := testResult{}

	_ = json.Unmarshal(file, &res)
	if res.Rc != 0 {
		return matrixEntryResultFail
	}
	return matrixEntryResultOK
}

func (m matrixOutput) addTestResult(path string) {
	subPaths := strings.Split(path, "/")
	testName := strings.TrimSuffix(subPaths[len(subPaths)-1], ".json")
	machineName := subPaths[len(subPaths)-2]

	if _, ok := m.entries[machineName]; !ok {
		m.entries[machineName] = make(map[string]matrixEntryResult)
	}
	matrixentry := m.entries[machineName]
	matrixentry[testName] = loadTestResult(path)
	m.entries[machineName] = matrixentry
}

func (m matrixOutput) Store() {
	fW, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fW.Close()

	table := tablewriter.NewWriter(fW)

	headers := []string{"Kernel"}
	for testName, _ := range m.testList {
		headers = append(headers, testName)
	}

	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	for kernel, tests := range m.entries {
		data := make([]string, len(tests)+1)
		data[0] = kernel
		var idx = 1
		for testName, _ := range m.testList {
			testRes := matrixEntryResultSkip
			if _, ok := tests[testName]; ok {
				testRes = tests[testName]
			}
			switch testRes {
			case matrixEntryResultOK:
				data[idx] = "🟢"
			case matrixEntryResultFail:
				data[idx] = "🟡"
			case matrixEntryResultSkip:
				data[idx] = " "
			}
			idx++
		}
		table.Append(data)
	}
	table.Render() // Send output
}

func main() {
	flag.Parse()

	matrix := matrixOutput{
		entries:  make(map[string]matrixEntry),
		testList: make(map[string]struct{}),
	}

	err := filepath.WalkDir(*rootFolder, func(path string, d fs.DirEntry, err error) error {
		if d.Type() == 0 { // regular file
			matrix.addTestResult(path)

			testName := strings.TrimSuffix(d.Name(), ".json")
			if _, ok := matrix.testList[testName]; !ok {
				matrix.testList[testName] = struct{}{}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to walk directory: %s", err)
	}

	matrix.Store()
}
