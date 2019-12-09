package router

import (
	"net/http"
	"strconv"

	"micro-microblog/database"
	"micro-microblog/typings"

	"github.com/gin-gonic/gin"
)

func routeUser() {
	router.POST("/api/users", createUserId)
	router.GET("/api/user/:id", getUserById)
	router.GET("/api/users", getUsers)
	router.PUT("/api/user/:id", modifyInfo)
}

/*GET /api/user:
    response:
		200: user*/
/*POST /api/user:
    body:
        username: string
		name: string
		motto: string
		birthday: string
		studentId: string
    response:
        201:
            userId
        409:
            errMsg*/
func createUserId(c *gin.Context) {
	registerBody := &typings.Registerbody{
		Username:  "",
		Name:      "",
		StudentId: "",
		Motto:     "",
		Password:  "",
		Birthday:  "",
	}
	if err := c.ShouldBindJSON(registerBody); err != nil {
		//400
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": err.Error(),
		})
		return
	}
	if userID, err := database.CreateUserIdByRegister(registerBody); err != nil {
		//409请求资源和当前状态存在冲突
		c.JSON(http.StatusConflict, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		//201
		c.JSON(http.StatusCreated, gin.H{
			"userid": userID,
		})
	}
}

func getUsers(c *gin.Context) {

	users := database.GetAllUsers()
	c.JSON(200, gin.H{
		"user": users,
	})

}
func modifyInfo(c *gin.Context) {
	sessionID := c.MustGet("SessionID").(string)
	userId, _ := strconv.Atoi(c.Param("id"))
	userInfo := &typings.User_tem{
		Name:      "",
		StudentId: "",
		Motto:     "",
		Birthday:  "",
	}
	if err := c.ShouldBindJSON(userInfo); err != nil {
		c.JSON(400, gin.H{
			"errMsg": err.Error(),
		})
	}
	authId, _ := database.GetUserIDBySessionID(sessionID)
	if userId != authId {
		c.JSON(403, gin.H{
			"errMsg": "拒绝",
		})
	} else {
		if err := database.ModifyInfo(userId, userInfo); err != nil {
			c.JSON(500, gin.H{
				"errMsg": err.Error(),
			})
		}
		c.JSON(204, gin.H{})
	}
}

func getUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if user, err := database.GetUserByUserID(userId); err != nil {
		c.JSON(404, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"user": user,
		})
	}

}
