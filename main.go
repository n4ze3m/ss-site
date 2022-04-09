package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n4ze3m/ss-site/routes"
)

func main() {
	r := gin.Default()
	routes.Screenshot(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
