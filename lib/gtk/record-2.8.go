// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// BackwardVisibleLine is a wrapper around the C function gtk_text_iter_backward_visible_line.
func (recv *TextIter) BackwardVisibleLine() bool {
	retC := C.gtk_text_iter_backward_visible_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleLines is a wrapper around the C function gtk_text_iter_backward_visible_lines.
func (recv *TextIter) BackwardVisibleLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_visible_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleLine is a wrapper around the C function gtk_text_iter_forward_visible_line.
func (recv *TextIter) ForwardVisibleLine() bool {
	retC := C.gtk_text_iter_forward_visible_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleLines is a wrapper around the C function gtk_text_iter_forward_visible_lines.
func (recv *TextIter) ForwardVisibleLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_visible_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// GetModel is a wrapper around the C function gtk_tree_row_reference_get_model.
func (recv *TreeRowReference) GetModel() *TreeModel {
	retC := C.gtk_tree_row_reference_get_model((*C.GtkTreeRowReference)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}
