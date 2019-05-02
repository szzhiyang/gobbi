// This is a generated file - DO NOT EDIT
// +build gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetEllipsizeMode is a wrapper around the C function gtk_tool_shell_get_ellipsize_mode.
func (recv *ToolShell) GetEllipsizeMode() pango.EllipsizeMode {
	retC := C.gtk_tool_shell_get_ellipsize_mode((*C.GtkToolShell)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// GetTextAlignment is a wrapper around the C function gtk_tool_shell_get_text_alignment.
func (recv *ToolShell) GetTextAlignment() float32 {
	retC := C.gtk_tool_shell_get_text_alignment((*C.GtkToolShell)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetTextOrientation is a wrapper around the C function gtk_tool_shell_get_text_orientation.
func (recv *ToolShell) GetTextOrientation() Orientation {
	retC := C.gtk_tool_shell_get_text_orientation((*C.GtkToolShell)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetTextSizeGroup is a wrapper around the C function gtk_tool_shell_get_text_size_group.
func (recv *ToolShell) GetTextSizeGroup() *SizeGroup {
	retC := C.gtk_tool_shell_get_text_size_group((*C.GtkToolShell)(recv.native))
	retGo := SizeGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}
