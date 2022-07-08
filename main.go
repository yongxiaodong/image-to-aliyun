package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image-to-aliyun/pkg"
)

func main() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	router := r.Group("/")
	router.GET("image-sync", pkg.ImageSync)

	if err := r.Run(); err != nil {
		fmt.Println("err")
		return
	}
}
