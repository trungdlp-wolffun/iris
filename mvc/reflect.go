package mvc

import (
	"reflect"

	"github.com/kataras/iris/v12/hero"
)

var (
	baseControllerTyp = reflect.TypeOf((*BaseController)(nil)).Elem()
	errorHandlerTyp   = reflect.TypeOf((*hero.ErrorHandler)(nil)).Elem()
	errorTyp          = reflect.TypeOf((*error)(nil)).Elem()
)

func isBaseController(ctrlTyp reflect.Type) bool {
	return ctrlTyp.Implements(baseControllerTyp)
}

// indirectType returns the value of a pointer-type "typ".
// If "typ" is a pointer, array, chan, map or slice it returns its Elem,
// otherwise returns the typ as it's.
func indirectType(typ reflect.Type) reflect.Type {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return typ.Elem()
	}
	return typ
}

func isErrorHandler(ctrlTyp reflect.Type) bool {
	return ctrlTyp.Implements(errorHandlerTyp)
}

func hasErrorOutArgs(fn reflect.Method) bool {
	n := fn.Type.NumOut()
	if n == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if out := fn.Type.Out(i); out.Kind() == reflect.Interface {
			if out.Implements(errorTyp) {
				return true
			}
		}
	}

	return false
}

func getInputArgsFromFunc(funcTyp reflect.Type) []reflect.Type {
	n := funcTyp.NumIn()
	funcIn := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		funcIn[i] = funcTyp.In(i)
	}
	return funcIn
}
