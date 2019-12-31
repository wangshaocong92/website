package main

import (
	"github.com/gin-gonic/gin"
	"src/src/control"
)

func main() {
	engine := gin.Default()
	control.InitRouter(engine)
	engine.Run("172.21.0.7:80")
}
