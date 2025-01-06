package main

import (
	"fmt"
	"lg"
)

func main() {
	r := lg.New()
	r.GET("/hello", func(c *lg.Context) {
		c.HTML(200, "<h1>hello</h1>")
	})
	r.GET("/hi", func(c *lg.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.String(200, "hi name %s and age %s", name, age)
	})

	r.POST("/login", func(c *lg.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		c.JSON(200, lg.H{
			"username": username,
			"password": password,
		})
	})
	r.Run(":7777")
}
