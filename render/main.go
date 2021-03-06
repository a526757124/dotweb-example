package main

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/devfeel/dotweb"
	"io"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置gzip开关
	//app.HttpServer.SetEnabledGzip(true)

	//设置路由
	InitRoute(app.HttpServer)

	//use Jet template
	//app.HttpServer.SetRenderer(NewJetRenderer().Reload(true))

	//启动 监控服务
	//pprofport := 8081
	//go app.StartPProfServer(pprofport)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

type UserInfo struct {
	UserName string
	Sex      bool
	Url      string
}

type BookInfo struct {
	Name string
	Size int64
}

func TestView(ctx *dotweb.HttpContext) {
	ctx.ViewData().Set("data", "图书信息")
	ctx.ViewData().Set("user", &UserInfo{UserName: "user1", Sex: true, Url: "client_id=dsfadafsdsfads1213fdsafdasfa&redirect_uri=http%3A%2F%2F192.168.1.100%3A9094%2Foauth2&response_type=code&scope=all&state=xyz"})
	m := make([]*BookInfo, 5)
	m[0] = &BookInfo{Name: "book0", Size: 1}
	m[1] = &BookInfo{Name: "book1", Size: 10}
	m[2] = &BookInfo{Name: "book2", Size: 100}
	m[3] = &BookInfo{Name: "book3", Size: 1000}
	m[4] = &BookInfo{Name: "book4", Size: 10000}
	ctx.ViewData().Set("Books", m)
	//ctx.ViewData().Set("urlPrams", "client_id=dsfadafsdsfads1213fdsafdasfa&redirect_uri=http%3A%2F%2F192.168.1.100%3A9094%2Foauth2&response_type=code&scope=all&state=xyz")
	//ctx.ViewData().Set("urlPrams", "i=1&x=2&c=3")
	//if use jet template, file name is testview_jet.html
	//if use go template, file name is testview.html
	ctx.View("d:/gotmp/templates/testview.html")

}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", TestView)
}

var views = jet.NewHTMLSet("d:/gotmp/templates")

type jetRenderer struct {
}

func (r *jetRenderer) Render(w io.Writer, tpl string, data interface{}, ctx *dotweb.HttpContext) error {
	view, err := views.GetTemplate(tpl)
	fmt.Println(view, err)
	if err != nil {
		fmt.Println("Unexpected template err:", err.Error())
	}
	//if use vars mode, template not use "."
	vars := convertMapToVar(data)
	return view.Execute(w, vars, nil)

	//if use data mode, template use "."
	//return view.Execute(w, nil, data)
}

func NewJetRenderer() *jetRenderer {
	r := new(jetRenderer)

	return r
}
func (jet *jetRenderer) Reload(b bool) *jetRenderer {
	views.SetDevelopmentMode(b)
	return jet
}

func convertMapToVar(data interface{}) jet.VarMap {
	vars := make(jet.VarMap, 0)
	if mapData, isMap := data.(map[string]interface{}); isMap {
		for k, v := range mapData {
			vars.Set(k, v)
		}
	} else {
		//TODO:log the error request
	}

	return vars
}
