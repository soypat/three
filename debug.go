package three

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// Debug prints JSON representation of underlying js.Value if found. Not for use
// with common Go types.
func Debug(a ...interface{}) {
	json := js.Global().Get("JSON")
	for _, v := range a {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr && rv.IsNil() {
			fmt.Println("<nil>")
			continue
		}
		rv = reflect.Indirect(rv)
		if rv.Type() == reflect.TypeOf(js.Value{}) {
			fmt.Println(json.Call("stringify", v.(js.Value)).String())
			continue
		}
		if rv.Kind() == reflect.Struct && rv.Field(0).Type() == reflect.TypeOf(js.Value{}) {
			fmt.Println(json.Call("stringify", rv.Field(0).Interface().(js.Value)).String())
			continue
		}
		jv := objectify(v)
		str := json.Call("stringify", jv).String()
		fmt.Println(str)
	}
}
