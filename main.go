package main

import (
	"log"

	"github.com/Yuichang/simple-bbs/handlers"
	"github.com/Yuichang/simple-bbs/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB接続
	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal("DB connect error:", err)
	}

	defer db.Close()

	// ルータ作成
	r := gin.Default()

	// テンプレート読み込み
	r.LoadHTMLGlob("templates/*")

	// "indexページ"
	r.GET("/", handlers.ShowIndex)

	// サーバ起動(http://localhost:8080)
	r.Run()
}
