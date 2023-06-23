package xssh

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/clearcodecn/xssh/config"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

func Del() error {
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
		tab := fmt.Sprintf("[%s]\t%s@%s\t-p\t%s\t-i\t%s\n", data.Alias, data.Username, data.Host, data.Port, data.SshKey)
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

	return db.Delete([]byte(mdl.Key()), nil)
}
