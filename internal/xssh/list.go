package xssh

import (
	"encoding/json"
	"fmt"
	"github.com/clearcodecn/xssh/config"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

func List(subStr string) {
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
		fmt.Println(tab)
	}
}
