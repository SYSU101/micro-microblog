package database

import (
	"fmt"
	"log"
	"strconv"
	"github.com/boltdb/bolt"
	"errors"
	"crypto/sha256"
	"github.com/satori/go.uuid"
	"reflect"
	"micro-microblog/typings"
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
		bucket_key := [9][]byte{
			[]byte("SystemValue"),
			[]byte("SessionID"), []byte("Id"),
			[]byte("Username"), []byte("Name"),
			[]byte("StudentId"), []byte("Motto"),
			[]byte("Password"), []byte("Birthday")}
		for i := 0; i < 9; i++ {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket_key[i]))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		b_SystemValue := tx.Bucket([]byte("SystemValue"))
		if b_SystemValue.Get([]byte("id_counter")) == nil {
			b_SystemValue.Put([]byte("id_counter"), []byte("0"))
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
	var return_id int = 0
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("SessionID"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(k) == sessionID {
				return_id, _ = strconv.Atoi(string(v))
				return nil
			}
		}
		return nil
	})
	if return_id == 0{
		return return_id, errors.New("获取失败，用户不存在")
	} else {
		return return_id, nil
	}
}

// DeleteSessionByID 根据给定的 sessionID 删除对应的 session 记录
func DeleteSessionByID(sessionID string) error{
	var err error = nil
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("SessionID"))
		if b.Delete([]byte(sessionID)) != nil{
			err = errors.New("删除失败，用户不存在")
		}
		return nil
	})
	return err
}

// UserLogin 根据给定的 username 和 Password（明文）查询对应的用户是否存在
// 若用户名和密码均正确，则使用 uuid.NewV4() 创建一个新的 session 记录，并返回对应的 sessionID
// 若该用户不存在或密码不正确，则返回 errors.New("用户名或密码错误")
func UserLogin(username, Password string) (string, error){
	var err error = nil
	var new_uuid []byte
	db.Update(func(tx *bolt.Tx) error {
		b_Username := tx.Bucket([]byte("Username"))
		b_Password := tx.Bucket([]byte("Password"))
		b_SessionID := tx.Bucket([]byte("SessionID"))
		id := b_Username.Get([]byte(username))
		encode_password := sha256.Sum256([]byte(Password))
		if (id == nil) || !reflect.DeepEqual(b_Password.Get(id), encode_password[:]) {
			err = errors.New("用户名或密码错误")
		} else {
			new_uuid := uuid.NewV4()
			b_SessionID.Put(id, new_uuid.Bytes())
		}
		return nil
	})
	return string(new_uuid), err
}

// CreateUserIdByRegister 根据给定的RegisterBody创建新的一个UserID
// 若该用户已存在，则返回 errors.New("用户冲突")
func CreateUserIdByRegister(userInfo *typings.Registerbody)(string, error) {
	var err error = nil
	var id []byte
	db.Update(func(tx *bolt.Tx) error {
		b_Id := tx.Bucket([]byte("Id"))
		if b_Id.Get([]byte(userInfo.Username)) != nil {
			err = errors.New("用户冲突")
		} else {
			b_SystemValue := tx.Bucket([]byte("SystemValue"))
			id = b_SystemValue.Get([]byte("id_counter"))
			id_int, _ := strconv.Atoi(string(id)) 
			b_SystemValue.Put([]byte("id_counter"), []byte(string(id_int + 1)))
			b_temp := tx.Bucket([]byte("Id"))
			b_temp.Put([]byte(userInfo.Username), id)
			b_temp = tx.Bucket([]byte("Username"))
			b_temp.Put(id, []byte(userInfo.Username))
			b_temp = tx.Bucket([]byte("Name"))
			b_temp.Put(id, []byte(userInfo.Name))
			b_temp = tx.Bucket([]byte("StudentId"))
			b_temp.Put(id, []byte(userInfo.StudentId))
			b_temp = tx.Bucket([]byte("Motto"))
			b_temp.Put(id, []byte(userInfo.Motto))
			b_temp = tx.Bucket([]byte("Password"))
			b_temp.Put(id, []byte(userInfo.Password))
			b_temp = tx.Bucket([]byte("Birthday"))
			b_temp.Put(id, []byte(userInfo.Birthday))
		}
		return nil
	})
	return string(id), nil
}


//得到所有的用户
func GetAllUsers() [](typings.User){
	var return_user_list []typings.User
	var index int = 0
	db.View(func(tx *bolt.Tx) error {
		b_Id := tx.Bucket([]byte("Id"))
		c_id := b_Id.Cursor()
		for _, v := c_id.First(); v != nil; _, v = c_id.Next() {
			return_user_list[index].Id = string(v)
			b_temp := tx.Bucket([]byte("Username"))
			return_user_list[index].Username = string(b_temp.Get(v))
			b_temp = tx.Bucket([]byte("Name"))
			return_user_list[index].Name = string(b_temp.Get(v))
			b_temp = tx.Bucket([]byte("StudentId"))
			return_user_list[index].StudentId = string(b_temp.Get(v))
			b_temp = tx.Bucket([]byte("Motto"))
			return_user_list[index].Motto = string(b_temp.Get(v))
			b_temp = tx.Bucket([]byte("Password"))
			return_user_list[index].Password = string(b_temp.Get(v))
			b_temp = tx.Bucket([]byte("Birthday"))
			return_user_list[index].Birthday = string(b_temp.Get(v))
		}
		return nil
	})
	return return_user_list
}

//通过用户id得到user信息
func GetUserByUserID(userId int) (user typings.User ,err error){
	db.View(func(tx *bolt.Tx) error {
		Id := []byte(string(userId))
		b_temp := tx.Bucket([]byte("Username"))
		user.Username = string(b_temp.Get(Id))
		b_temp = tx.Bucket([]byte("Name"))
		user.Name = string(b_temp.Get(Id))
		b_temp = tx.Bucket([]byte("StudentId"))
		user.StudentId = string(b_temp.Get(Id))
		b_temp = tx.Bucket([]byte("Motto"))
		user.Motto = string(b_temp.Get(Id))
		b_temp = tx.Bucket([]byte("Password"))
		user.Password = string(b_temp.Get(Id))
		b_temp = tx.Bucket([]byte("Birthday"))
		user.Birthday = string(b_temp.Get(Id))
		return nil
	})
	return user, err
}

//通过用户id修改用户信息返回报错信息
func modifyInfo(userId int, userInfo typings.User_tem) (error) {
	db.Update(func(tx *bolt.Tx) error {
		Id := []byte(string(userId))
		b_temp := tx.Bucket([]byte("Name"))
		b_temp.Put(Id, []byte(userInfo.Name))
		b_temp = tx.Bucket([]byte("StudentId"))
		b_temp.Put(Id, []byte(string(userInfo.StudentId)))
		b_temp = tx.Bucket([]byte("Motto"))
		b_temp.Put(Id, []byte(userInfo.Motto))
		b_temp = tx.Bucket([]byte("Birthday"))
		b_temp.Put(Id, []byte(userInfo.Birthday))
		return nil
	})
	return nil
}