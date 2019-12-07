package middlewares

import "github.com/gin-gonic/gin"

// GetSessionID 取出请求 cookie 的 SESSIONID 字段
func GetSessionID() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("SESSIONID")
		if err != nil {
			sessionID = ""
		}
		c.Set("SessionID", sessionID)
	}
}
