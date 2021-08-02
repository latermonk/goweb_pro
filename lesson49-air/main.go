package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//r.GET("/", handleroot)
	r.Run(":9997")

}
