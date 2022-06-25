package postgres

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	v1beta "scheduler/pkg/manifest/redis/v1beta1"

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

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		panic(err)
	}

	cluster, err := cmd.Flags().GetString("cluster")
	if err != nil {
		panic(err)
	}

	manifest := v1beta.Redis{
		ApiVersion: "redis/v1beta",
		Spec: &v1beta.Spec{
			ClusterName: cluster,
			Configuration: &v1beta.Configuration{
				Name: name,
				Network: &v1beta.Network{
					PrivateIP: privateIP,
				},
			},
		},
		Metadata: &v1beta.Metadata{
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

	payload, err := json.Marshal(map[string]v1beta.Redis{
		"manifest": manifest,
	})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:9003/v1/database/redis", "application/json",
		bytes.NewBuffer(payload))

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		log.Println(fmt.Sprintf("resource %v was created", name))
	}

}
