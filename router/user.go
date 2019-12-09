package router

import (
	"net/http"

	"micro-microblog/database"
	"micro-microblog/typings"

	"github.com/gin-gonic/gin"
)

func routeUser() {
	router.POST("/api/user", post)
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
func post(c *gin.Context) {
	userInfo := &typings.User{
		id: "",
		username: "",
		name: "",
		studentId: "",
		motto: "",
		password: "",
		birthday: "",
	}
	if err := c.ShouldBindJSON(userInfo); err != nil {
		//400
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": "错误请求",
		})
		return
	}
	if ID, err := database.GetUserIdByUserName(userInfo); err != nil {
		//401未授权
		c.JSON(http.StatusUnauthorized, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		//202
		c.JSON(http.StatusAccepted, gin.H{
			"userid": userID,
		})
	}
}

/*GET /api/user/:id
    response:
        200:
            user: user
        404:
            errMsg: string*/
