package router

import (
	"net/http"

	"micro-microblog/database"
	"micro-microblog/typings"

	"github.com/gin-gonic/gin"
)

func routeSession() {
	router.GET("/api/seesion", getUserID)
	router.PUT("/api/session", login)
	router.DELETE("/api/session", logout)
}

func getUserID(c *gin.Context) {
	sessionID := c.MustGet("SessionID").(string)
	if userID, err := database.GetUserIDBySessionID(sessionID); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"userId": userID,
		})
	}
}

func login(c *gin.Context) {
	loginBody := &typings.LoginBody{
		Username: "",
		Password: "",
	}
	if err := c.ShouldBindJSON(loginBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": "无法识别的登录请求",
		})
		return
	}
	if sessionID, err := database.UserLogin(loginBody.Username, loginBody.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		c.SetCookie("SESSIONID", sessionID, -1, "/", "", false, false)
		c.JSON(http.StatusAccepted, gin.H{})
	}
}

func logout(c *gin.Context) {
	sessionID := c.MustGet("SessionID").(string)
	database.DeleteSessionByID(sessionID)
	c.SetCookie("SESSIONID", "", -1, "/", "", false, false)
	c.JSON(http.StatusNoContent, gin.H{})
}
