package schedulelistners

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"
	tpl "scheduler/pkg/template"
)

func Worker(data interface{}) {
	manifest := data.(*nodesv1beta1.KubeNodes)

	payload := struct {
		Values      *nodesv1beta1.KubeNodes
		AnsiblePath string
	}{
		Values:      manifest,
		AnsiblePath: "../../../ansible/kubernetes-setup/node-playbook.yml",
	}
	tpl.NewVagrantNodeTemplate().Parse("/Users/igor.rodrigues/Documents/study/ansible/scheduler/templates/Vagrantfile.tpl").Execute(manifest, payload)
	command := "vagrant"
	cmd := exec.Command(command, "up")

	cmd.Dir = "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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

func MasterPrimary(data interface{}) {
	manifest := data.(*nodesv1beta1.KubeNodes)

	payload := struct {
		Values      *nodesv1beta1.KubeNodes
		AnsiblePath string
	}{
		Values:      manifest,
		AnsiblePath: "../../../ansible/kubernetes-setup/master-playbook.yml",
	}
	tpl.NewVagrantNodeTemplate().Parse("/Users/igor.rodrigues/Documents/study/ansible/scheduler/templates/Vagrantfile.tpl").Execute(manifest, payload)
	command := "vagrant"
	cmd := exec.Command(command, "up")

	cmd.Dir = "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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

func Master(data interface{}) {
	manifest := data.(*nodesv1beta1.KubeNodes)

	payload := struct {
		Values      *nodesv1beta1.KubeNodes
		AnsiblePath string
	}{
		Values:      manifest,
		AnsiblePath: "../../../ansible/kubernetes-setup/master-replica-playbook.yml",
	}
	tpl.NewVagrantNodeTemplate().Parse("/Users/igor.rodrigues/Documents/study/ansible/scheduler/templates/Vagrantfile.tpl").Execute(manifest, payload)
	command := "vagrant"
	cmd := exec.Command(command, "up")

	cmd.Dir = "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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

func DeleteNode(data interface{}) {
	manifest := data.(*nodesv1beta1.KubeNodes)
	command := "vagrant"
	cmd := exec.Command(
		command,
		"destroy",
		manifest.Spec.Node.Name,
		"-f",
	)

	cmd.Dir = "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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
