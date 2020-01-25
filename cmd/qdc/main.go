package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jonpalmisc/qdc/quartz"
	"github.com/spf13/pflag"
)

// PrintFatal prints an error to stderr and terminates the program.
func PrintFatal(msg string) {
	_, err := fmt.Fprintf(os.Stderr, "error: %v\n", msg)
	if err != nil {
		panic(err)
	}

	os.Exit(-1)
}

// ShowUsage prints the program's usage information.
func ShowUsage() {
	usage := "usage: qdc [options] <index>\n" +
		"  -x, --mirror <index>\n" +
		"  -r, --resolution <wxh>"

	_, err := fmt.Fprintln(os.Stderr, usage)
	if err != nil {
		panic(err)
	}

	os.Exit(-1)
}

func main() {
	mirror := pflag.IntP("mirror", "x", -1, "")
	res := pflag.StringP("resolution", "r", "", "")

	pflag.Usage = ShowUsage
	pflag.Parse()

	// Abort and show usage if we haven't received a display index.
	if pflag.NArg() < 1 {
		ShowUsage()
	}

	// Attempt to convert the supplied display index to a string.
	idx, err := strconv.Atoi(pflag.Arg(0))
	if err != nil {
		PrintFatal("invalid display index")
	}

	// Make sure at least one operation has been requested.
	if *mirror == -1 && *res == "" {
		PrintFatal("no operation requested")
	}

	// Get all displays and ensure the selected display is in bounds.
	displays := quartz.OnlineDisplays()
	if idx >= len(displays) {
		PrintFatal("display index out of bounds")
	}

	display := quartz.OnlineDisplays()[idx]

	// Attempt to configure mirroring if it has been requested.
	if *mirror != -1 {
		if *mirror >= len(displays) {
			PrintFatal("mirror target index out of bounds")
		}

		target := displays[*mirror]
		display.MirrorTo(&target)
	}

	// Attempt to apply a resolution if it has been requested.
	if *res != "" {
		match, err := display.FindMode(*res)
		if err != nil {
			PrintFatal(err.Error())
		}

		display.ApplyMode(match)
	}
}
