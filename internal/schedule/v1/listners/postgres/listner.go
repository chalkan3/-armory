package listners

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	v1beta "scheduler/pkg/manifest/postgres/v1beta1"
	tpl "scheduler/pkg/template"
)

func Postgres(data interface{}) {
	manifest := data.(*v1beta.Postgres)

	payload := struct {
		Values      *v1beta.Postgres
		AnsiblePath string
	}{
		Values:      manifest,
		AnsiblePath: "../../../../../ansible/postgres-setup/setup.yml",
	}
	tpl.NewPostgresTemplate().Parse("/Users/igor.rodrigues/Documents/study/ansible/scheduler/templates/postgres.tpl").Execute(manifest, payload)
	command := "vagrant"
	cmd := exec.Command(command, "up")

	cmd.Dir = "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/database/postgres/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\t > %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return
	}

}
