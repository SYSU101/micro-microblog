package database

import (
	"fmt"
	"log"
	"strconv"
	"github.com/boltdb/bolt"
	"errors"
)

var db *DB

func init() {
	StartDB("test.db")
}

// 通过数据库文件名加载数据库
func StartDB(db_file_name string) error {
	var err
	db, err = bolt.Open(db_file_name, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// GetUserIDBySessionID 根据给定的 sessionID 查找对应的用户 ID，若该用户不存在，则返回 errors.New("用户不存在")
func GetUserIDBySessionID(sessionID string) (int, error) {

}

// DeleteSessionByID 根据给定的 sessionID 删除对应的 session 记录
func DeleteSessionByID(sessionID string) error

// UserLogin 根据给定的 username 和 password（明文）查询对应的用户是否存在
// 若用户名和密码均正确，则使用 uuid.NewV4() 创建一个新的 session 记录，并返回对应的 sessionID
// 若该用户不存在或密码不正确，则返回 errors.New("用户名或密码错误")
func UserLogin(username, password string) (sessionID string, err error)
