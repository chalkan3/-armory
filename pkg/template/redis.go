package tpl

import (
	"log"
	"os"
	v1beta1 "scheduler/pkg/manifest/redis/v1beta1"
	"text/template"
)

type RedisTemplate struct {
	tpl *template.Template
}

func NewRedisTemplate() *RedisTemplate { return new(RedisTemplate) }

func (t *RedisTemplate) Parse(path string) *RedisTemplate {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return t
	}
	t.tpl = tpl
	return t
}

func (t *RedisTemplate) Execute(manifest *v1beta1.Redis, payload interface{}) {
	path := "/Users/igor.rodrigues/Documents/study/ansible/scheduler/vms/database/redis/" + manifest.Spec.ClusterName + "/" + manifest.Metadata.ID
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
