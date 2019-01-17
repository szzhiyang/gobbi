// This is a generated file - DO NOT EDIT
// +build gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
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

	GtkWidget* callback_flowboxcreatewidgetfuncHandler(GObject *, gpointer, gpointer, gpointer);

*/
import "C"

var callbackFlowboxcreatewidgetfuncId int
var callbackFlowboxcreatewidgetfuncMap = make(map[int]FlowboxcreatewidgetfuncCallback)
var callbackFlowboxcreatewidgetfuncLock sync.RWMutex

// FlowboxcreatewidgetfuncCallback is a callback function for a 'FlowBoxCreateWidgetFunc' callback.
type FlowboxcreatewidgetfuncCallback func(item *gobject.Object) *Widget
