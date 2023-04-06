package v1

import (
	"net/http"
	"race-proj/util/dirsearch"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Request_DIRSearch(ctx *gin.Context) {
	targetURL := ctx.Query("url")
	how := ctx.Query("flag")
	result := dirsearch.Search(targetURL, com.StrTo(how).MustInt())
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
	})
}
