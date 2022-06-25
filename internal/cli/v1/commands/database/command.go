package lbcommand

import (
	postgrescommand "scheduler/internal/cli/commands/database/postgres"
	rediscommand "scheduler/internal/cli/commands/database/redis"

	"github.com/spf13/cobra"
)

type PostgresCommand struct {
	root *cobra.Command
}

func NewPostgresCommand() *PostgresCommand {
	return &PostgresCommand{
		root: &cobra.Command{
			Aliases: []string{"db"},
			Use:     "database",
			Short:   "configure postgres",
		},
	}
}

func (nc *PostgresCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		postgrescommand.NewPostgresClusterCommand().Register(),
		rediscommand.NewRedisCommand().Register(),
	)
	return nc.root
}
