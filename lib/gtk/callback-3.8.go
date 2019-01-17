// This is a generated file - DO NOT EDIT
// +build gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"sync"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gboolean callback_tickcallbackHandler(GObject *, GtkWidget*, GdkFrameClock*, gpointer, gpointer);

*/
import "C"

var callbackTickcallbackId int
var callbackTickcallbackMap = make(map[int]TickcallbackCallback)
var callbackTickcallbackLock sync.RWMutex

// TickcallbackCallback is a callback function for a 'TickCallback' callback.
type TickcallbackCallback func(widget *Widget, frameClock *gdk.FrameClock) bool
