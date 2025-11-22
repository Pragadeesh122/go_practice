package api


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)



func Main(){

	router := gin.Default()


	router.GET("/", func(res *gin.Context){
		res.JSON(http.StatusOK, gin.H{
			"Message" : "The server is up and running",
		})
	})


	router.GET("/gomodule", func(res *gin.Context){
		res.JSON(http.StatusOK, gin.H{
			"GoModule" : "This is the Go Module on the gomodule path",
		})
	})

	err := router.Run("localhost:3000")
	if err != nil{
		fmt.Println("There was an error starting the server")
	}
}
