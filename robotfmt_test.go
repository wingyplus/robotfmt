package main

import (
	"bytes"
	"testing"
)

var tests = []struct {
	name string
	in   string
	out  string
}{
	{
		name: "format code",
		in: `
*** Settings ***
Library  CollectionLibrary
`,
		out: `
*** Settings ***
Library    CollectionLibrary
`,
	},
	{
		name: "format empty column",
		in: `
*** Settings ***
Library  Selenium2Library
Library		CollectionLibrary
Test Template	Work Around
`,
		out: `
*** Settings ***
Library          Selenium2Library
Library          CollectionLibrary
Test Template    Work Around
`,
	},
	{
		name: "format tab that follow by space",
		in: `
*** Test Cases ***	fieldA	fieldB
Case 1	 1	2
`,
		out: `
*** Test Cases ***    fieldA    fieldB
Case 1                1         2
`,
	},
	{
		name: "format multiple tabs to 4 padding spaces",
		in: `
*** Test Cases ***
Test Multiple Tabs To 4 Padding Spaces
			Keyword Step
`,
		out: `
*** Test Cases ***
Test Multiple Tabs To 4 Padding Spaces
    Keyword Step
`,
	},
}

func TestFormat(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			format(&buf, test.in)
			if s := buf.String(); s != test.out {
				t.Error("output from format() is", s)
			}
		})
	}
}
