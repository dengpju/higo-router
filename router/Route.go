package router

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

type Route struct {
	serve        string        // 服务
	method       string        // 请求方法 GET/POST/DELETE/PATCH/OPTIONS/HEAD
	groupPrefix  string        // 组前缀
	relativePath string        // 后端url
	fullPath     string        // 完整url (组前缀 + 后端url)
	handle       interface{}   // 后端控制器函数
	flag         string        // 后端控制器函数标记
	frontPath    string        // 前端 path(前端菜单路由)
	isStatic     bool          // 是否静态文件
	desc         string        // 描述
	middleware   []interface{} // 中间件
	groupMiddle  interface{}   // 组中间件
	unique       string        // 唯一标识 md5(method + ":" fullPath)
	header       http.Header
}

func NewRoute() *Route {
	return &Route{}
}

func (this *Route) Prefix() string {
	return this.groupPrefix
}

func (this *Route) Method() string {
	return this.method
}

func (this *Route) RelativePath() string {
	return this.relativePath
}

func (this *Route) FullPath() string {
	return this.fullPath
}

func (this *Route) Handle() interface{} {
	return this.handle
}

func (this *Route) Flag() string {
	return this.flag
}

func (this *Route) FrontPath() string {
	return this.frontPath
}

func (this *Route) IsStatic() bool {
	return this.isStatic
}

func (this *Route) Desc() string {
	return this.desc
}

func (this *Route) Middleware() interface{} {
	return this.middleware
}

func (this *Route) GroupMiddle() interface{} {
	return this.groupMiddle
}

func (this *Route) Serve() string {
	return this.serve
}

func (this *Route) UniMd5() *Route {
	this.unique = UniMd5(this.method, this.fullPath)
	return this
}

func (this *Route) Unique() string {
	return this.unique
}

func (this *Route) Header() http.Header {
	return this.header
}

func (this *Route) SetHeader(header http.Header) *Route {
	this.header = header
	return this
}

func UniMd5(method, fullPath string) string {
	m5 := md5.New()
	m5.Write([]byte(method + ":" + fullPath))
	return hex.EncodeToString(m5.Sum(nil))
}
