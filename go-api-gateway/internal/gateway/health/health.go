package health

import "github.com/gin-gonic/gin"

func RegisterHealthHandler(g *gin.Engine) {
	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
