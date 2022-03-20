package three

import (
	"errors"
	"syscall/js"

	"github.com/soypat/gwasm"
)

// three is the handle to the global instance of THREE object.
// Init() must be called before using three.
var three js.Value

// Init must be called after three.js script has been loaded so that
// everything works properly.
func Init() error {
	three = js.Global().Get("THREE")
	if !three.Truthy() {
		return errors.New("unable to find THREE in global namespace. Have you added the script?")
	}
	return nil
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
