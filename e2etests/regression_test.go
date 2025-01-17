package e2etests

import (
	"testing"
)

func testRegression(t *testing.T) {
	tcs := []testCase{
		{
			name: "dagre_special_ids",
			script: `
ninety\nnine
eighty\reight
seventy\r\nseven
a\\yode -> there
a\\"ode -> there
a\\node -> there
`,
		},
		{
			name: "empty_sequence",
			script: `
A: hello {
  shape: sequence_diagram
}

B: goodbye {
  shape: sequence_diagram
}

A->B
`,
		},
	}

	runa(t, tcs)
}
