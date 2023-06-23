package xssh

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type XSSHModel struct {
	Port     string `json:"port"`
	Host     string `json:"host"`
	SshKey   string `json:"ssh_key"`
	Username string `json:"username"`
	Alias    string `json:"alias"`
}

func (k *XSSHModel) Key() string {
	data, _ := json.Marshal(k)
	m := md5.New()
	m.Write(data)
	x := m.Sum(nil)
	return fmt.Sprintf("%x", x)
}
