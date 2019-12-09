package database
import (
	"micro-microblog/typings"
)
// GetUserIDBySessionID 根据给定的 sessionID 查找对应的用户 ID，若该用户不存在，则返回 errors.New("用户不存在")
func GetUserIDBySessionID(sessionID string) (int, error)

// DeleteSessionByID 根据给定的 sessionID 删除对应的 session 记录
func DeleteSessionByID(sessionID string) error

// UserLogin 根据给定的 username 和 password（明文）查询对应的用户是否存在
// 若用户名和密码均正确，则使用 uuid.NewV4() 创建一个新的 session 记录，并返回对应的 sessionID
// 若该用户不存在或密码不正确，则返回 errors.New("用户名或密码错误")
func UserLogin(username, password string) (sessionID string, err error)

//得到所有的用户
func GetAllUsers() [](typings.User)

//通过用户id得到user信息
func GetUserByUserID(userId int) (user typings.User ,err error)
//通过用户id修改用户信息返回报错信息
func modifyInfo(userId int,userInfo typings.User_tem) (error)