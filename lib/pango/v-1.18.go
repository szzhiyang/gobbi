// Code generated - DO NOT EDIT.
// +build pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// IsSynthesized is a wrapper around the C function pango_font_face_is_synthesized.
func (recv *FontFace) IsSynthesized() bool {
	retC := C.pango_font_face_is_synthesized((*C.PangoFontFace)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}