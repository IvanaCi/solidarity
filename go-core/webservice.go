package main

import (
	"github.com/gin-gonic/gin"
	"goapi/mappings"
	"github.com/go-sql-driver/mysql"
)



func main() {
    
	webservice := gin.Default()

	func setupRouter() *gin.Engine {
		webservice := gin.Default()
		webservice.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H {
				"message" : "Hello World!",
			})
		})
		return webservice
	}

	webservice.GET("/test", func(ctx *gin.Context){
		
		webservice:= setupRouter()
		webservice.Run("80:80")
	})
}