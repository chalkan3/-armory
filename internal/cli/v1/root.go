package rootcommand

import (
	postgrescommand "scheduler/internal/cli/v1/commands/database"
	nodecommand "scheduler/internal/cli/v1/commands/kubernetes"
	lbcommand "scheduler/internal/cli/v1/commands/loadbalancer"
	machinecommand "scheduler/internal/cli/v1/commands/machine"
	streamcommand "scheduler/internal/cli/v1/commands/stream"

	"github.com/spf13/cobra"
)

type RootCommand struct {
	root *cobra.Command
}

func NewRootCommand() *RootCommand {
	return &RootCommand{
		root: &cobra.Command{
			Use:   "managectl",
			Short: "configure kubernetes",
		},
	}
}

func (nc *RootCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		nodecommand.NewKubernetesCommand().Register(),
		lbcommand.NewLoadBalancerCommand().Register(),
		machinecommand.NewMachineommand().Register(),
		postgrescommand.NewPostgresCommand().Register(),
		streamcommand.NewStreamCommand().Register(),
	)
	return nc.root
}
