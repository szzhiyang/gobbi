package gi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"unsafe"
)

func TestMain(m *testing.M) {
	SetErrorHandler(func(err error) {
		panic(err)
	})

	Require("GLib", "2.0")

	os.Exit(m.Run())
}

func TestFunctionIntegerReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "get_num_processors")

	args := []Arg{}
	retArg := Arg{typ: ArgType_uint}
	ret := fn.Invoke(args, 0, 0, retArg)
	assert.True(t, ret.value.(uint32) > 0)
}

func TestFunctionBooleanReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "mem_is_system_malloc")

	args := []Arg{}
	retArg := Arg{typ: ArgType_boolean}
	ret := fn.Invoke(args, 0, 0, retArg)
	assert.True(t, ret.value.(bool))
}

func TestFunctionIntegerArgPointerReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "malloc")

	in1 := Arg{typ: ArgType_size, value: uint(8), in: true}
	args := []Arg{in1}
	retArg := Arg{typ: ArgType_pointer}
	ret := fn.Invoke(args, 1, 0, retArg)
	assert.NotNil(t, (*byte)(ret.value.(unsafe.Pointer)))
}

func TestFunctionPointerArg(t *testing.T) {
	// malloc some memory, ready to free it
	malloc, _ := Function2InvokerNew("GLib", "malloc")
	in1 := Arg{typ: ArgType_size, value: uint(8), in: true}
	args := []Arg{in1}
	retArg := Arg{typ: ArgType_pointer}
	ret := malloc.Invoke(args, 1, 0, retArg)

	// free is the function with a pointer arg
	free, _ := Function2InvokerNew("GLib", "free")
	in1 = Arg{typ: ArgType_pointer, value: ret.value.(unsafe.Pointer), in: true}
	args = []Arg{in1}
	retArg = Arg{}
	free.Invoke(args, 1, 0, retArg)
}

func BenchmarkFuncCall(b *testing.B) {
	fn, _ := Function2InvokerNew("GLib", "get_num_processors")

	for n := 0; n < b.N; n++ {
		fn.Invoke([]Arg{}, 0, 0, Arg{typ: ArgType_uint})
	}
}
