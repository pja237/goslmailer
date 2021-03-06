package cmdline

import (
	"errors"
	"flag"
)

// CmdArgs holds currently supported command line parameters.
type CmdArgs struct {
	CfgFile *string
	Version *bool
}

// NewCmdArgs return the CmdArgs structure built from command line parameters.
// `prog` string is used to build the default /etc/slurm/`prog`.conf configFile string.
func NewCmdArgs(prog string) (*CmdArgs, error) {
	c := new(CmdArgs)

	c.CfgFile = flag.String("c", "/etc/slurm/"+prog+".conf", "Specify configuration file to use")
	c.Version = flag.Bool("v", false, "Display version")
	flag.Parse()
	if !flag.Parsed() {
		return nil, errors.New("failed to parse command line flags")
	}

	return c, nil
}
