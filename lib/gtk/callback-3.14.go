// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void callback_listboxforeachfuncHandler(GObject *, GtkListBox*, GtkListBoxRow*, gpointer, gpointer);

*/
import "C"

var callbackListboxforeachfuncId int
var callbackListboxforeachfuncMap = make(map[int]ListboxforeachfuncCallback)
var callbackListboxforeachfuncLock sync.RWMutex

// ListboxforeachfuncCallback is a callback function for a 'ListBoxForeachFunc' callback.
type ListboxforeachfuncCallback func(box *ListBox, row *ListBoxRow)
