package counter

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Println("Serving response!")
		c.JSON(http.StatusNoContent, gin.H{})
	})
	r.Run(":8080")
}
