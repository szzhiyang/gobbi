// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// TreePathNewFromIndicesv is a wrapper around the C function gtk_tree_path_new_from_indicesv.
func TreePathNewFromIndicesv(indices []int32) *TreePath {
	c_indices_array := make([]C.gint, len(indices)+1, len(indices)+1)
	for i, item := range indices {
		c := (C.gint)(item)
		c_indices_array[i] = c
	}
	c_indices_array[len(indices)] = 0
	c_indices_arrayPtr := &c_indices_array[0]
	c_indices := (*C.gint)(unsafe.Pointer(c_indices_arrayPtr))

	c_length := (C.gsize)(len(indices))

	retC := C.gtk_tree_path_new_from_indicesv(c_indices, c_length)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}
