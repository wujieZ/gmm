package main

import (
	"gmm"
	"log"
	"net/http"
	"time"
)

func onlyForV1() gmm.HandlerFunc {
	return func(c *gmm.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gmm.New()
	r.Use(gmm.Logger())

	r.GET("/hello/:name", func(c *gmm.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v1 := r.Group("/v1")
	v1.Use(onlyForV1())
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
