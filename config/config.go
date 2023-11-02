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
	dbPath   string
)

func init() {
	dir, _ := homedir.Dir()
	dbPath = filepath.Join(dir, confPath)
	os.MkdirAll(dbPath, 0777)
}

func GetDB() *leveldb.DB {
	db, _ = leveldb.OpenFile(dbPath, nil)
	return db
}

func CloseDB(db *leveldb.DB) {
	db.Close()
}
