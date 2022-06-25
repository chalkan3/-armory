package tpl

import (
	"log"
	"os"
	v1beta1 "scheduler/pkg/manifest/postgres/v1beta1"
	"text/template"
)

type PostgresTemplate struct {
	tpl *template.Template
}

func NewPostgresTemplate() *PostgresTemplate { return new(PostgresTemplate) }

func (t *PostgresTemplate) Parse(path string) *PostgresTemplate {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return t
	}
	t.tpl = tpl
	return t
}

func (t *PostgresTemplate) Execute(manifest *v1beta1.Postgres, payload interface{}) {
	path := "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/database/postgres/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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
