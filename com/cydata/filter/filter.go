package filter

import (
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

func (f *Filter) Handle(webHandle WebHandle) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var matcher antpath.PathMatcher
		for path, handle := range f.filterMap {
			if matcher.Match(r.RequestURI, path) {
				//执行拦截业务逻辑
				err := handle(rw, r)
				if err != nil {
					rw.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		//执行正常注册的函数
		webHandle(rw, r)
	}
}
