package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/boltdb/bolt"
	"errors"
)

func test_read(){
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("name"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key = %s, value = %s\n", k, v)
		}
		return nil
	})
}

func test_GetUserIDBySessionID(sessionID string){
	fmt.Printf("test_GetUserIDBySessionID\n")
	fmt.Printf("sessionID = %s\n", sessionID)
	x, err := GetUserIDBySessionID(sessionID)
	if err == nil {
		fmt.Printf("userID = %d\n", x)
	} else {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("====================\n")
}

func test_DeleteSessionByID(sessionID string){
	fmt.Printf("test_DeleteSessionByID\n")
	fmt.Printf("sessionID = %s\n", sessionID)
	err := DeleteSessionByID(sessionID)
	if err == nil {
		fmt.Printf("删除成功\n")
	} else {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("====================\n")
}

func GetUserIDBySessionID(sessionID string) (int, error){
	var return_int int = 0
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sessionID"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(k) == sessionID {
				return_int, _ = strconv.Atoi(string(v))
				return nil
			}
		}
		return nil
	})
	if return_int == 0{
		return return_int, errors.New("获取失败，用户不存在")
	} else {
		return return_int, nil
	}
}

func DeleteSessionByID(sessionID string) error{
	var err error = nil
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sessionID"))
		if b.Delete([]byte(sessionID)) != nil{
			err = errors.New("删除失败，用户不存在")
		}
		return nil
	})
	return err
}

// func UserLogin(username, password string) (sessionID string, err error){

// }

func main(){
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sessionID"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("sessionID1"), []byte("17343092"))
		b.Put([]byte("sessionID2"), []byte("17343093"))
		b.Put([]byte("sessionID3"), []byte("17343098"))
		b.Put([]byte("sessionID4"), []byte("17343099"))
		b.Put([]byte("sessionID5"), []byte("17343101"))
		b.Put([]byte("sessionID6"), []byte("17343102"))	

		b, err = tx.CreateBucketIfNotExists([]byte("id"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("17343092"))
		b.Put([]byte("17343093"), []byte("17343093"))
		b.Put([]byte("17343098"), []byte("17343098"))
		b.Put([]byte("17343099"), []byte("17343099"))
		b.Put([]byte("17343101"), []byte("17343101"))
		b.Put([]byte("17343102"), []byte("17343102"))

		b, err = tx.CreateBucketIfNotExists([]byte("username"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("PPC"))
		b.Put([]byte("17343093"), []byte("PHC"))
		b.Put([]byte("17343098"), []byte("QA"))
		b.Put([]byte("17343099"), []byte("SFZ"))
		b.Put([]byte("17343101"), []byte("SQD"))
		b.Put([]byte("17343102"), []byte("SYH"))

		b, err = tx.CreateBucketIfNotExists([]byte("name"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("潘鹏程"))
		b.Put([]byte("17343093"), []byte("庞海成"))
		b.Put([]byte("17343098"), []byte("全奥"))
		b.Put([]byte("17343099"), []byte("沈方哲"))
		b.Put([]byte("17343101"), []byte("苏祺达"))
		b.Put([]byte("17343102"), []byte("苏禹行"))

		b, err = tx.CreateBucketIfNotExists([]byte("studentId"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("17343092"))
		b.Put([]byte("17343093"), []byte("17343093"))
		b.Put([]byte("17343098"), []byte("17343098"))
		b.Put([]byte("17343099"), []byte("17343099"))
		b.Put([]byte("17343101"), []byte("17343101"))
		b.Put([]byte("17343102"), []byte("17343102"))

		b, err = tx.CreateBucketIfNotExists([]byte("motto"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("------"))
		b.Put([]byte("17343093"), []byte("------"))
		b.Put([]byte("17343098"), []byte("------"))
		b.Put([]byte("17343099"), []byte("------"))
		b.Put([]byte("17343101"), []byte("------"))
		b.Put([]byte("17343102"), []byte("------"))

		b, err = tx.CreateBucketIfNotExists([]byte("password"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("17343092"))
		b.Put([]byte("17343093"), []byte("17343093"))
		b.Put([]byte("17343098"), []byte("17343098"))
		b.Put([]byte("17343099"), []byte("17343099"))
		b.Put([]byte("17343101"), []byte("17343101"))
		b.Put([]byte("17343102"), []byte("17343102"))

		b, err = tx.CreateBucketIfNotExists([]byte("brithday"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("17343092"), []byte("17343092"))
		b.Put([]byte("17343093"), []byte("17343093"))
		b.Put([]byte("17343098"), []byte("17343098"))
		b.Put([]byte("17343099"), []byte("17343099"))
		b.Put([]byte("17343101"), []byte("17343101"))
		b.Put([]byte("17343102"), []byte("17343102"))

		return nil
	})
	db.Close()

	// test_read()
	test_GetUserIDBySessionID("sessionID0")
	test_GetUserIDBySessionID("sessionID1")

	test_DeleteSessionByID("sessionID0")
	test_DeleteSessionByID("sessionID1")
	
	test_GetUserIDBySessionID("sessionID1")
}