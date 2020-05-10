package counter

import (
	"net/http"
	"strconv"
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
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/", func(c *gin.Context) {
		if !hasVisited(c) {
			setVisited(c)
			app.mux.Lock()
			app.numClients++
			app.mux.Unlock()
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"totalClients": strconv.Itoa(app.numClients),
		})
	})
	r.Run(":8080")
}

// set is visited cookie
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
