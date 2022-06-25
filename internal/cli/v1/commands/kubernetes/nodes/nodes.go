package nodecommand

import (
	"github.com/spf13/cobra"
)

type NodeCommand struct {
	root *cobra.Command
}

func NewNodeCommand() *NodeCommand {
	return &NodeCommand{
		root: &cobra.Command{
			Aliases: []string{"node"},
			Use:     "node",
			Short:   "configure kubernetes nodes",
		},
	}
}

func (nc *NodeCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		NewCreateCommand().Command(),
		NewListCommand().Command(),
		NewConnectCommand().Command(),
		NewPortFowardCommand().Command(),
		NewDeleteCommand().Command(),
	)
	return nc.root
}
