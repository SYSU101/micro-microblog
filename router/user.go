package router

import (
	"net/http"

	"micro-microblog/database"
	"micro-microblog/typings"

	"github.com/gin-gonic/gin"
)
func routeUser() {
	router.GET("/api/user/:id", getUserById)
	router.GET("/api/users", getUsers)
}
func getUsers(c *gin.Context){
	user := &typings.User{
		id: 0,
		username: "",
		name: "",
		studentId: 0, 
		motto: "",
		password: "",
		brithday: "",
	}
	
	users:=database.GetAllUsers()
	c.JSON(200, gin.H{
		"user": users
	})
	

}
func getUserById(c *gin.Context){
	user := &typings.User{
		id: 0,
		username: "",
		name: 0,
		studentId: "", 
		motto: "",
		password: "",
		brithday: "",
	}
	
	userId:=c.Param("id")
	if user, err := database.GetUserByUserID(userId); err != nil {
		c.JSON(404, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"user": user
		})
	}

}
