package gie

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc  func(*Context)


type RouterGroup struct{
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}

type Engine struct{
	*RouterGroup
	route *router
	group []*RouterGroup
}

func New() *Engine{
	engine:=&Engine{route:newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.group = []*RouterGroup{engine.RouterGroup}
	return engine
}


func (g *RouterGroup)Use(middlewares ...HandlerFunc){
	g.middlewares = append(g.middlewares,middlewares...)
}

func (g *RouterGroup)Group(prefix string) *RouterGroup{
	engine:=g.engine
	newGroup:=&RouterGroup{
		engine:engine,
		prefix:g.prefix + prefix,
		parent:g,
	}
	engine.group = append(engine.group, newGroup)
	return newGroup
}

func (g *RouterGroup)addroute(method string ,comp string, handler HandlerFunc){
	patter:=g.prefix + comp
	g.engine.route.addRoute(method,patter,handler)
}


func (g *RouterGroup)Get(patter string,handler HandlerFunc){
	g.addroute("GET",patter,handler)
}

func(g *RouterGroup)Post(patter string, handler HandlerFunc){
	g.addroute("POST", patter,handler)
}


func (engine *Engine)Run(addr string){
	fmt.Println(http.ListenAndServe(addr,engine))
}

func (engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
	
	var middlewares []HandlerFunc
	for _,g:=range engine.group{
		if strings.HasPrefix(req.URL.Path,g.prefix){
			middlewares = append(middlewares, g.middlewares...)
		}
	}
	
	c:=newContext(w,req)
	c.handlers = middlewares
	engine.route.handle(c)
}