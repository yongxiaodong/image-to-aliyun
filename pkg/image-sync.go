package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImageMetadata struct {
	SourceImage string `form:"source"`
}

func ImageSync(context *gin.Context) {
	var imageMetadata ImageMetadata
	err := context.ShouldBind(&imageMetadata)
	if err != nil {
		fmt.Println(err)
		//context.String(200, err.Error())
	}
	fmt.Println(imageMetadata.SourceImage)
	context.String(200, "ing....")
}
