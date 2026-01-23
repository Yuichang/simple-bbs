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

	// ハンドラにDBを渡す
	h:=&handlers.Handler{DB:db}
	// "indexページ"
	r.GET("/", h.ShowIndex)
	// 投稿

	// サーバ起動(http://localhost:8080)
	r.Run()
}
