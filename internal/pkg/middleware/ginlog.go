package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func GinLog() gin.HandlerFunc {
	return sloggin.New(slog.Default())
}
