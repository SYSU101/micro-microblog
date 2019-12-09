package typings
// Post User 的请求体
type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	StudentId string `json:"studentId"`
	Motto string `json:"motto"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}