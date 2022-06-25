package hacommand

import (
	"github.com/spf13/cobra"
)

type RawCommand struct {
	root *cobra.Command
}

func NewRawCommand() *RawCommand {
	return &RawCommand{
		root: &cobra.Command{
			Aliases: []string{"raw"},
			Use:     "raw",
			Short:   "configure kubernetes nodes",
		},
	}
}

func (nc *RawCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		NewCreateCommand().Command(),
	)
	return nc.root
}
