package main

import (
	"morfa/server/pkg/setting"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initRoute() {

	// Group clients related routes together
	router.Use(CORSMiddleware())

	// store := persistence.NewInMemoryStore(time.Second)
	clientsRoutes := router.Group("/setting")
	{
		clientsRoutes.GET("/get/:id", setting.GetById)
		clientsRoutes.POST("/create", setting.Create)
		clientsRoutes.POST("/update", setting.Update)

	}

}

func main() {

	// Normal options
	// flag.BoolVar(&opts.Single, "s", false, "")                 // single
	// flag.CommandLine.SetOutput(ioutil.Discard)
	// flag.Parse()

	router = gin.Default()
	initRoute()

	_ = router.Run("0.0.0.0:8080")

}
