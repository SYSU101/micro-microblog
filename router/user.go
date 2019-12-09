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
	registerBody := &typings.Registerbody{
		username: "",
		name: "",
		studentId: "",
		motto: "",
		password: "",
		birthday: "",
	}
	if err := c.ShouldBindJSON(registerBody); err != nil {
		//400
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": "错误请求",
		})
		return
	}
	if userID, err := database.CreateUserIdByRegister(registerBody); err != nil {
		//401未授权
		c.JSON(http.StatusUnauthorized, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		//201
		c.JSON(http.StatusCreated, gin.H{
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
