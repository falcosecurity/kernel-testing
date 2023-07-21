package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
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
	Rc             int    `json:"rc"`
	Skipped        bool   `json:"skipped"`
	StdErr         string `json:"stderr"`
	Msg            string `json:"msg"`
	FalseCondition string `json:"false_condition"`
}

type matrixEntry map[string]testResult

type matrixOutput struct {
	entries  map[string]matrixEntry
	testList map[string]time.Time
}

type matrixErrorReportKey struct {
	Kernel string
	Test   string
	Res    testResult
}

func (m matrixErrorReportKey) ToMDSection() string {
	key := fmt.Sprint("#" + m.Kernel + "-" + m.Test)
	// "." is not available, ie:
	// #archlinux-5.18-build-kernel-module should become
	// #archlinux-518-build-kernel-module
	return strings.Replace(key, ".", "", -1)
}

func init() {
	rootFolder = flag.String("root-folder", "~/ansible_output", "ansible output root folder")
	outputFile = flag.String("output-file", "matrix.md", "output file where the generated matrix is stored")
}

func loadTestResult(path string) testResult {
	file, _ := os.ReadFile(path)
	res := testResult{}
	err := json.Unmarshal(file, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (m matrixOutput) addTestResult(path string) {
	subPaths := strings.Split(path, "/")
	testName := strings.TrimSuffix(subPaths[len(subPaths)-1], ".json")
	machineName := subPaths[len(subPaths)-2]

	if _, ok := m.entries[machineName]; !ok {
		m.entries[machineName] = make(map[string]testResult)
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

func writeMDCodeBlock(w io.StringWriter, block string) {
	w.WriteString("```\n")
	w.WriteString(block + "\n")
	w.WriteString("```\n")
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

	// list of tests that need to be reported to user
	// because either failed or skipped
	toBeReported := make([]matrixErrorReportKey, 0)

	for _, kernel := range kernels {
		tests := m.entries[kernel]
		data := make([]string, len(headers))
		for idx, testName := range headers {
			if idx == 0 {
				data[idx] = kernel
				continue
			}
			testRes := tests[testName]
			mErrKey := matrixErrorReportKey{
				Kernel: kernel,
				Test:   testName,
				Res:    testRes,
			}
			if testRes.Skipped {
				data[idx] = fmt.Sprintf("[🟡](%s)", mErrKey.ToMDSection())
				toBeReported = append(toBeReported, mErrKey)
			} else if testRes.Rc != 0 {
				data[idx] = fmt.Sprintf("[❌](%s)", mErrKey.ToMDSection())
				toBeReported = append(toBeReported, mErrKey)
			} else {
				data[idx] = "🟢"
			}
			idx++
		}
		table.Append(data)
	}
	table.Render() // Send output

	// After the table, append all the failed/skipped tests
	// outputs, each as a separate section,
	// to allow users to quickly heck them.
	fW.WriteString("\n\n")
	for _, mErrReport := range toBeReported {
		fW.WriteString(mErrReport.ToMDSection() + "\n\n")
		if mErrReport.Res.Skipped {
			fW.WriteString("Skipped Condition:\n")
			writeMDCodeBlock(fW, mErrReport.Res.FalseCondition)
		} else {
			fW.WriteString("Msg:\n")
			writeMDCodeBlock(fW, mErrReport.Res.Msg)
			fW.WriteString("Err:\n")
			if mErrReport.Res.StdErr != "" {
				writeMDCodeBlock(fW, mErrReport.Res.StdErr)
			} else {
				writeMDCodeBlock(fW, fmt.Sprintf("Exit Code: %d", mErrReport.Res.Rc))
			}
		}
		fW.WriteString("\n")
	}
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
