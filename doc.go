package yod

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type errorType = error

type Yod struct {
}

func New() Yod {
	return Yod{}
}

// MapError set return http status code to specific error type
func (y *Yod) MapError(code int, err errorType) {

}

func Merge(y ...Yod) Yod {
	return Yod
}

type Request {
	*http.Request
	Asset map[string]interface{}
}

func (r *Request) Decode(v interface{}) error {
	return nil
}

func (r *Request) Param(s string) string {
	return ""
}

func (r *Request) Query(s string) string {
	return ""
}

func (r *Request) FormValue(s string) string {
	return ""
}

func (r *Request) File(s string) (b []byte,err error) {
	return nil,nil
}

func (r *Request) Set(k string, v inteface{}) {}
func (r *Request) Get(k string) v inteface{} {
	return nil
}

type ResponseWriter interface {
	Header(v interface{}) error
	Write(code int, v interface{}) error
	Code(i int)
}

type Handler func(r *Request, w ResponseWriter) (err error) {
	return nil
}

type Middleware func(h Handler) Handler

type Route struct {
	Method string
	Path   string
	H      Handler
	MW     []Middleware
	Name string
}

func NewRoute(method, path string, h Handler, mw ...Middleware) Route {
	fn := strings.Split(runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name(), "/")
	return Route{
		Method: method,
		Path:   path,
		H:      h,
		MW:     mw,
		Name :fn[len(fn)-1],
	}
}

func (r Route) HandlerName() string {
	return r.Name
}
