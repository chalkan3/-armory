package hacommand

import (
	"github.com/spf13/cobra"
)

type HaProxyCommand struct {
	root *cobra.Command
}

func NewHaProxyCommand() *HaProxyCommand {
	return &HaProxyCommand{
		root: &cobra.Command{
			Aliases: []string{"ha"},
			Use:     "ha-proxy",
			Short:   "configure kubernetes nodes",
		},
	}
}

func (nc *HaProxyCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		NewCreateCommand().Command(),
	)
	return nc.root
}
