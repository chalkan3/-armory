package postgres

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

type ListCommand struct {
	root *cobra.Command
}

func NewListCommand() *ListCommand {
	return &ListCommand{
		root: &cobra.Command{
			Aliases: []string{"list"},
			Use:     "list",
			Short:   "list nodes",
		},
	}
}

func (nc *ListCommand) Command() *cobra.Command {
	nc.root.Run = nc.Create

	return nc.root
}

func (nc *ListCommand) Create(cmd *cobra.Command, args []string) {
	var manifests map[string][]nodesv1beta1.KubeNodes
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
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"id", "cluster", "name", "type", "private ip", "primary", "created at"})
		var rows []table.Row
		for _, manifest := range manifests["manifest"] {
			rows = append(rows, table.Row{
				strings.Replace(manifest.Metadata.ID, "nodes-", "", 1),
				manifest.Spec.ClusterName,
				manifest.Metadata.Name,
				manifest.Spec.Node.Types,
				manifest.Spec.Node.Network.PrivateIP,
				manifest.Spec.Node.Primary,
				manifest.Metadata.CreatedAT,
			})
		}
		t.AppendRows(rows)

		t.AppendSeparator()
		t.Render()

	}

}
