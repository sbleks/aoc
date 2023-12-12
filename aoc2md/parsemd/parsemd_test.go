package parsemd_test

import (
	"aoc2md/parsemd"
	"testing"
)

func TestParse(t *testing.T) {

	html := []string{
		"<div><code>I shouldn't be emphasized <em>but I should be</em></code></div>",
		"<div><code><em>I should be all emphasized</em></code></div>",
		"<strong>Important</strong>",
	}

	expect := []string{
		"`I shouldn't be emphasized _but I should be_`",
		"_`I should be all emphasized`_",
		"**Important**",
	}

	converter := parsemd.New()

	for i, str := range html {
		markdown, err := converter.ConvertString(str)

		if err != nil {
			t.FailNow()
		}

		if markdown != expect[i] {
			t.FailNow()
		}
	}

}
