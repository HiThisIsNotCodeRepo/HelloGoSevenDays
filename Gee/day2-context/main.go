package main

import (
	"context"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		//given: /hello?name=geektutu
		//want: hello geektutu, you're at /hello
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		// given: in postman "Body"
		//		key: username
		//			value: geektutu
		//		key: password
		//			value: 1234
		// want:
		//{
		//	"password": "1234",
		//	"username": "geektutu"
		//}
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
