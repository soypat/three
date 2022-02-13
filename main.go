package three

import (
	"reflect"
	"syscall/js"
	"time"
)

var (
	three = js.Global().Get("THREE")
)

func Init() {
	three = js.Global().Get("THREE")
	if !three.Truthy() {
		panic("unable to find THREE in global namespace. Have you added the script?")
	}
}

// AddScript is a helper function for adding a script and verifying it was added correctly given a namespace.
func AddScript(url string, objName string) {
	script := js.Global().Get("document").Call("createElement", "script")
	script.Set("src", url)
	js.Global().Get("document").Get("head").Call("appendChild", script)
	if objName == "" {
		return
	}
	count := 0
	for {
		count++
		time.Sleep(25 * time.Millisecond)
		if jsObject := js.Global().Get(objName); !jsObject.IsUndefined() {
			break
		} else if count > 100 {
			panic("could not obtain " + objName + " from URL: " + url)
		}
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
		if fv.IsZero() {
			// Skip zero values and nil pointers.
			continue
		}
		switch field.Type.Kind() {
		case reflect.Float64:
			obj.Set(tag, fv.Float())
		case reflect.String:
			obj.Set(tag, fv.String())
		case reflect.Int:
			obj.Set(tag, fv.Int())
		case reflect.Ptr:
			if fv.IsNil() {
				break
			}
			fv = reflect.Indirect(fv)
			if fv.Kind() != reflect.Struct {
				break
			}
			fallthrough
		case reflect.Struct:
			if fv.NumField() == 0 || fv.Field(0).Type() != reflect.TypeOf(js.Value{}) {
				break
			}
			jsv := fv.Field(0).Interface().(js.Value)
			if jsv.Truthy() {
				obj.Set(tag, jsv)
			}
		case reflect.Interface:
			if ifv, ok := fv.Interface().(objecter); ok {
				obj.Set(tag, ifv.getInternalObject())
			}
		}
	}
	return obj
}
