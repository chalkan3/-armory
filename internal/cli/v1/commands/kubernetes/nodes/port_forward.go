package nodecommand

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"
	ssh "scheduler/pkg/ssh"

	"github.com/spf13/cobra"
)

type PortFowardCommand struct {
	root *cobra.Command
}

func NewPortFowardCommand() *PortFowardCommand {
	return &PortFowardCommand{
		root: &cobra.Command{
			Aliases: []string{"pf"},
			Use:     "port-forward",
			Short:   "ssh a node",
		},
	}
}

func (nc *PortFowardCommand) Command() *cobra.Command {
	nc.root.Run = nc.Create
	nc.root.Flags().String("name", "", "set node name")
	nc.root.Flags().StringP("addr", "a", "", "set node name")
	nc.root.Flags().StringP("port", "p", "", "set node name")

	return nc.root
}

func (nc *PortFowardCommand) Create(cmd *cobra.Command, args []string) {
	var manifests map[string][]nodesv1beta1.KubeNodes

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}

	port, err := cmd.Flags().GetString("port")
	if err != nil {
		panic(err)
	}

	addr, err := cmd.Flags().GetString("addr")
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
				ssh.NewSSH().Connect("vagrant", "vagrant", manifest.Spec.Node.Network.PrivateIP, port).PortForward(addr)
				break
			}
		}
	}

}
