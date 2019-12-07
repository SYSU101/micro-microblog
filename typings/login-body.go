package typings

// LoginBody 登录的请求体
type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
