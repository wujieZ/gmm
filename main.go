package main

import (
	"gmm"
	"net/http"
)

func main() {
	r := gmm.New()

	r.GET("/", func(c *gmm.Context) {
		c.HTML(200, "<h1>Hello Gmm</h1>")
	})

	r.GET("/hello", func(c *gmm.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gmm.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.Run(":3000")
}
