// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

	gboolean callback_pixbufsavefuncHandler(GObject *, gchar*, gsize, GError**, gpointer, gpointer);

*/
import "C"

var callbackPixbufsavefuncId int
var callbackPixbufsavefuncMap = make(map[int]PixbufsavefuncCallback)
var callbackPixbufsavefuncLock sync.RWMutex

// PixbufsavefuncCallback is a callback function for a 'PixbufSaveFunc' callback.
type PixbufsavefuncCallback func(buf []uint8) bool
