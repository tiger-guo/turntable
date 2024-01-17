package main

import "github.com/gin-gonic/gin"

func Turntable(ctx *gin.Context) {
	ctx.HTML(200, "turntable.html", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./template/*")
	router.GET("/", Turntable)
	router.Run(":11314")
}
