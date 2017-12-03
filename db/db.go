package db

import (
	"log"
	"strconv"
	"path/filepath"
	"github.com/boltdb/bolt"
	"github.com/webus/mypass/base"
	"errors"
)

var cfg base.MyPassConfiguration

const passwordBucket = "pass"
const passwordLengthBucket = "pass_len"
const loginBucket = "login"
const loginLengthBucket = "login_len"

func createBucket(DB *bolt.DB, bucketName string) {
	DB.Update(func(tx *bolt.Tx) error {
		_, _ = tx.CreateBucket([]byte(bucketName))
		//FIXME: how process this error ?
		return nil
	})
}

func getDbPath() string {
	return filepath.Join(cfg.DatabaseLocation, cfg.DatabaseName)
}

func InitDatabase() {
	cfg = base.MyPassConfiguration{}
	cfg.InitConfiguration()

	dbPath := getDbPath()
	DB, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	for _, bckt := range []string{
		passwordBucket,
		passwordLengthBucket,
		loginBucket,
		loginLengthBucket,
	} {
		createBucket(DB, bckt)
	}
}

func openDatabase() (*bolt.DB, error) {
	dbPath := getDbPath()
	return bolt.Open(dbPath, 0600, nil)
}

// buckets: login, pass
func UpdateDataBucket(bucketName string, dataName string, dataValue string) error {
	db, err := openDatabase()
	defer db.Close()
	if err != nil {
		return err
	}
	dataValueEnc, dataValueLen := base.EncString(dataValue, cfg.Key)
	if err != nil {
		return err
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(dataName), []byte(dataValueEnc))
		return err
	})
	err = updateDataLenBucket(bucketName+"_len", dataName, strconv.Itoa(dataValueLen), db)
	if err != nil {
		return err
	}
	return nil
}

func updateDataLenBucket(bucketName string, dataName string, dataValue string, db *bolt.DB) error {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(dataName), []byte(dataValue))
		return err
	})
	return nil
}

func GetDataBucket(bucketName string, dataName string) (string, error) {
	db, err := openDatabase()
	defer db.Close()
	if err != nil {
		return "", err
	}
	var dataLen string
	dataLen, err = getDataLenBucket(bucketName + "_len", dataName, db)
	var dataLenInt int
	dataLenInt, err = strconv.Atoi(dataLen)
	var data string
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(dataName))
		if v != nil {
			data = string(v)
			return nil
		} else {
			return errors.New("Key not found")
		}
	})
	if err != nil {
		return "", err
	}
	dataClean := base.DecString(data, dataLenInt, cfg.Key)
	return dataClean, nil
}

func getDataLenBucket(bucketName string, dataName string, db *bolt.DB) (string, error) {
	var data string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(dataName))
		if v != nil {
			data = string(v)
		}
		return nil
	})
	return data, nil
}
