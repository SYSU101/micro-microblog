package router
type User_tem struct {
	Name string `json:"name"`
	StudentId int `json:"studentId"`
	Motto string `json:"motto"`
	Birthday string `json:"birthday"`
}
import (
	"net/http"

	"micro-microblog/database"
	"micro-microblog/typings"

	"github.com/gin-gonic/gin"
)
func routeUser() {
	router.GET("/api/user/:id", getUserById)
	router.GET("/api/users", getUsers)
	router.PUT("/api/user/:id", modifyInfo)
}
func getUsers(c *gin.Context){
	
	
	users:=database.GetAllUsers()
	c.JSON(200, gin.H{
		"user": users,
	})

}
func modifyInfo(c *gin.Context){
	userId:=c.Param("id")
	userInfo:=&typings.User_tem{
		Name: "",
		StudentId: 0,
		Motto :"",
		Birthday:"",
	}
	if err := c.ShouldBindJSON(userInfo); err != nil {
		c.JSON(404, gin.H{
			"errMsg": "找不到",
		})
	}
	if err :=database.modifyInfo(userId,userInfo){
		c.JSON(403, gin.H{
			"errMsg": "拒绝",
		})
	}
	else{
		c.JSON(204, gin.H{})
	}
}
func getUserById(c *gin.Context){
	
	
	userId:=c.Param("id")
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


