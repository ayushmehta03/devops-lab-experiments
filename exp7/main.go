package main

import (
	"fmt"
	"log"
	"example.com/githubactions/controller"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()

	router.GET("/health",controller.HealthCheck)


	port:="8080"
	

	log.Printf("Server started on port %s",port)

	if err:=router.Run(":"+port);err!=nil{
		fmt.Println("Failed to start server:",err)
	}



}

