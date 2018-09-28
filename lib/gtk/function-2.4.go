// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// BindingsActivateEvent is a wrapper around the C function gtk_bindings_activate_event.
func BindingsActivateEvent(object *gobject.Object, event *gdk.EventKey) bool {
	c_object := (*C.GObject)(object.ToC())

	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_bindings_activate_event(c_object, c_event)
	retGo := retC == C.TRUE

	return retGo
}

// FileChooserErrorQuark is a wrapper around the C function gtk_file_chooser_error_quark.
func FileChooserErrorQuark() glib.Quark {
	retC := C.gtk_file_chooser_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gtk_rc_reset_styles : no return generator