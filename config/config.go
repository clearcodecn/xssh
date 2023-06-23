package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"path/filepath"
)

var (
	confPath = ".xssh"
	db       *leveldb.DB
)

func init() {
	dir, _ := homedir.Dir()
	dbPath := filepath.Join(dir, confPath)
	os.MkdirAll(dbPath, 0777)

	db, _ = leveldb.OpenFile(dbPath, nil)
}

func GetDB() *leveldb.DB {
	return db
}
