package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/open-plm/timestone/pkg/config"
)

func ForPublic(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ForInternal(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ForPrivate(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}