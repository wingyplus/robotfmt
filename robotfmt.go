package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"text/tabwriter"
)

var (
	write = flag.Bool("w", false, "write result to file instead of stdout")
	debug = flag.Bool("d", false, "enable debug mode")
)

func readFile(filename string) string {
	b, _ := ioutil.ReadFile(filename)
	return string(b)
}

func format(w io.Writer, content string) {
	var (
		mode    uint = 0
		padchar byte = ' '
	)
	if *debug {
		mode = tabwriter.Debug
		padchar = '-'
	}

	tabWriter := tabwriter.NewWriter(w, 0, 4, 4, padchar, mode)
	fmt.Fprint(tabWriter, content)
	tabWriter.Flush()
}

func main() {
	flag.Parse()

	filename := flag.Args()[0]

	re := regexp.MustCompile(" {2,}")
	content := re.ReplaceAllString(readFile(filename), "\t")

	var buf bytes.Buffer
	format(&buf, content)

	if *write {
		ioutil.WriteFile(filename, buf.Bytes(), 0644)
	} else {
		io.Copy(os.Stdout, &buf)
	}

}
