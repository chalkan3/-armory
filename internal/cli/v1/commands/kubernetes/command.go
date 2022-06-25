package kubernetescommand

import (
	nodecommand "scheduler/internal/cli/commands/kubernetes/nodes"

	"github.com/spf13/cobra"
)

type KubernetesCommand struct {
	root *cobra.Command
}

func NewKubernetesCommand() *KubernetesCommand {
	return &KubernetesCommand{
		root: &cobra.Command{
			Aliases: []string{"k8s"},
			Use:     "kubernetes",
			Short:   "configure kubernetes",
		},
	}
}

func (nc *KubernetesCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		nodecommand.NewNodeCommand().Register(),
	)
	return nc.root
}
