package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// welcomeページ 後で変えるかも。
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
