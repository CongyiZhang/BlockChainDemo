package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	fmt.Println("hello world")

	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket失败")
			}
		}
		bucket.Put([]byte("1111"), []byte("hello"))
		bucket.Put([]byte("2222"), []byte("hello"))
		return nil
	})
}
