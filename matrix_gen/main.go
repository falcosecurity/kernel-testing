package main

import (
	"encoding/json"
	"flag"
	"github.com/olekukonko/tablewriter"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	rootFolder *string
	outputFile *string
)

type testResult struct {
	Rc      matrixEntryResult `json:"rc"`
	Skipped bool              `json:"skipped"`
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
	testList map[string]time.Time
}

func init() {
	rootFolder = flag.String("root-folder", "~/ansible_output", "ansible output root folder")
	outputFile = flag.String("output-file", "matrix.md", "output file where the generated matrix is stored")
}

func loadTestResult(path string) matrixEntryResult {
	file, _ := os.ReadFile(path)

	res := testResult{}

	_ = json.Unmarshal(file, &res)
	if res.Skipped {
		return matrixEntryResultSkip
	}
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

func (m matrixOutput) loadSortTestByModTime() []string {
	type kv struct {
		Key   string
		Value time.Time
	}

	ss := make([]kv, 0, len(m.testList))
	for k, v := range m.testList {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value.Before(ss[j].Value)
	})

	testList := make([]string, 0, len(m.testList))
	for _, val := range ss {
		testList = append(testList, val.Key)
	}
	return testList
}

func (m matrixOutput) Store() {
	fW, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fW.Close()

	// Load sorted by mod time test list, so that they appear
	// in correct order
	testList := m.loadSortTestByModTime()

	headers := []string{"Kernel"}
	for _, testName := range testList {
		headers = append(headers, testName)
	}

	table := tablewriter.NewWriter(fW)
	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	// Sort by kernel
	kernels := make([]string, 0, len(m.entries))
	for k := range m.entries {
		kernels = append(kernels, k)
	}
	sort.Strings(kernels)

	for _, kernel := range kernels {
		tests := m.entries[kernel]
		data := make([]string, len(headers))
		for idx, testName := range headers {
			if idx == 0 {
				data[idx] = kernel
				continue
			}
			// This should never happen; leave this in case.
			testRes := matrixEntryResultSkip
			if _, ok := tests[testName]; ok {
				testRes = tests[testName]
			}
			switch testRes {
			case matrixEntryResultOK:
				data[idx] = "🟢"
			case matrixEntryResultFail:
				data[idx] = "❌"
			case matrixEntryResultSkip:
				data[idx] = "🟡"
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
		testList: make(map[string]time.Time),
	}

	err := filepath.WalkDir(*rootFolder, func(path string, d fs.DirEntry, err error) error {
		if d.Type() == 0 { // regular file
			matrix.addTestResult(path)

			testName := strings.TrimSuffix(d.Name(), ".json")
			if _, ok := matrix.testList[testName]; !ok {
				info, _ := d.Info()
				matrix.testList[testName] = info.ModTime()
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to walk directory: %s", err)
	}

	matrix.Store()
}
