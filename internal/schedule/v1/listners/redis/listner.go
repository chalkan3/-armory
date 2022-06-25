package listners

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	v1beta "scheduler/pkg/manifest/redis/v1beta1"
	tpl "scheduler/pkg/template"
)

func Redis(data interface{}) {
	manifest := data.(*v1beta.Redis)

	payload := struct {
		Values      *v1beta.Redis
		AnsiblePath string
	}{
		Values:      manifest,
		AnsiblePath: "../../../../../ansible/redis-setup/setup.yml",
	}
	tpl.NewRedisTemplate().Parse("/Users/igor.rodrigues/Documents/study/ansible/scheduler/templates/redis.tpl").Execute(manifest, payload)
	command := "vagrant"
	cmd := exec.Command(command, "up")

	cmd.Dir = "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/database/redis/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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
