package main

import (
	"gmm"
)

func main() {
	r := gmm.New()

	r.GET("/", func(c *gmm.Context) {
		c.HTML(200, "<h1>Hello Gmm</h1>")
	})

	r.Run(":3000")
}
