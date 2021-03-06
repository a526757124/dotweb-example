package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/cache"
	"github.com/devfeel/dotweb/framework/file"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置gzip开关
	//app.SetEnabledGzip(true)

	//设置路由
	InitRoute(app.HttpServer)

	//启动 监控服务
	//pprofport := 8081
	//go app.StartPProfServer(pprofport)

	//app.SetCache(cache.NewRuntimeCache())
	app.SetCache(cache.NewRedisCache("192.168.8.175:6379"))

	app.Cache().Set("g", "gv", 60)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

type UserInfo struct {
	UserName string
	Sex      int
}

func One(ctx *dotweb.HttpContext) {
	g := ctx.Cache().GetString("g")
	ctx.Cache().Incr("count")
	ctx.WriteString("One [" + g + "] ")
}

func Two(ctx *dotweb.HttpContext) {
	g := ctx.Cache().GetString("g")
	ctx.Cache().Incr("count")
	ctx.WriteString("Two [" + g + "] [" + ctx.Cache().GetString("count") + "]")
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/1", One)
	server.Router().GET("/2", Two)
}
