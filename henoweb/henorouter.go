package henoweb

import (
	"fmt"
	"net/http"
	"strings"
)

type HenoRouter struct {
	HandleMap map[string]interface{}
}

var hRouter = NewRouter()

func NewRouter() *HenoRouter {
	r := &HenoRouter{}
	r.HandleMap = make(map[string]interface{})
	return r
}

func RegisterFuncRouter(pattern string, h http.HandlerFunc) {
	hRouter.HandleMap[pattern] = h
}

func RegisterControllerRouter(pattern string, c Controller) {
	hRouter.HandleMap[strings.TrimLeft(pattern, "/")] = c
}

func (router *HenoRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, ok := router.HandleMap[r.URL.Path]
	var action, controller string
	action, controller = getControllerAction(r.URL.Path)
	if !ok {
		h, ok = router.HandleMap[controller]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "~opps 404~")
			return
		}
	}

	switch h.(type) {
	case http.HandlerFunc:
		h.(http.HandlerFunc).ServeHTTP(w, r)
	case Controller:
		action = Ucfirst(action)
		err := Invoke(h, action, w, r)
		if err != nil {
			http.HandlerFunc(Page404).ServeHTTP(w, r)
		}
	}
}

func Page404(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "page-404-no-action")
}

func getControllerAction(path string) (action, controller string) {
	pathItem := strings.Split(strings.Trim(path, "/"), "/")
	fmt.Println(pathItem)
	controller = "index"
	action = "index"
	if len(pathItem) >= 1 {
		controller = pathItem[0]
	}
	if len(pathItem) > 1 {
		action = pathItem[1]
	}
	return
}

func Ucfirst(str string) string {
	if str == "" {
		return str
	}
	b := []byte(str)
	if b[0] >= 97 && b[0] <= 122 {
		b[0] = b[0] - 32
	}
	return string(b)
}
