package lbcommand

import (
	rawcommand "scheduler/internal/cli/commands/machine/raw_machine"

	"github.com/spf13/cobra"
)

type Machineommand struct {
	root *cobra.Command
}

func NewMachineommand() *Machineommand {
	return &Machineommand{
		root: &cobra.Command{
			Aliases: []string{"ma"},
			Use:     "machine",
			Short:   "configure machine",
		},
	}
}

func (nc *Machineommand) Register() *cobra.Command {
	nc.root.AddCommand(
		rawcommand.NewCreateCommand().Command(),
	)
	return nc.root
}
