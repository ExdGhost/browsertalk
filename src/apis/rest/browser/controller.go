package browser

import (
	"net/http"
	browser "projects/browsertalk/src/modules/browser"

	"github.com/gin-gonic/gin"
)

// Launch ...
func Launch(c *gin.Context) {

	var query *browser.Query

	if err := c.BindQuery(&query); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := browser.Start(query); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, "Success")
}
