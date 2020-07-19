package main

import (
	"gmm"
	"net/http"
)

func main() {
	r := gmm.New()

	r.GET("/hello/:name", func(c *gmm.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gmm.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gmm</h1>")
		})
		v1.GET("/hello", func(c *gmm.Context) {
			c.JSON(http.StatusOK, gmm.H{
				"username": c.Query("username"),
				"password": c.Query("password"),
			})
		})
	}
	r.Run(":3000")
}
