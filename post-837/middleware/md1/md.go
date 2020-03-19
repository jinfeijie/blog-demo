package md1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Md1(ctx *gin.Context) {
	fmt.Println("md1 start")
	if ctx.Query("danger") == "1" {
		ctx.Abort()
	}
	ctx.Next()
	fmt.Println("md1 end")
}
