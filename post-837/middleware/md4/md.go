package md4

import (
"fmt"
"github.com/gin-gonic/gin"
)

func Md4(ctx *gin.Context) {
	fmt.Println("md4 start")
	ctx.Next()
	fmt.Println("md4 end")
}

