// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type ApplicationInhibitFlags C.GtkApplicationInhibitFlags

const (
	GTK_APPLICATION_INHIBIT_LOGOUT  ApplicationInhibitFlags = 1
	GTK_APPLICATION_INHIBIT_SWITCH  ApplicationInhibitFlags = 2
	GTK_APPLICATION_INHIBIT_SUSPEND ApplicationInhibitFlags = 4
	GTK_APPLICATION_INHIBIT_IDLE    ApplicationInhibitFlags = 8
)