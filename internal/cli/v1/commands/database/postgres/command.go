package postgres

import (
	"github.com/spf13/cobra"
)

type PostgresClusterCommand struct {
	root *cobra.Command
}

func NewPostgresClusterCommand() *PostgresClusterCommand {
	return &PostgresClusterCommand{
		root: &cobra.Command{
			Aliases: []string{"pg"},
			Use:     "postgres",
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
