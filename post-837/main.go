package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinfeijie/blog-demo/post-837/middleware/md1"
	"github.com/jinfeijie/blog-demo/post-837/middleware/md2"
	"github.com/jinfeijie/blog-demo/post-837/middleware/md3"
	"github.com/jinfeijie/blog-demo/post-837/middleware/md4"
	"math"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), md1.Md1, md2.Md2, md3.Md3, md4.Md4)

	fmt.Println(math.MaxInt8 / 2)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, []int{1, 2, 3})
	})

	if err := http.ListenAndServe(":1080", router); err != nil {
		panic(err.Error())
	}
}
