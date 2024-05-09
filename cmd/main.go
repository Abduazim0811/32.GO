package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type User struct{
	Name string `json:"name"`
}

var users = []User{{Name: "John"}, {Name: "Doe"}}

func main() {
	router := gin.Default()

	router.GET("/users", GetUsers)
	router.POST("/users", AddUsers)
	router.PUT("/users/:name", UpdateUsers)
	router.Run(":8080")
}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,users)
}

func UserName(ctx *gin.Context){
	name:=ctx.Param("name")
	ctx.JSON(http.StatusOK,"Hello " + name)
}

func AddUsers(ctx *gin.Context){
	var user User
	err:=ctx.BindJSON(&user)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	users=append(users, user)
	ctx.JSON(http.StatusOK,"Yaxshi qoshildi")

}

func UpdateUsers(ctx *gin.Context){
	var user User
	err:=ctx.BindJSON(&user)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	name:=ctx.Param("name")

	for index,v:=range users{
		if v.Name==name{
			users[index].Name=user.Name
		}
	}
	
	ctx.JSON(http.StatusOK,"O'zgartirildi")
}


