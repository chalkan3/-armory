package tpl

import (
	"log"
	"os"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"
	"text/template"
)

type VagrantNodeTemplate struct {
	tpl *template.Template
}

func NewVagrantNodeTemplate() *VagrantNodeTemplate { return new(VagrantNodeTemplate) }

func (t *VagrantNodeTemplate) Parse(path string) *VagrantNodeTemplate {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return t
	}
	t.tpl = tpl
	return t
}

func (t *VagrantNodeTemplate) Execute(manifest *nodesv1beta1.KubeNodes, payload interface{}) {
	path := "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
	os.MkdirAll(path, 0777)

	f, err := os.Create(path + "/Vagrantfile")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.tpl.Execute(f, payload)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}
