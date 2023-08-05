package boa

import (
	"flag"
	"fmt"
)

type Cmd interface {
	Init(args []string) error
	Called() bool
	Run()
}

type Command struct {
	Name        string
	Flags       *flag.FlagSet
	Execute     func(cmd *Command, args []string)
	Subcommands []*Command
}

func (c *Command) Init(args []string) error {
	return c.Flags.Parse(args)
}

func (c *Command) Called() bool {
	return c.Flags.Parsed()
}

func (c *Command) Run() {
	fmt.Printf("running command %s\n", c.Name)
	fmt.Printf("args: %v\n", c.Flags.Args())

	if len(c.Flags.Args()) == 0 {
		c.Execute(c, c.Flags.Args())
		return
	}

	for _, cmd := range c.Subcommands {
		if cmd.Name == c.Flags.Arg(0) {
			cmd.Init(c.Flags.Args()[1:])
			cmd.Run()
		}
	}

}

func (c *Command) AddCommand(cmd *Command) {
	fmt.Println("adding command")
	c.Subcommands = append(c.Subcommands, cmd)
}
