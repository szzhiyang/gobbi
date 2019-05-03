// Code generated - DO NOT EDIT.
// +build pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// QuantizeLineGeometry is a wrapper around the C function pango_quantize_line_geometry.
func QuantizeLineGeometry(thickness int32, position int32) {
	c_thickness := (C.int)(thickness)

	c_position := (C.int)(position)

	C.pango_quantize_line_geometry(&c_thickness, &c_position)

	return
}

// GetFontScaleFactor is a wrapper around the C function pango_matrix_get_font_scale_factor.
func (recv *Matrix) GetFontScaleFactor() float64 {
	retC := C.pango_matrix_get_font_scale_factor((*C.PangoMatrix)(recv.native))
	retGo := (float64)(retC)

	return retGo
}