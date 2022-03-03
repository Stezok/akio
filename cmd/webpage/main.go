package main

import "NFTProject/internal/handler/webpage"

func main() {
	handler := webpage.Handler{
		HTMLGlobPath: "web/html/*",
		CSSPath:      "web/css/",
		AssetsPath:   "assets/",
	}

	handler.InitRoutes().Run()
}
