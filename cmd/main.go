package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"

	"github.com/komisan19/tcl/action"
	"github.com/olekukonko/tablewriter"
)

const (
	name    = "sgp"
	version = "v.0.0.1"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
}

func main() {
	var strFlag string
	var showVersion bool

	flag.StringVar(&strFlag, "f", "", "target json file")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("%s: %s (rev: %s)\n", name, version, runtime.Version())
		return
	}

	if strFlag != "" {
		data, err := action.ResultAction(&strFlag)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "URL", "Status"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output

		os.Exit(0)
	} else {
		slog.Warn("not read file")
		os.Exit(1)
	}
}
