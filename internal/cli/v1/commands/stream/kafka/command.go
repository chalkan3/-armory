package hacommand

import (
	"github.com/spf13/cobra"
)

type KafkaCommand struct {
	root *cobra.Command
}

func NewKafkaCommand() *KafkaCommand {
	return &KafkaCommand{
		root: &cobra.Command{
			Aliases: []string{"kafka"},
			Use:     "kafka",
			Short:   "configure kafka",
		},
	}
}

func (nc *KafkaCommand) Register() *cobra.Command {
	nc.root.AddCommand(
		NewCreateCommand().Command(),
	)
	return nc.root
}
