package xssh

import (
	"encoding/json"
	"errors"
	"github.com/clearcodecn/xssh/config"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	keyPrefix = "xssh-"
)

func Add(model *XSSHModel, overWrite bool) error {
	var write bool
	db := config.GetDB()
	defer config.CloseDB(db)
	
	data, err := db.Get([]byte(keyPrefix+model.Key()), nil)
	if err != nil {
		if err != leveldb.ErrNotFound {
			return err
		}
		write = true
	}
	if err != nil {
		if !overWrite && !write {
			return errors.New("记录已存在")
		}
	}

	data, _ = json.Marshal(model)
	err = db.Put([]byte(keyPrefix+model.Key()), data, nil)
	if err != nil {
		return err
	}
	return nil
}
