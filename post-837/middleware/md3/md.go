package md3

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Md3(ctx *gin.Context) {
	fmt.Println("md3 start")
	ctx.Next()
	fmt.Println("md3 end")
}

