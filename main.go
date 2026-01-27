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

	// 静的ファイル
	r.Static("/static", "./static")

	// テンプレート読み込み
	r.LoadHTMLGlob("templates/*")

	// ハンドラにDBを渡す
	h := &handlers.Handler{DB: db}
	// "indexページ"
	r.GET("/", h.ShowIndex)

	// "homeページ"
	r.GET("/home", h.ShowHome)

	// 投稿
	r.POST("/home", h.CreatePost)
	// サーバ起動(http://localhost:8080)
	r.Run()
}
