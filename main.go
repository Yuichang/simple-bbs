package main

import (
	"github.com/Yuichang/simple-bbs/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// テンプレート読み込み
	r.LoadHTMLGlob("templates/*")

	// "indexページ"
	r.GET("/", handlers.ShowIndex)

	// サーバ起動(http://localhost:8080)
	r.Run()
}
