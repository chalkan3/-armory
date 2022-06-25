package nodecommand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"

	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"
)

type CreateCommand struct {
	root *cobra.Command
}

func NewCreateCommand() *CreateCommand {
	return &CreateCommand{
		root: &cobra.Command{
			Aliases: []string{"create"},
			Use:     "create",
			Short:   "create a new node",
		},
	}
}

func (nc *CreateCommand) Command() *cobra.Command {
	nc.root.Run = nc.Create
	nc.root.Flags().String("name", "", "set node name")
	nc.root.Flags().String("cluster", "", "cluster target")
	nc.root.Flags().String("private-ip", "", "set node private ip")
	nc.root.Flags().String("type", "", "set node type (master, worker)")
	nc.root.Flags().Bool("primary", false, "set node type (master, worker)")
	nc.root.Flags().String("lb-ip", "", "set node type (master, worker)")
	nc.root.Flags().StringP("file", "f", "", "set node type (master, worker)")

	return nc.root
}

func (nc *CreateCommand) Create(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}
	privateIP, err := cmd.Flags().GetString("private-ip")
	if err != nil {
		panic(err)
	}
	types, err := cmd.Flags().GetString("type")
	if err != nil {
		panic(err)
	}

	lbIP, err := cmd.Flags().GetString("lb-ip")
	if err != nil {
		panic(err)
	}

	primary, err := cmd.Flags().GetBool("primary")
	if err != nil {
		panic(err)
	}

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		panic(err)
	}

	cluster, err := cmd.Flags().GetString("cluster")
	if err != nil {
		panic(err)
	}

	manifest := nodesv1beta1.KubeNodes{
		ApiVersion: "nodesv1beta1",
		Spec: &nodesv1beta1.Spec{
			ClusterName: cluster,
			Node: &nodesv1beta1.Node{
				Name:    name,
				Types:   types,
				Primary: primary,
				Network: &nodesv1beta1.Network{
					PrivateIP:      privateIP,
					LoadBalancerIP: lbIP,
				},
			},
		},
		Metadata: &nodesv1beta1.Metadata{
			Name: name,
		},
	}

	if file != "" {

		path, _ := filepath.Abs(file)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		yaml.Unmarshal(content, &manifest)

	}

	payload, err := json.Marshal(map[string]nodesv1beta1.KubeNodes{
		"manifest": manifest,
	})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:9003/v1/kubernetes/nodes", "application/json",
		bytes.NewBuffer(payload))

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		log.Println(fmt.Sprintf("resource %v was created", name))
	}

}
