package lbcommand

import (
	kafkacmd "scheduler/internal/cli/commands/stream/kafka"

	"github.com/spf13/cobra"
)

type StreamCommand struct {
	root *cobra.Command
}

func NewStreamCommand() *StreamCommand {
	return &StreamCommand{
		root: &cobra.Command{
			Aliases: []string{"st"},
			Use:     "stream",
			Short:   "configure stream types like (kafka)",
		},
	}
}

func (nc *StreamCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		kafkacmd.NewKafkaCommand().Register(),
	)
	return nc.root
}
