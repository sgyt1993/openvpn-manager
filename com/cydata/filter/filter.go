package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ovpn-admin/com/cydata/antpath"
)

type FilterHandle func(rw http.ResponseWriter, req *http.Request) error
type WebHandle func(rw http.ResponseWriter, req *http.Request)

type Filter struct {
	filterMap map[string]FilterHandle
}

func New() *Filter {
	return &Filter{
		filterMap: make(map[string]FilterHandle),
	}
}

func (f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	f.filterMap[uri] = handler
}

func (f *Filter) GetFilterHandle(uri string) FilterHandle {
	return f.filterMap[uri]
}

var matcher = antpath.New()

func (f *Filter) Handle(ctx *gin.Context) {
	for path, handle := range f.filterMap {
		if matcher.Match(path, ctx.Request.RequestURI) {
			//执行拦截业务逻辑
			err := handle(ctx.Writer, ctx.Request)
			if err != nil {
				ctx.Writer.Write([]byte(err.Error()))
				return
			}
			break
		}
	}
}

//func (f *Filter) Handle(webHandle WebHandle) func(rw http.ResponseWriter, r *http.Request) {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		for path, handle := range f.filterMap {
//			if matcher.Match(path, r.RequestURI) {
//				//执行拦截业务逻辑
//				err := handle(rw, r)
//				if err != nil {
//					rw.Write([]byte(err.Error()))
//					return
//				}
//				break
//			}
//		}
//		//执行正常注册的函数
//		webHandle(rw, r)
//	}
//}
