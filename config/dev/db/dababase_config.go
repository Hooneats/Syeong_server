package db

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Hooneats/Syeong_server/common/ciper"
	"os"
	"time"
)

type DB struct {
	URI        string
	DBName     string
	User       string
	PWD        string
	BackupPath string
}

func (d *DB) DecryptFields() error {
	block := ciper.GetCipherBlock()
	URI, err := ciper.AESDecrypt(block, d.URI)
	if err != nil {
		return err
	}
	d.URI = URI

	name, err := ciper.AESDecrypt(block, d.DBName)
	if err != nil {
		return err
	}
	d.DBName = name
	return nil
}

func WriteBackup(fPath string, T any) error {
	data, err := json.MarshalIndent(T, "", "    ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(fPath, 0700); err != nil {
		return err
	}

	path := fPath + time.Now().Format("2006-01-02") + ".txt"
	file := fmt.Sprintf(path)
	f, err := os.OpenFile(
		file, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0700),
	)
	defer f.Close()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	if _, err = fmt.Fprint(w, string(data)+"\n"); err != nil {
		return err
	}

	if err = w.Flush(); err != nil {
		return err
	}

	return nil
}
