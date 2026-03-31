package main

import (
	"log"

	"github.com/Yuichang/simple-bbs/handlers"
	"github.com/Yuichang/simple-bbs/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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

	// ユーザー登録ページ
	r.GET("/register", h.ShowRegister)

	// ログインページ
	r.GET("/login", h.ShowLogin)

	// アカウント登録
	r.POST("/register", h.AccountRegister)

	// 投稿作成
	r.POST("/home", h.CreatePost)

	// 投稿削除
	r.POST("/delete/:id", h.DeletePost)

	// ログイン処理
	r.POST("/login", h.Login)

	// ログアウト処理
	r.GET("/logout", h.Logout)

	// サーバ起動(http://localhost:8080)
	r.Run()
}
