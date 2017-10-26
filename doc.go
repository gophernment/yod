package yod

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type ErrorType = error

type App struct {
	errorMapping map[int]ErrorType
}

func New() App {
	return App{}
}

// MapError set return http status code to specific error type
func (a *App) MapError(code int, err ErrorType) {
	a.errorMapping[code] = err
}

func (a *App) Serve(port string) {

}

// Merge every App into one
func Merge(a ...App) App {
	return App{}
}

type ContentDecoder func(v interface{}) error

func JSONDecoder(v interface{}) error {
	return nil
}

func XMLDecoder(v interface{}) error {
	return nil
}

type Request struct {
	*http.Request
	DecodeContent ContentDecoder
	Asset         map[string]interface{}
}

func (r *Request) Decode(v interface{}) error {
	return r.DecodeContent(v)
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

func (r *Request) File(s string) (b []byte, err error) {
	return nil, nil
}

func (r *Request) Set(k string, v interface{}) {}
func (r *Request) Get(k string) interface{} {
	return nil
}

type ResponseWriter interface {
	SetHeader(key, value string)
	OK(v interface{}) error
	InternalServerError(v interface{}) error
	Informational(code int, v interface{}) error
	Successful(code int, v interface{}) error
	Redirection(code int, v interface{}) error
	ClientError(code int, v interface{}) error
	ServerError(code int, v interface{}) error
}

type Handler interface {
	Serve(r *Request, w ResponseWriter) (err error)
}

type HandlerFunc func(r *Request, w ResponseWriter) (err error)

func (fn HandlerFunc) Serve(r *Request, w ResponseWriter) (err error) {
	return fn(r, w)
}

type MiddlewareFunc func(h HandlerFunc) HandlerFunc

func JSONMiddleware() MiddlewareFunc {
	return func(h HandlerFunc) HandlerFunc {
		return func(r *Request, w ResponseWriter) (err error) {
			w = &JSONWriter{}
			r.DecodeContent = JSONDecoder
			return h(r, w)
		}
	}
}

func XMLMiddleware() MiddlewareFunc {
	return func(h HandlerFunc) HandlerFunc {
		return func(r *Request, w ResponseWriter) (err error) {
			w = &XMLWriter{}
			r.DecodeContent = XMLDecoder
			return h(r, w)
		}
	}
}

type Route struct {
	Method string
	Path   string
	H      Handler
	MW     []MiddlewareFunc
	Name   string
}

func NewRoute(method, path string, h Handler, mw ...MiddlewareFunc) Route {
	fn := strings.Split(runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name(), "/")
	return Route{
		Method: method,
		Path:   path,
		H:      h,
		MW:     mw,
		Name:   fn[len(fn)-1],
	}
}

func (r Route) HandlerName() string {
	return r.Name
}
