package boa

import (
	"flag"
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := &Command{
		Flags: flag.NewFlagSet("test", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			fmt.Println("Hello World!")
		},
	}
	cmd.Flags.Parse([]string{})
	cmd.Run()
}
