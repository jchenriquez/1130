package minimumcosttree

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input  []int `json:"input"`
	Output int   `json:"output"`
}

func TestMctFromLeafValues(test *testing.T) {
	f, err := os.Open("./tests.json")

	if err != nil {
		test.Error(err)
	}

	defer f.Close()

	var tests map[string]Test
	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	for {
		err := decoder.Decode(&tests)

		if err == nil {
			for name, tst := range tests {
				test.Run(name, func(st *testing.T) {
					if MctFromLeafValues(tst.Input) != tst.Output {
						st.Errorf("Test %v", tst)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			test.Error(err)
		}
	}
}
