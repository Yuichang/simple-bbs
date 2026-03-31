package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Yuichang/simple-bbs/models"
	"github.com/Yuichang/simple-bbs/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

// welcomeページ
func (h Handler) ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// アカウント登録ページ
func (h Handler) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// ホーム画面
func (h Handler) ShowHome(c *gin.Context) {
	posts, err := models.ListPosts(c.Request.Context(), h.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	session := sessions.Default(c)
	user := session.Get("user_name")
	uid := session.Get("user_id")

	var username string
	var userID int

	if user != nil {
		username = user.(string)
	}
	if uid != nil {
		userID = uid.(int)
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"posts":    posts,
		"username": username,
		"user_id":  userID,
	})
}

// ログインページ
func (h Handler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// 投稿を作成するハンドラ
func (h Handler) CreatePost(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("user_id")
	body := c.PostForm("body")

	if body == "" {
		c.String(http.StatusBadRequest, "コメント内容を入力してください")
		return
	}

	var userID int
	var name string

	if uid != nil {
		userID = uid.(int) // ログイン中
	} else {
		userID = 0 // 未ログイン
	}
	name = c.PostForm("name")
	if name == "" {
		name = "名無しさん"
	}

	err := models.CreatePost(c.Request.Context(), h.DB, userID, name, body)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}

// 投稿を削除するハンドラ
func (h Handler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	session := sessions.Default(c)
	uid := session.Get("user_id")

	if uid == nil {
		c.String(http.StatusForbidden, "ログインしてください")
		return
	}

	_, err := models.GetPostByID(c.Request.Context(), h.DB, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	err = models.DeletePost(c.Request.Context(), h.DB, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}

// アカウントを登録するハンドラ
func (h Handler) AccountRegister(c *gin.Context) {

	name := c.PostForm("username")
	pass := c.PostForm("password")
	gender := c.PostForm("gender")

	hashedPass, err := utils.GeneratedHash(pass)
	if err != nil {
		c.String(http.StatusInternalServerError, "Hash error")
		return
	}

	err = models.CreateAccount(c.Request.Context(), h.DB, name, gender, hashedPass)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	// 登録後すぐにログインさせる
	userID, _, _, err := models.GetUserByName(c.Request.Context(), h.DB, name)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", userID)
	session.Set("user_name", name)
	session.Save()

	c.Redirect(http.StatusFound, "/home")
}

// ログイン
func (h Handler) Login(c *gin.Context) {
	name := c.PostForm("username")
	pass := c.PostForm("password")

	userID, username, hashedPass, err := models.GetUserByName(c.Request.Context(), h.DB, name)
	if err != nil {
		c.String(http.StatusUnauthorized, "ユーザーが存在しません")
		return
	}

	if !utils.VeryifyPassword(hashedPass, pass) {
		c.String(http.StatusUnauthorized, "パスワードが違います")
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", userID)
	session.Set("user_name", username)
	session.Save()

	c.Redirect(http.StatusFound, "/home")
}

// ログアウト
func (h Handler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
