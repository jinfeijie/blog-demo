package md2

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Md2(ctx *gin.Context) {
	fmt.Println("md2 start")
	if ctx.Query("login") == "0" {
		ctx.Abort()
	}
	ctx.Next()
	fmt.Println("md2 end")
}
