package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware(defaultPage int, defaultLimit int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = defaultPage
		}

		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			limit = defaultLimit
		}

		ctx.Set("page", page)
		ctx.Set("limit", limit)

		ctx.Next()
	}
}
