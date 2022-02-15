package three

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// Debug prints JSON representation of underlying js.Value if found. Not for use
// with common Go types.
func Debug(a ...interface{}) {
	for _, v := range a {
		fmt.Print(debugs(v) + " ")
	}
	fmt.Println()
}

func debugs(a interface{}) string {
	if s, ok := a.(string); ok {
		return s
	}
	stringify := func(jsv js.Value) string {
		if !jsv.Truthy() {
			return js.Global().Get("String").New(jsv).String()
		}
		return js.Global().Get("JSON").Call("stringify", jsv).String()
	}
	rv := reflect.ValueOf(a)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return "<nil>"
	}
	rv = reflect.Indirect(rv)
	switch {
	case rv.Type() == reflect.TypeOf(js.Value{}):
		// interface is a js.Value.
		return stringify(a.(js.Value))

	case rv.Kind() == reflect.Struct:
		if rv.NumField() == 1 && rv.Field(0).Type() == reflect.TypeOf(js.Value{}) {
			// Single field struct of a js.Value, like most binded types in this package.
			return stringify(rv.Field(0).Interface().(js.Value))
		}
		return stringify(objectify(a))
	}
	return fmt.Sprintf("%+v", a)
}
