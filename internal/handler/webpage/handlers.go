package webpage

import "github.com/gin-gonic/gin"

func handlerRoot(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
