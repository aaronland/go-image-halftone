package halftone

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var process string
var scale_factor int
var source_uri string
var target_uri string

var extra_transformations multi.MultiCSVString

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("halftone")

	fs.IntVar(&scale_factor, "scale-factor", 2, "")
	fs.StringVar(&process, "process", "atkinson", "...")

	fs.StringVar(&source_uri, "source-uri", "file:///", "")
	fs.StringVar(&target_uri, "target-uri", "file:///", "")
	fs.Var(&extra_transformations, "transformation-uri", "")

	return fs
}
