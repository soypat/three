package three

import (
	"reflect"
	"syscall/js"
)

var (
	three = js.Global().Get("THREE")
)

func init() {
	if three.IsUndefined() {
		panic("THREE (three.js) is undefined at global level")
	}
}

// getModule gets a module or property added to THREE or globalThis.
func getModule(namespace string) js.Value {
	mod := three.Get(namespace)
	if !mod.IsUndefined() {
		return mod
	}
	mod = js.Global().Get(namespace)
	if !mod.IsUndefined() {
		return mod
	}
	panic("three:failed to get " + namespace + " namespace from THREE and global namespace")
}

func objectify(Struct interface{}) js.Value {
	v := reflect.ValueOf(Struct)
	vi := reflect.Indirect(v)
	if vi.Kind() != reflect.Struct {
		panic("expected struct input to objectify")
	}
	obj := js.Global().Get("Object").New()
	recordType := vi.Type()
	for i, field := range reflect.VisibleFields(recordType) {
		tag := field.Tag.Get("js")
		if tag == "" {
			continue
		}
		fv := vi.Field(i)
		switch field.Type.Kind() {
		case reflect.Float64, reflect.Int, reflect.String:
			obj.Set(tag, fv.Interface())
		case reflect.Ptr:
			fv = reflect.Indirect(fv)
			if fv.Kind() != reflect.Struct {
				break
			}
			fallthrough
		case reflect.Struct:
			if fv.NumField() == 0 || fv.Field(0).Type() != reflect.TypeOf(js.Value{}) {
				break
			}
			obj.Set(tag, fv.Interface().(js.Value))
		}
	}
	return obj
}
