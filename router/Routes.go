package router

import (
	"strings"
)

type RoutesCallable func(index int, route *Route)

type Routes struct {
	serve    string
	unique   *UniqueString
	list     []*Route
	routeMap map[string]*Route
}

func NewRoutes(name string) *Routes {
	return &Routes{serve: name, unique: NewUniqueString(), list: make([]*Route, 0), routeMap: make(map[string]*Route)}
}

func (this *Routes) ForEach(callable RoutesCallable) {
	for key, value := range this.list {
		callable(key, value)
	}
}

func (this *Routes) Route(method, url string) *Route {
	if route, ok := this.routeMap[UniMd5(method, url)]; ok {
		return route
	} else {
		panic(route.serve +  "route " + method + ":" + url + " non-existent")
	}
}

// 追加 route
func (this *Routes) Append(route *Route) *Routes {
	this.unique.Append(route.unique)
	this.list = append(this.list, route)
	this.routeMap[route.unique] = route
	return this
}

// 收集 route
func CollectRoute(route *Route) {
	route.method = strings.ToUpper(route.method)
	if ! onlySupportMethods.Exist(route.method) {
		panic(route.serve + " route " + route.method + " error, only support:" + onlySupportMethods.String())
	}

	// 生成唯一标识
	route.UniMd5()

	if serve.Routes(route.serve).Unique().Exist(route.unique) {
		panic(route.serve + " route " + route.method + ":" + route.fullPath + " already exist")
	}

	serve.Routes(route.serve).Unique().Append(route.unique)
	serve.AddRoute(route.serve, route)
}

func (this *Routes) Serve() string {
	return this.serve
}

func (this *Routes) Unique() *UniqueString {
	return this.unique
}

func (this *Routes) List() []*Route {
	return this.list
}

// 获取路由集
func GetRoutes(name string) *Routes {
	return serve.Routes(name)
}

func (this *Routes) AddRoute(method string, relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	if nil == RouteAttributes(attributes).Find(ROUTE_SERVE) {
		attributes = RouteAttributes(attributes).Append(SetServe(this.serve))
	}
	addRoute(method, relativePath, handler, attributes...)
	return this
}

func (this *Routes) AddGroup(prefix string, callable interface{}, attributes ...*RouteAttribute) *Routes {
	if nil == RouteAttributes(attributes).Find(ROUTE_SERVE) {
		attributes = RouteAttributes(attributes).Append(SetServe(this.serve))
	}
	addGroup(prefix, callable, attributes...)
	return this
}

func (this *Routes) Ws(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(WEBSOCKET, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Get(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(GET, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Post(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(POST, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Put(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(PUT, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Delete(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(DELETE, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Patch(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(PATCH, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Head(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(HEAD, relativePath, handler, attributes...)
	return this
}
