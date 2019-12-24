// geektutu.com
// main.go
package main

import (

	"time"
	"log"
//	"net/http"
	"./gie"
	"fmt"
	"net/http"
)


//curl "http://localhost:6666/hello?name=geektutu"
//curl "http://localhost:6666/login" -X POST -d 'username=geektutu&password=1234'
//curl "http://localhost:9999/hello/geektutu"
//curl "http://localhost:9999/assets/css/geektutu.css"
func main() {
	engine:=gie.New()	
	engine.Get("/",func(c *gie.Context){
		fmt.Fprintf(c.Write,"url = %s\n",c.Req.URL.Path)
	})
	engine.Get("/hello/name/123",func(c *gie.Context){
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)//Query -- curl "http://localhost:6666/hello?name=geektutu"
	})
	engine.Get("/hello/name/456",func(c *gie.Context){
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})	
	engine.Post("/login",func(c *gie.Context){
		c.Json(http.StatusOK, gie.H{
			"username": c.PostForm("username"),//PostForm -- curl "http://localhost:6666/login" -X POST -d 'username=geektutu&password=1234'
			"password": c.PostForm("password"),
		})
	})
	engine.Get("/hello/:name/123", func(c *gie.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello1 %s, you're at %s\n", c.Param("name"), c.Path)//Param -- curl "http://localhost:9999/hello/geektutu"
	})

	engine.Get("/hello/:hhh/123", func(c *gie.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello2 %s, you're at %s\n", c.Query("hhh"), c.Path)
	})


	engine.Get("/assets/*filepath", func(c *gie.Context) {
		c.Json(http.StatusOK, gie.H{"filepath": c.Param("filepath")}) // curl "http://localhost:9999/assets/css/geektutu.css"
	})


	g1:=engine.Group("chen")
	g1.Get("123",func(c *gie.Context){
		c.Html(http.StatusOK, "<h1>Index Page</h1>")//curl "http://localhost:9999/chen123"
	})

	g2:=engine.Group("tao")
	g2.Get("/123",func(c *gie.Context){
		c.Data(http.StatusOK, []byte("chentao"))//curl "http://localhost:9999/tao/123"
	})

	middware(engine)

	engine.Run(":9999")
}




func middware(r *gie.Engine){
	r.Use(Logger()) // global midlleware
	r.Get("/mid", func(c *gie.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2(),onlyForV22()) // v2 group middleware
	{
		v2.Get("/mid/:name", func(c *gie.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
			log.Println("v2")
		})
	}
}


func Logger() gie.HandlerFunc {
	return func(c *gie.Context) {
		// Start timer
		t := time.Now()
		log.Println(t)
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func onlyForV2() gie.HandlerFunc {
	return func(c *gie.Context) {
		c.Next()
		t := time.Now()
		log.Printf("onlyForV2 after in %v ", time.Since(t))
	}
}


func onlyForV22() gie.HandlerFunc {
	return func(c *gie.Context) {
		log.Printf("onlyForV22 before\n", )
		c.Next()
		log.Printf("onlyForV22 after\n", )
	}
}