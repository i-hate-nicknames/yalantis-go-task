package counter

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const visitedCookieName = "visited"

type App struct {
	numClients int
	mux        sync.Mutex
}

func StartApp() {
	app := &App{}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		if !hasVisited(c) {
			setVisited(c)
			app.mux.Lock()
			app.numClients++
			app.mux.Unlock()
		}
		log.Println("Serving response!")
		c.JSON(http.StatusOK, gin.H{
			"clients": app.numClients,
		})
	})
	r.Run(":8080")
}

// true if user has visited cookie
func setVisited(c *gin.Context) {
	// todo: secure cookie
	cookie := &http.Cookie{Name: visitedCookieName, Value: "1"}
	http.SetCookie(c.Writer, cookie)
}

// true if user has visited cookie
func hasVisited(c *gin.Context) bool {
	// todo: validate cookie
	_, err := c.Request.Cookie(visitedCookieName)
	if err != nil {
		return false
	}
	return true
}
