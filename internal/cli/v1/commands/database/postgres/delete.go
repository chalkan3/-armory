package postgres

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"

	"github.com/spf13/cobra"
)

type DeleteCommand struct {
	root *cobra.Command
}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{
		root: &cobra.Command{
			Aliases: []string{"del"},
			Use:     "delete",
			Short:   "ssh a node",
		},
	}
}

func (nc *DeleteCommand) Command() *cobra.Command {
	nc.root.Run = nc.Create
	nc.root.Flags().String("name", "", "set node name")

	return nc.root
}

func (nc *DeleteCommand) Create(cmd *cobra.Command, args []string) {
	var manifests map[string][]nodesv1beta1.KubeNodes

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}

	resp, err := http.Get("http://localhost:9003/v1/kubernetes/nodes")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &manifests)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {

		for _, manifest := range manifests["manifest"] {
			if manifest.Spec.Node.Name == name {
				// Create client
				client := &http.Client{}

				// Create request
				req, err := http.NewRequest("DELETE", "http://localhost:9003/v1/kubernetes/nodes/"+manifest.Metadata.ID, nil)
				if err != nil {
					fmt.Println(err)
					return
				}

				// Fetch Request
				_, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				break
			}
		}
	}

}
