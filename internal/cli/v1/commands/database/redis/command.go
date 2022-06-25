package postgres

import (
	"github.com/spf13/cobra"
)

type PostgresClusterCommand struct {
	root *cobra.Command
}

func NewRedisCommand() *PostgresClusterCommand {
	return &PostgresClusterCommand{
		root: &cobra.Command{
			Aliases: []string{"rd"},
			Use:     "redis",
			Short:   "configure cluster",
		},
	}
}

func (nc *PostgresClusterCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		NewCreateCommand().Command(),
	)
	return nc.root
}
