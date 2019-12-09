package typings
// Register 的请求体
type Registerbody struct {
	Username string `json:"username"`
	Name string `json:"name"`
	StudentId string `json:"studentId"`
	Motto string `json:"motto"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}