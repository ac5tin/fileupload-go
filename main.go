package main

import (
	//"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"flag"
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"fileupload/file"
)

var (
	// string flag "addr" stored in pointer ip with type *String
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {

	flag.Parse() // parse command line into defined flags
	router := gin.Default()

	// api routes
	router.GET("/api/ping", func(c *gin.Context) { c.String(200, "pong") })
	/*
	   apiroute := router.Group("/api")
	   file.Routes(apiroute.Group("/file"))
	*/
	fileapi := router.Group("/api/file")
	file.Routes(fileapi)

	// serve static built js
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./client/build/index.html")
	})

	log.Println(router.Run(*addr))
}
