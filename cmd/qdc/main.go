package main

import (
	"fmt"
	"os"

	"github.com/jonpalmisc/qdc/quartz"
	flag "github.com/spf13/pflag"
)

const Version string = "0.2.2"

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
	u := "Quartz Display Configurator " + Version + "\n"
	u += "Copyright (c) 2019-2020 Jon Palmisciano\n\n"
	u += "Usage:\n"
	u += "  qdc [options]\n"
	u += "\n"
	u += "Options:\n"
	u += "  -d <display>    the display to adjust (defaults to main)\n"
	u += "  -r <size>       set the display resolution\n"
	u += "  -x <target>     mirror the  display to target\n"
	u += "  -h              show help & usage\n"

	_, err := fmt.Fprintln(os.Stderr, u)
	if err != nil {
		panic(err)
	}

	os.Exit(1)
}

func main() {
	idx := flag.IntP("display", "d", 0, "")
	mirror := flag.IntP("mirror", "x", -1, "")
	res := flag.StringP("resolution", "r", "", "")

	flag.Usage = ShowUsage
	flag.Parse()

	// Make sure at least one operation has been requested.
	if *mirror == -1 && *res == "" {
		ShowUsage()
	}

	// Get all displays and ensure the selected display is in bounds.
	displays := quartz.OnlineDisplays()
	if *idx >= len(displays) {
		PrintFatal("display index out of bounds")
	}

	display := quartz.OnlineDisplays()[*idx]

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
