// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

func (recv *Pixbuf) Object() *gobject.Object {}

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

func (recv *PixbufAnimation) Object() *gobject.Object {}

func (recv *PixbufAnimationIter) Object() *gobject.Object {}

func (recv *PixbufLoader) Object() *gobject.Object {}

// PixbufSimpleAnimNew is a wrapper around the C function gdk_pixbuf_simple_anim_new.
func PixbufSimpleAnimNew(width int32, height int32, rate float32) *PixbufSimpleAnim {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_rate := (C.gfloat)(rate)

	retC := C.gdk_pixbuf_simple_anim_new(c_width, c_height, c_rate)
	retGo := PixbufSimpleAnimNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFrame is a wrapper around the C function gdk_pixbuf_simple_anim_add_frame.
func (recv *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gdk_pixbuf_simple_anim_add_frame((*C.GdkPixbufSimpleAnim)(recv.native), c_pixbuf)

	return
}

func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {}

// Unsupported : PixbufSimpleAnimIter : no CType
