package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
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
	var err error
	db, err = leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *leveldb.DB {
	return db
}
