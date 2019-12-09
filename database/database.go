package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/boltdb/bolt"
	"errors"
	"crypto/sha256"
	"github.com/satori/go.uuid"
	"reflect"
)

var db *bolt.DB

func init() {
	StartDB("test.db")
}

// 通过数据库文件名加载数据库
func StartDB(db_file_name string) {
	CloseDB()
	var err error
	db, err = bolt.Open(db_file_name, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	//所有的bucket的键为id，id的键为username
	db.Update(func(tx *bolt.Tx) error {
		bucket_key := [8][]byte{
			[]byte("sessionID"), []byte("id"),
			[]byte("username"), []byte("name"),
			[]byte("studentId"), []byte("motto"),
			[]byte("password"), []byte("brithday")}
		for i := 0; i < 8; i++ {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket_key[i]))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		return nil
	})
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// GetUserIDBySessionID 根据给定的 sessionID 查找对应的用户 ID，若该用户不存在，则返回 errors.New("用户不存在")
func GetUserIDBySessionID(sessionID string) (int, error){
	var return_int int = 0
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

// DeleteSessionByID 根据给定的 sessionID 删除对应的 session 记录
func DeleteSessionByID(sessionID string) error{
	var err error = nil
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sessionID"))
		if b.Delete([]byte(sessionID)) != nil{
			err = errors.New("删除失败，用户不存在")
		}
		return nil
	})
	return err
}

// UserLogin 根据给定的 username 和 password（明文）查询对应的用户是否存在
// 若用户名和密码均正确，则使用 uuid.NewV4() 创建一个新的 session 记录，并返回对应的 sessionID
// 若该用户不存在或密码不正确，则返回 errors.New("用户名或密码错误")
func UserLogin(username, password string) (string, error){
	var err error = nil
	var new_uuid []byte
	db.Update(func(tx *bolt.Tx) error {
		b_username := tx.Bucket([]byte("username"))
		b_password := tx.Bucket([]byte("password"))
		b_sessionID := tx.Bucket([]byte("sessionID"))
		id := b_username.Get([]byte(username))
		encode_password := sha256.Sum256([]byte(password))

		if (id == nil) || !reflect.DeepEqual(b_password.Get(id), encode_password[:]) {
			err = errors.New("用户名或密码错误")
			return nil
		}
		new_uuid, _ := uuid.NewV4()
		b_sessionID.Put(id, new_uuid.Bytes())
		return nil
	})
	return string(new_uuid), err
}

// 注册用户，用户存在则返回 errors.New("用户名已被使用")
func UserRegister(username, name, password string) error {

}

// 编辑用户属性, sessionId 无用户时返回errors.New("用户未登录")
func UserEditName(sessionID, name string) error {

}

// 编辑用户属性, sessionId 无用户时返回errors.New("用户未登录")
func UserEditStudentId(sessionID, studentId string) error {

}

// 编辑用户属性, sessionId 无用户时返回errors.New("用户未登录")
func UserEditBrithday(sessionID, brithday string) error {

}

// 编辑用户属性, sessionId 无用户时返回errors.New("用户未登录")
func UserEditMotto(sessionID, motto string) error {

}

// 编辑用户属性, sessionId 无用户时返回errors.New("用户未登录")
func UserEditPassword(sessionID, password string) error {

}

func UserEditInfo(sessionID, name, studentId, brithday, motto string) error {

}

func main() {
}