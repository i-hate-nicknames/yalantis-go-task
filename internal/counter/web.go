package counter

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	numClients int
}

func StartApp() {
	app := &App{}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		app.numClients++
		log.Println("Serving response!")
		c.JSON(http.StatusOK, gin.H{
			"clients": app.numClients,
		})
	})
	r.Run(":8080")
}
