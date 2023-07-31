package matrix

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Result struct {
	Rc             int    `json:"rc"`
	Skipped        bool   `json:"skipped"`
	StdErr         string `json:"stderr"`
	Msg            string `json:"msg"`
	FalseCondition string `json:"false_condition"`
}

type Entry map[string]Result

type Output struct {
	entries  map[string]Entry
	testList map[string]time.Time
}

func loadTestResult(path string) Result {
	file, _ := os.ReadFile(path)
	res := Result{}
	err := json.Unmarshal(file, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (m *Output) addTestResult(path string) {
	subPaths := strings.Split(path, "/")
	testName := strings.TrimSuffix(subPaths[len(subPaths)-1], ".json")
	machineName := subPaths[len(subPaths)-2]

	if _, ok := m.entries[machineName]; !ok {
		m.entries[machineName] = make(map[string]Result)
	}
	matrixentry := m.entries[machineName]
	matrixentry[testName] = loadTestResult(path)
	m.entries[machineName] = matrixentry
}

func (m *Output) loadSortTestByModTime() []string {
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

func (m *Output) Store(outputFile string) {
	fW, err := os.Create(outputFile)
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
	toBeReported := make([]ErrorReportKey, 0)

	for _, kernel := range kernels {
		tests := m.entries[kernel]
		data := make([]string, len(headers))
		for idx, testName := range headers {
			if idx == 0 {
				data[idx] = kernel
				continue
			}
			testRes := tests[testName]
			mErrKey := newErrorReportKey(kernel, testName, testRes)
			if testRes.Skipped {
				data[idx] = fmt.Sprintf("[🟡](%s)", mErrKey.ToMDSectionLink())
				toBeReported = append(toBeReported, mErrKey)
			} else if testRes.Rc != 0 {
				data[idx] = fmt.Sprintf("[❌](%s)", mErrKey.ToMDSectionLink())
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
		mErrReport.Dump(fW)
	}
}

func NewOutput() *Output {
	return &Output{
		entries:  make(map[string]Entry),
		testList: make(map[string]time.Time),
	}
}

func (m *Output) Loop(rootFolder string) error {
	err := filepath.WalkDir(rootFolder, func(path string, d fs.DirEntry, err error) error {
		if d.Type() == 0 { // regular file
			m.addTestResult(path)

			testName := strings.TrimSuffix(d.Name(), ".json")
			if _, ok := m.testList[testName]; !ok {
				info, _ := d.Info()
				m.testList[testName] = info.ModTime()
			}
		}
		return nil
	})
	return err
}
