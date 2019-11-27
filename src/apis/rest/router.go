package rest

import (
	"log"
	"net/http"
	"projects/browsertalk/src/apis/rest/browser"
	cfg "projects/browsertalk/src/config"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Build ...
func Build(config *cfg.Model) {
	port := config.Server.HTTP.Port
	router := gin.Default()

	cnf := cors.DefaultConfig()
	cnf.AllowAllOrigins = true
	cnf.AllowCredentials = true
	cnf.AddAllowHeaders("cache-control")
	cnf.AddAllowMethods("GET", "PUT", "POST", "DELETE")
	router.Use(cors.New(cnf))

	// Inits routes
	browser.Routes(router)

	// Debug statement
	router.GET("/ping", checkPing)

	// Starts router
	router.Run(port)

}

func checkPing(c *gin.Context) {
	log.Println("INSIDE CHECKPING")
	c.String(http.StatusOK, "PONG")
}
