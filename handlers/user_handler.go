package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Yuichang/simple-bbs/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

// welcomeページ 後で変えるかも。
func (h Handler) ShowIndex(c *gin.Context) {
	posts, err := models.ListPosts(c.Request.Context(), h.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error", err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": posts,
	})
}
