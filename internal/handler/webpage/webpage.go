package webpage

import "github.com/gin-gonic/gin"

type Handler struct {
	HTMLGlobPath string
	CSSPath      string
	AssetsPath   string
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob(h.HTMLGlobPath)
	router.Static("/resource", h.CSSPath)
	router.Static("/assets", h.AssetsPath)

	root := router.Group("/")
	{
		root.GET("", handlerRoot)
	}

	return router
}
