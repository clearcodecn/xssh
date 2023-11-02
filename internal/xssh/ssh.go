package xssh

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/clearcodecn/xssh/config"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"os"
	"os/exec"
)

func Ssh() error {
	var options []string
	prompt := &survey.Select{
		Message: "Choose a server to login:",
	}

	var models []*XSSHModel
	db := config.GetDB()
	it := db.NewIterator(util.BytesPrefix([]byte(keyPrefix)), nil)
	for it.Next() {
		val := it.Value()
		var data XSSHModel
		err := json.Unmarshal(val, &data)
		if err != nil {
			log.Println("[warn] parse json failed:", err.Error())
			continue
		}
		tab := fmt.Sprintf("[%s]\t%s@%s\t-p\t%s\t-i\t%s", data.Alias, data.Username, data.Host, data.Port, data.SshKey)
		options = append(options, tab)
		models = append(models, &data)
	}
	prompt.Options = options
	var answer string
	if err := survey.AskOne(prompt, &answer); err != nil {
		return err
	}
	var mdl *XSSHModel
	for i, opt := range options {
		if opt == answer {
			mdl = models[i]
			break
		}
	}
	if mdl == nil {
		return nil
	}
	config.CloseDB(db)

	cmd := exec.Command("ssh",
		fmt.Sprintf("%s@%s", mdl.Username, mdl.Host),
		"-p", mdl.Port,
		"-i", mdl.SshKey)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
