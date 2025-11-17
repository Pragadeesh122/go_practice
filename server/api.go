package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)



func main(){

	router := gin.Default()


	router.GET("/", func(res *gin.Context){
		res.JSON(http.StatusOK, gin.H{
			"Message" : "The server is up and running",
		})
	})

	err := router.Run("localhost:3000")
	if err != nil{
		fmt.Println("There was an error starting the server")
	}
}
