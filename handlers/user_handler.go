package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Yuichang/simple-bbs/models"
	"github.com/Yuichang/simple-bbs/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

// welcomeページ 後で変えるかも。
func (h Handler) ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// アカウント登録ページ
func (h Handler) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// welcomeページ 後で変えるかも。
func (h Handler) ShowHome(c *gin.Context) {
	posts, err := models.ListPosts(c.Request.Context(), h.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error", err)
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"posts": posts,
	})
}

// 投稿を作成するハンドラ
func (h Handler) CreatePost(c *gin.Context) {
	name := c.PostForm("name")
	body := c.PostForm("body")
	if name == "" {
		name = "名無しさん"
	}

	err := models.CreatePost(c.Request.Context(), h.DB, name, body)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}

// 投稿を削除するハンドラ
func (h Handler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := models.DeletePost(c.Request.Context(), h.DB, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error", err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/home")
}

// アカウントを登録するハンドラ
func (h Handler) AccountRegister(c *gin.Context) {

	// JSでバリデーションチェック済
	name := c.PostForm("username")
	pass := c.PostForm("password")
	gender := c.PostForm("gender")

	// パスワードをハッシュ化させる
	hashedPass, err := utils.GeneratedHash(pass)
	if err != nil {
		c.String(http.StatusInternalServerError, "Hash error", err)
		return
	}
	err = models.CreateAccount(c.Request.Context(), h.DB, name, gender, hashedPass)
}
