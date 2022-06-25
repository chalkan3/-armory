package lbcommand

import (
	hacommand "scheduler/internal/cli/commands/loadbalancer/haproxy"

	"github.com/spf13/cobra"
)

type LoadBalancerCommand struct {
	root *cobra.Command
}

func NewLoadBalancerCommand() *LoadBalancerCommand {
	return &LoadBalancerCommand{
		root: &cobra.Command{
			Aliases: []string{"lb"},
			Use:     "load-balancer",
			Short:   "configure loadbalancer",
		},
	}
}

func (nc *LoadBalancerCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		hacommand.NewHaProxyCommand().Register(),
	)
	return nc.root
}
