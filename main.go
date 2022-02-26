package three

import (
	"syscall/js"

	"github.com/soypat/gwasm"
)

var three = js.Global().Get("THREE")

func Init() {
	three = js.Global().Get("THREE")
	if !three.Truthy() {
		panic("unable to find THREE in global namespace. Have you added the script?")
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

// objectify converts a struct with `js` field tags to
// a javascript Object type with the non-zero, non-nil
// fields set to the struct's values.
func objectify(Struct interface{}) js.Value {
	return gwasm.ValueFromStruct(Struct, true)
}
